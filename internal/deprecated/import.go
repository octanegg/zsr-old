package deprecated

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/octanegg/zsr/ballchasing"
	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const (
	linkagesSQL = "SELECT old_event, old_stage, new_event, new_stage FROM mapping"
	numWorkers  = 40
)

// EventLinkage .
type EventLinkage struct {
	OldEvent int                `json:"old_event"`
	OldStage int                `json:"old_stage"`
	NewEvent primitive.ObjectID `json:"new_event"`
	NewStage int                `json:"new_stage"`
}

func (h *handler) Import(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var events []int
	if err := json.NewDecoder(r.Body).Decode(&events); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	linkages, err := h.Deprecated.getLinkages(events)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	fmt.Printf("Importing %d events\n", len(linkages))

	done := 0
	sem := semaphore.NewWeighted(numWorkers)
	errs, _ := errgroup.WithContext(context.TODO())
	for _, linkage := range linkages {
		linkage := linkage
		errs.Go(func() error {
			if err = sem.Acquire(context.TODO(), 1); err != nil {
				return err
			}
			defer sem.Release(1)

			start := time.Now()
			if err = h.singleImport(linkage); err != nil {
				return err
			}
			done++

			fmt.Printf("%d / %d: Finished importing %d (%d) => %s (%d) in %f seconds\n",
				done,
				len(linkages),
				linkage.OldEvent,
				linkage.OldStage,
				linkage.NewEvent.Hex(),
				linkage.NewStage,
				time.Since(start).Seconds(),
			)

			return nil
		})
	}

	if err = errs.Wait(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) singleImport(linkage *EventLinkage) error {
	data, err := h.Octane.Events().FindOne(bson.M{"_id": &linkage.NewEvent})
	if err != nil {
		return err
	}

	event := data.(octane.Event)

	_, err = h.Octane.Statlines().Delete(bson.M{"game.match.event._id": linkage.NewEvent, "game.match.stage._id": linkage.NewStage})
	if err != nil {
		return err
	}

	matches, err := h.getMatches(linkage, &event)
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		return nil
	}
	for _, m := range matches {
		var allGames []*octane.Game
		allPlayerStats := map[string][]*octane.Statline{}
		match := m.(*octane.Match)
		if match.Blue != nil && match.Orange != nil {
			games, err := h.getGames(match)
			if err != nil {
				return err
			}

			if len(games) > 0 {
				var matchStats []*octane.Statline
				for _, game := range games {
					gameStats := h.getStats(game)
					matchStats = append(matchStats, gameStats...)

					for game.Number > len(match.Games)+1 {
						match.Games = append(match.Games, &octane.GameOverview{})
					}
					match.Games = append(match.Games, &octane.GameOverview{
						ID:       game.ID,
						Blue:     game.Blue.Team.Stats.Core.Goals,
						Orange:   game.Orange.Team.Stats.Core.Goals,
						Duration: game.Duration,
					})

					allPlayerStats[fmt.Sprintf("%s-%d", game.OctaneID, game.Number)] = gameStats
				}

				blue, orange := map[string]*octane.PlayerStats{}, map[string]*octane.PlayerStats{}
				match.Blue.Team.Stats = &ballchasing.TeamStats{
					Core: &ballchasing.TeamCore{},
				}
				match.Orange.Team.Stats = &ballchasing.TeamStats{
					Core: &ballchasing.TeamCore{},
				}
				for _, stat := range matchStats {
					if stat.Team.Team.ID == match.Blue.Team.Team.ID {
						match.Blue.Team.Stats.Core.Score += stat.Player.Stats.Core.Score
						match.Blue.Team.Stats.Core.Goals += stat.Player.Stats.Core.Goals
						match.Blue.Team.Stats.Core.Assists += stat.Player.Stats.Core.Assists
						match.Blue.Team.Stats.Core.Saves += stat.Player.Stats.Core.Saves
						match.Blue.Team.Stats.Core.Shots += stat.Player.Stats.Core.Shots
						if match.Blue.Team.Stats.Core.Shots > 0 {
							match.Blue.Team.Stats.Core.ShootingPercentage = float64(match.Blue.Team.Stats.Core.Goals) / float64(match.Blue.Team.Stats.Core.Shots) * 100
						}

						if _, ok := blue[stat.Player.Player.ID.Hex()]; !ok {
							blue[stat.Player.Player.ID.Hex()] = &octane.PlayerStats{
								Player: stat.Player.Player,
								Stats: &ballchasing.PlayerStats{
									Core: &ballchasing.PlayerCore{},
								},
								Advanced: &octane.AdvancedStats{},
							}
						}

						blue[stat.Player.Player.ID.Hex()].Stats.Core.Score += stat.Player.Stats.Core.Score
						blue[stat.Player.Player.ID.Hex()].Stats.Core.Goals += stat.Player.Stats.Core.Goals
						blue[stat.Player.Player.ID.Hex()].Stats.Core.Assists += stat.Player.Stats.Core.Assists
						blue[stat.Player.Player.ID.Hex()].Stats.Core.Saves += stat.Player.Stats.Core.Saves
						blue[stat.Player.Player.ID.Hex()].Stats.Core.Shots += stat.Player.Stats.Core.Shots
						blue[stat.Player.Player.ID.Hex()].Advanced.Rating += stat.Player.Advanced.Rating
					} else {
						match.Orange.Team.Stats.Core.Score += stat.Player.Stats.Core.Score
						match.Orange.Team.Stats.Core.Goals += stat.Player.Stats.Core.Goals
						match.Orange.Team.Stats.Core.Assists += stat.Player.Stats.Core.Assists
						match.Orange.Team.Stats.Core.Saves += stat.Player.Stats.Core.Saves
						match.Orange.Team.Stats.Core.Shots += stat.Player.Stats.Core.Shots
						if match.Orange.Team.Stats.Core.Shots > 0 {
							match.Orange.Team.Stats.Core.ShootingPercentage = float64(match.Orange.Team.Stats.Core.Goals) / float64(match.Orange.Team.Stats.Core.Shots) * 100
						}

						if _, ok := orange[stat.Player.Player.ID.Hex()]; !ok {
							orange[stat.Player.Player.ID.Hex()] = &octane.PlayerStats{
								Player: stat.Player.Player,
								Stats: &ballchasing.PlayerStats{
									Core: &ballchasing.PlayerCore{},
								},
								Advanced: &octane.AdvancedStats{},
							}
						}

						orange[stat.Player.Player.ID.Hex()].Stats.Core.Score += stat.Player.Stats.Core.Score
						orange[stat.Player.Player.ID.Hex()].Stats.Core.Goals += stat.Player.Stats.Core.Goals
						orange[stat.Player.Player.ID.Hex()].Stats.Core.Assists += stat.Player.Stats.Core.Assists
						orange[stat.Player.Player.ID.Hex()].Stats.Core.Saves += stat.Player.Stats.Core.Saves
						orange[stat.Player.Player.ID.Hex()].Stats.Core.Shots += stat.Player.Stats.Core.Shots
						orange[stat.Player.Player.ID.Hex()].Advanced.Rating += stat.Player.Advanced.Rating
					}
				}

				for _, player := range blue {
					if player.Stats.Core.Shots > 0 {
						player.Stats.Core.ShootingPercentage = float64(player.Stats.Core.Goals) / float64(player.Stats.Core.Shots) * 100
					}
					if match.Blue.Team.Stats.Core.Goals > 0 {
						player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / float64(match.Blue.Team.Stats.Core.Goals) * 100
					}
					player.Advanced.Rating /= float64(len(games))
					match.Blue.Players = append(match.Blue.Players, player)
				}

				for _, player := range orange {
					if player.Stats.Core.Shots > 0 {
						player.Stats.Core.ShootingPercentage = float64(player.Stats.Core.Goals) / float64(player.Stats.Core.Shots) * 100
					}
					if match.Orange.Team.Stats.Core.Goals > 0 {
						player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / float64(match.Orange.Team.Stats.Core.Goals) * 100
					}
					player.Advanced.Rating /= float64(len(games))
					match.Orange.Players = append(match.Orange.Players, player)
				}

				allGames = append(allGames, games...)

				reverseSweepAttempt, reverseSweep := getSweepData(games)
				match.ReverseSweepAttempt = reverseSweepAttempt
				match.ReverseSweep = reverseSweep
			}
		}

		var matchID, gameID primitive.ObjectID

		m, err := h.Octane.Matches().FindOne(bson.M{"octane_id": match.OctaneID})
		if err != nil {
			id, err := h.Octane.Matches().InsertOne(match)
			if err != nil {
				return err
			}
			matchID = *id
		} else {
			match := m.(octane.Match)
			matchID = *match.ID
			match.ID = nil
			if _, err = h.Octane.Matches().UpdateOne(bson.M{"_id": matchID}, bson.M{"$set": match}); err != nil {
				return err
			}
		}

		for _, game := range allGames {
			game.Match.ID = &matchID
			g, err := h.Octane.Games().FindOne(bson.M{"octane_id": game.OctaneID, "number": game.Number})
			if err != nil {
				id, err := h.Octane.Games().InsertOne(game)
				if err != nil {
					return err
				}
				gameID = *id
			} else {
				gameID = *g.(octane.Game).ID
				game.ID = nil
				if _, err = h.Octane.Games().UpdateOne(bson.M{"_id": gameID}, bson.M{"$set": game}); err != nil {
					return err
				}
			}

			var finalStats []interface{}
			for _, stat := range allPlayerStats[fmt.Sprintf("%s-%d", game.OctaneID, game.Number)] {
				stat.Game.Match.ID = &matchID
				stat.Game.ID = &gameID
				finalStats = append(finalStats, stat)
			}

			if len(finalStats) > 0 {
				if _, err := h.Octane.Statlines().Insert(finalStats); err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func getSweepData(games []*octane.Game) (bool, bool) {
	var format int
	var _games []*octane.Game
	for _, game := range games {
		_games = append(_games, game)
		format = game.Match.Format.Length
	}

	var isReverseSweep, isReverseSweepAttempt bool

	if len(games) == format {
		isReverseSweepAttempt = true
		firstWinnerIsBlue := _games[0].Blue.Winner
		for i := 1; i < format/2; i++ {
			if (firstWinnerIsBlue && !_games[i].Blue.Winner) || (!firstWinnerIsBlue && !_games[i].Orange.Winner) {
				isReverseSweepAttempt = false
				break
			}
		}

		if isReverseSweepAttempt {
			isReverseSweep = (firstWinnerIsBlue && _games[format-1].Orange.Winner) || (!firstWinnerIsBlue && _games[format-1].Blue.Winner)
		}
	}

	return isReverseSweepAttempt, isReverseSweep
}

func (d *deprecated) getLinkages(events []int) ([]*EventLinkage, error) {
	stmt := linkagesSQL
	if len(events) > 0 {
		stmt += fmt.Sprintf(" WHERE old_event IN (%s)", strings.Join(intsToStrings(events), ","))
	}

	results, err := d.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var linkages []*EventLinkage
	for results.Next() {
		var linkage EventLinkage
		var newEvent string
		err = results.Scan(&linkage.OldEvent, &linkage.OldStage, &newEvent, &linkage.NewStage)
		if err != nil {
			return nil, err
		}

		if linkage.NewEvent, err = primitive.ObjectIDFromHex(newEvent); err != nil {
			return nil, err
		}

		linkages = append(linkages, &linkage)
	}

	return linkages, nil
}

func (h *handler) getMatches(linkage *EventLinkage, event *octane.Event) ([]interface{}, error) {
	oldMatches, err := h.Deprecated.GetMatches(&GetMatchesContext{
		Event: strconv.Itoa(linkage.OldEvent),
		Stage: strconv.Itoa(linkage.OldStage),
	})
	if err != nil {
		return nil, err
	}

	var newMatches []interface{}
	for _, match := range oldMatches {
		id := primitive.NewObjectID()
		newMatch := &octane.Match{
			ID:       &id,
			OctaneID: match.OctaneID,
			Date:     match.Date,
			Number:   match.Number,
			Event: &octane.Event{
				ID:     event.ID,
				Name:   event.Name,
				Mode:   event.Mode,
				Region: event.Region,
				Tier:   event.Tier,
				Image:  event.Image,
				Groups: event.Groups,
			},
			Stage: &octane.Stage{
				ID:     linkage.NewStage,
				Name:   event.Stages[linkage.NewStage].Name,
				Format: event.Stages[linkage.NewStage].Format,
			},
		}

		if event.Stages[linkage.NewStage].Qualifier {
			newMatch.Stage.Qualifier = true
		}

		if match.Blue.Name != "" && match.Orange.Name != "" {
			newMatch.Blue = &octane.MatchSide{
				Team: &octane.TeamStats{
					Team: h.findOrInsertTeam(match.Blue.Name),
				},
				Score:  match.Blue.Score,
				Winner: match.Blue.Winner,
			}
			newMatch.Orange = &octane.MatchSide{
				Team: &octane.TeamStats{
					Team: h.findOrInsertTeam(match.Orange.Name),
				},
				Score:  match.Orange.Score,
				Winner: match.Orange.Winner,
			}

			winnerScore := match.Blue.Score
			if match.Orange.Score > match.Blue.Score {
				winnerScore = match.Orange.Score
			}

			newMatch.Format = &octane.Format{
				Type:   "best",
				Length: winnerScore*2 - 1,
			}
		}

		newMatches = append(newMatches, newMatch)
	}

	return newMatches, nil
}

func (h *handler) getGames(match *octane.Match) ([]*octane.Game, error) {
	oldGames, err := h.Deprecated.GetGames(&GetGamesContext{
		OctaneID: match.OctaneID,
		Blue:     match.Blue.Team.Team.Name,
		Orange:   match.Orange.Team.Team.Name,
	})
	if err != nil {
		return nil, err
	}

	var newGames []*octane.Game
	for _, game := range oldGames {
		if game.Blue.Name == match.Orange.Team.Team.Name {
			game.Blue, game.Orange = game.Orange, game.Blue
		}

		id := primitive.NewObjectID()
		newGame := &octane.Game{
			ID:       &id,
			OctaneID: game.OctaneID,
			Number:   game.Number,
			Map: &octane.Map{
				Name: game.Map,
			},
			Duration: game.Duration,
			Blue: &octane.GameSide{
				Team: &octane.TeamStats{
					Team:  match.Blue.Team.Team,
					Stats: toTeamStats(game.Blue.Players),
				},
				Players: h.toPlayers(game.Blue.Players),
			},
			Orange: &octane.GameSide{
				Team: &octane.TeamStats{
					Team:  match.Orange.Team.Team,
					Stats: toTeamStats(game.Orange.Players),
				},
				Players: h.toPlayers(game.Orange.Players),
			},
			Match: &octane.Match{
				ID:     match.ID,
				Format: match.Format,
				Event:  match.Event,
				Stage:  match.Stage,
			},
			Date: match.Date,
		}

		newGame.Blue.Winner = newGame.Blue.Team.Stats.Core.Goals > newGame.Orange.Team.Stats.Core.Goals
		newGame.Orange.Winner = newGame.Orange.Team.Stats.Core.Goals > newGame.Blue.Team.Stats.Core.Goals

		newGames = append(newGames, newGame)
	}

	return newGames, nil
}

func (h *handler) getStats(game *octane.Game) []*octane.Statline {
	var stats []*octane.Statline

	for _, p := range game.Blue.Players {
		id := primitive.NewObjectID()
		stats = append(stats, &octane.Statline{
			ID: &id,
			Game: &octane.Game{
				ID:       game.ID,
				Match:    game.Match,
				Date:     game.Date,
				Map:      game.Map,
				Duration: game.Duration,
			},
			Team: &octane.StatlineSide{
				Score:   game.Blue.Team.Stats.Core.Goals,
				Winner:  game.Blue.Winner,
				Team:    game.Blue.Team.Team,
				Stats:   game.Blue.Team.Stats,
				Players: getGamePlayers(game.Blue.Players),
			},
			Opponent: &octane.StatlineSide{
				Score:   game.Orange.Team.Stats.Core.Goals,
				Winner:  game.Orange.Winner,
				Team:    game.Orange.Team.Team,
				Stats:   game.Orange.Team.Stats,
				Players: getGamePlayers(game.Orange.Players),
			},
			Player: p,
		})
	}

	for _, p := range game.Orange.Players {
		id := primitive.NewObjectID()
		stats = append(stats, &octane.Statline{
			ID: &id,
			Game: &octane.Game{
				ID:       game.ID,
				Match:    game.Match,
				Date:     game.Date,
				Map:      game.Map,
				Duration: game.Duration,
			},
			Team: &octane.StatlineSide{
				Score:   game.Orange.Team.Stats.Core.Goals,
				Winner:  game.Orange.Winner,
				Team:    game.Orange.Team.Team,
				Stats:   game.Orange.Team.Stats,
				Players: getGamePlayers(game.Orange.Players),
			},
			Opponent: &octane.StatlineSide{
				Score:   game.Blue.Team.Stats.Core.Goals,
				Winner:  game.Blue.Winner,
				Team:    game.Blue.Team.Team,
				Stats:   game.Blue.Team.Stats,
				Players: getGamePlayers(game.Blue.Players),
			},
			Player: p,
		})
	}

	return stats
}

func toTeamStats(logs []Log) *ballchasing.TeamStats {
	stats := &ballchasing.TeamStats{
		Core: &ballchasing.TeamCore{},
	}

	for _, log := range logs {
		stats.Core.Score += log.Score
		stats.Core.Goals += log.Goals
		stats.Core.Assists += log.Assists
		stats.Core.Saves += log.Saves
		stats.Core.Shots += log.Shots
		if stats.Core.Shots > 0 {
			stats.Core.ShootingPercentage = float64(stats.Core.Goals) / float64(stats.Core.Shots) * 100
		}
	}

	return stats
}

func (h *handler) toPlayers(logs []Log) []*octane.PlayerStats {
	var players []*octane.PlayerStats
	for _, log := range logs {
		player := &octane.PlayerStats{
			Player: h.findOrInsertPlayer(log.Player),
			Stats: &ballchasing.PlayerStats{
				Core: &ballchasing.PlayerCore{
					Score:   log.Score,
					Goals:   log.Goals,
					Assists: log.Assists,
					Saves:   log.Saves,
					Shots:   log.Shots,
				},
			},
			Advanced: &octane.AdvancedStats{
				MVP:    log.MVP,
				Rating: log.Rating,
			},
		}

		if log.Shots > 0 {
			player.Stats.Core.ShootingPercentage = float64(log.Goals) / float64(log.Shots) * 100
		}

		if log.TeamGoals > 0 {
			player.Advanced.GoalParticipation = float64(log.Goals+log.Assists) / float64(log.TeamGoals) * 100
		}

		players = append(players, player)
	}

	return players
}

func getGamePlayers(playerStats []*octane.PlayerStats) []*octane.Player {
	var players []*octane.Player
	for _, stats := range playerStats {
		players = append(players, stats.Player)
	}
	return players
}

func intsToStrings(a []int) []string {
	b := make([]string, len(a))
	for i, n := range a {
		b[i] = strconv.Itoa(n)
	}
	return b
}

func (h *handler) findOrInsertTeam(name string) *octane.Team {
	t, err := h.Octane.Teams().FindOne(bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf("^%s$", name), Options: "i"}}})
	if err == nil {
		team := t.(octane.Team)
		return &octane.Team{
			ID:    team.ID,
			Name:  team.Name,
			Image: team.Image,
		}
	}

	id := primitive.NewObjectID()
	team, _ := h.Octane.Teams().InsertOne(&octane.Team{
		ID:   &id,
		Name: name,
	})

	return &octane.Team{
		ID:   team,
		Name: name,
	}
}

func (h *handler) findOrInsertPlayer(tag string) *octane.Player {
	p, err := h.Octane.Players().FindOne(bson.M{"tag": bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf("^%s$", tag), Options: "i"}}})
	if err == nil {
		player := p.(octane.Player)
		return &octane.Player{
			ID:      player.ID,
			Tag:     player.Tag,
			Country: player.Country,
		}
	}

	id := primitive.NewObjectID()
	player, _ := h.Octane.Players().InsertOne(&octane.Player{
		ID:  &id,
		Tag: tag,
	})

	return &octane.Player{
		ID:  player,
		Tag: tag,
	}
}
