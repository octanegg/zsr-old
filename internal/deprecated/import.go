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
	numWorkers  = 50
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

			if err = h.singleImport(linkage); err != nil {
				return err
			}
			done++

			fmt.Printf("%d / %d: Finished importing %+v\n", done, len(linkages), linkage)

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

	_, err = h.Octane.Matches().Delete(bson.M{"event._id": linkage.NewEvent, "stage._id": linkage.NewStage})
	if err != nil {
		return err
	}

	_, err = h.Octane.Games().Delete(bson.M{"match.event._id": linkage.NewEvent, "match.stage._id": linkage.NewStage})
	if err != nil {
		return err
	}

	_, err = h.Octane.Statlines().Delete(bson.M{"game.match.event._id": linkage.NewEvent, "game.match.stage._id": linkage.NewStage})
	if err != nil {
		return err
	}

	allMatches, err := h.getMatches(linkage, &event)
	if err != nil {
		return err
	}

	if len(allMatches) == 0 {
		return nil
	}

	var allPlayerStats, allGames []interface{}
	for _, m := range allMatches {
		match := m.(*octane.Match)
		if match.Blue.Team == nil || match.Orange.Team == nil {
			continue
		}

		games, err := h.getGames(match)
		if err != nil {
			return err
		}

		if len(games) == 0 {
			continue
		}

		for _, g := range games {
			game := g.(*octane.Game)
			allPlayerStats = append(allPlayerStats, h.getStats(game)...)
		}
		allGames = append(allGames, games...)
	}

	if len(allMatches) > 0 {
		_, err = h.Octane.Matches().Insert(allMatches)
		if err != nil {
			return err
		}
	}

	if len(allGames) > 0 {
		_, err = h.Octane.Games().Insert(allGames)
		if err != nil {
			return err
		}
	}

	if len(allPlayerStats) > 0 {
		_, err = h.Octane.Statlines().Insert(allPlayerStats)
		if err != nil {
			return err
		}
	}

	return nil
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
			Format:   match.Format,
			Number:   match.Number,
			Blue: &octane.MatchSide{
				Score:  match.Blue.Score,
				Winner: match.Blue.Winner,
			},
			Orange: &octane.MatchSide{
				Score:  match.Orange.Score,
				Winner: match.Orange.Winner,
			},
			Event: &octane.Event{
				ID:     event.ID,
				Name:   event.Name,
				Mode:   event.Mode,
				Region: event.Region,
				Tier:   event.Tier,
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
			newMatch.Blue.Team = h.findOrInsertTeam(match.Blue.Name)
			newMatch.Orange.Team = h.findOrInsertTeam(match.Orange.Name)
		}

		newMatches = append(newMatches, newMatch)
	}

	return newMatches, nil
}

func (h *handler) getGames(match *octane.Match) ([]interface{}, error) {
	oldGames, err := h.Deprecated.GetGames(&GetGamesContext{
		OctaneID: match.OctaneID,
		Blue:     match.Blue.Team.Name,
		Orange:   match.Orange.Team.Name,
	})
	if err != nil {
		return nil, err
	}

	var newGames []interface{}
	for _, game := range oldGames {
		if game.Blue.Name == match.Orange.Team.Name {
			game.Blue, game.Orange = game.Orange, game.Blue
		}

		id := primitive.NewObjectID()
		newGame := &octane.Game{
			ID:       &id,
			OctaneID: game.OctaneID,
			Number:   game.Number,
			Map:      game.Map,
			Duration: game.Duration,
			Blue: &octane.GameSide{
				Team:    match.Blue.Team,
				Stats:   toTeamStats(game.Blue.Players),
				Players: h.toPlayers(game.Blue.Players),
			},
			Orange: &octane.GameSide{
				Team:    match.Orange.Team,
				Stats:   toTeamStats(game.Orange.Players),
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

		newGame.Blue.Winner = newGame.Blue.Stats.Core.Goals > newGame.Orange.Stats.Core.Goals
		newGame.Orange.Winner = newGame.Orange.Stats.Core.Goals > newGame.Blue.Stats.Core.Goals

		newGames = append(newGames, newGame)
	}

	return newGames, nil
}

func (h *handler) getStats(game *octane.Game) []interface{} {
	var stats []interface{}

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
			Team:     game.Blue.Team,
			Opponent: game.Orange.Team,
			Winner:   game.Blue.Winner,
			Player:   p.Player,
			Stats: &octane.StatlineStats{
				Player: p.Stats,
				Team:   game.Blue.Stats,
			},
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
			Team:     game.Orange.Team,
			Opponent: game.Blue.Team,
			Winner:   game.Orange.Winner,
			Player:   p.Player,
			Stats: &octane.StatlineStats{
				Player: p.Stats,
				Team:   game.Orange.Stats,
			},
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
			stats.Core.ShootingPercentage = float64(stats.Core.Goals) / float64(stats.Core.Shots)
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
					Mvp:     log.MVP,
					Rating:  log.Rating,
				},
			},
		}

		if log.Shots > 0 {
			player.Stats.Core.ShootingPercentage = float64(log.Goals) / float64(log.Shots)
		}

		if log.TeamGoals > 0 {
			player.Stats.Core.GoalParticipation = float64(log.Goals+log.Assists) / float64(log.TeamGoals)
		}

		players = append(players, player)
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
	t, err := h.Octane.Teams().FindOne(bson.M{"name": name})
	if err == nil {
		team := t.(octane.Team)
		return &octane.Team{
			ID:   team.ID,
			Name: team.Name,
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
	p, err := h.Octane.Players().FindOne(bson.M{"tag": tag})
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
