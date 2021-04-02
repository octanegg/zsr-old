package octane

import (
	"time"

	"github.com/octanegg/zsr/ballchasing"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Event .
type Event struct {
	ID        *primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string              `json:"name,omitempty" bson:"name,omitempty"`
	StartDate *time.Time          `json:"startDate,omitempty" bson:"start_date,omitempty"`
	EndDate   *time.Time          `json:"endDate,omitempty" bson:"end_date,omitempty"`
	Region    string              `json:"region,omitempty" bson:"region,omitempty"`
	Mode      int                 `json:"mode,omitempty" bson:"mode,omitempty"`
	Prize     *Prize              `json:"prize,omitempty" bson:"prize,omitempty"`
	Tier      string              `json:"tier,omitempty" bson:"tier,omitempty"`
	Image     string              `json:"image,omitempty" bson:"image,omitempty"`
	Groups    []string            `json:"groups,omitempty" bson:"groups,omitempty"`
	Stages    []*Stage            `json:"stages,omitempty" bson:"stages,omitempty"`
}

// Stage .
type Stage struct {
	ID         int         `json:"_id" bson:"_id"`
	Name       string      `json:"name,omitempty" bson:"name,omitempty"`
	Format     string      `json:"format,omitempty" bson:"format,omitempty"`
	Region     string      `json:"region,omitempty" bson:"region,omitempty"`
	StartDate  *time.Time  `json:"startDate,omitempty" bson:"start_date,omitempty"`
	EndDate    *time.Time  `json:"endDate,omitempty" bson:"end_date,omitempty"`
	Prize      *Prize      `json:"prize,omitempty" bson:"prize,omitempty"`
	Liquipedia string      `json:"liquipedia,omitempty" bson:"liquipedia,omitempty"`
	Qualifier  bool        `json:"qualifier,omitempty" bson:"qualifier,omitempty"`
	Substages  []*Substage `json:"substages,omitempty" bson:"substages,omitempty"`
}

// Substage .
type Substage struct {
	ID     int    `json:"_id" bson:"_id"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Format string `json:"format,omitempty" bson:"format,omitempty"`
}

// Prize .
type Prize struct {
	Amount   float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency string  `json:"currency,omitempty" bson:"currency,omitempty"`
}

// Match .
type Match struct {
	ID                  *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OctaneID            string              `json:"octane_id,omitempty" bson:"octane_id,omitempty"`
	Event               *Event              `json:"event,omitempty" bson:"event,omitempty"`
	Stage               *Stage              `json:"stage,omitempty" bson:"stage,omitempty"`
	Substage            int                 `json:"substage,omitempty" bson:"substage,omitempty"`
	Date                *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Format              *Format             `json:"format,omitempty" bson:"format,omitempty"`
	Blue                *MatchSide          `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange              *MatchSide          `json:"orange,omitempty" bson:"orange,omitempty"`
	Number              int                 `json:"number,omitempty" bson:"number,omitempty"`
	ReverseSweep        bool                `json:"reverseSweep,omitempty" bson:"reverse_sweep,omitempty"`
	ReverseSweepAttempt bool                `json:"reverseSweepAttempt,omitempty" bson:"reverse_sweep_attempt,omitempty"`
	Games               []*GameOverview     `json:"games,omitempty" bson:"games,omitempty"`
}

// GameOverview .
type GameOverview struct {
	ID            *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Blue          float64             `json:"blue" bson:"blue"`
	Orange        float64             `json:"orange" bson:"orange"`
	Duration      int                 `json:"duration,omitempty" bson:"duration,omitempty"`
	BallchasingID string              `json:"ballchasing,omitempty" bson:"ballchasing,omitempty"`
}

// Format .
type Format struct {
	Type   string `json:"type,omitempty" bson:"type,omitempty"`
	Length int    `json:"length,omitempty" bson:"length,omitempty"`
}

// MatchSide .
type MatchSide struct {
	Score   int            `json:"score,omitempty" bson:"score,omitempty"`
	Winner  bool           `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *TeamStats     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*PlayerStats `json:"players,omitempty" bson:"players,omitempty"`
}

// Game .
type Game struct {
	ID            *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OctaneID      string              `json:"octane_id,omitempty" bson:"octane_id,omitempty"`
	Number        int                 `json:"number,omitempty" bson:"number,omitempty"`
	Match         *Match              `json:"match,omitempty" bson:"match,omitempty"`
	Map           *Map                `json:"map,omitempty" bson:"map,omitempty"`
	Duration      int                 `json:"duration,omitempty" bson:"duration,omitempty"`
	Date          *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Blue          *GameSide           `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange        *GameSide           `json:"orange,omitempty" bson:"orange,omitempty"`
	BallchasingID string              `json:"ballchasing,omitempty" bson:"ballchasing,omitempty"`
}

// GameSide .
type GameSide struct {
	Winner  bool           `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *TeamStats     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*PlayerStats `json:"players,omitempty" bson:"players,omitempty"`
}

// PlayerStats .
type PlayerStats struct {
	Player   *Player                  `json:"player,omitempty" bson:"player,omitempty"`
	Car      *Car                     `json:"car,omitempty" bson:"car,omitempty"`
	Camera   *ballchasing.Camera      `json:"camera,omitempty" bson:"camera,omitempty"`
	Stats    *ballchasing.PlayerStats `json:"stats,omitempty" bson:"stats,omitempty"`
	Advanced *AdvancedStats           `json:"advanced,omitempty" bson:"advanced,omitempty"`
}

// AdvancedStats .
type AdvancedStats struct {
	GoalParticipation float64 `json:"goalParticipation" bson:"goal_participation"`
	Rating            float64 `json:"rating,omitempty" bson:"rating,omitempty"`
	MVP               bool    `json:"mvp,omitempty" bson:"mvp,omitempty"`
}

// TeamStats .
type TeamStats struct {
	Team  *Team                  `json:"team,omitempty" bson:"team,omitempty"`
	Stats *ballchasing.TeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
}

// Map .
type Map struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

// Car .
type Car struct {
	ID   int    `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

// Player .
type Player struct {
	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tag      string              `json:"tag,omitempty" bson:"tag,omitempty"`
	Name     string              `json:"name,omitempty" bson:"name,omitempty"`
	Country  string              `json:"country,omitempty" bson:"country,omitempty"`
	Team     *Team               `json:"team,omitempty" bson:"team,omitempty"`
	Accounts []*Account          `json:"accounts,omitempty" bson:"accounts,omitempty"`
}

// Account .
type Account struct {
	Platform string `json:"platform,omitempty" bson:"platform,omitempty"`
	ID       string `json:"id,omitempty" bson:"id,omitempty"`
}

// Team .
type Team struct {
	ID     *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string              `json:"name,omitempty" bson:"name,omitempty"`
	Region string              `json:"region,omitempty" bson:"region,omitempty"`
	Image  string              `json:"image,omitempty" bson:"image,omitempty"`
}

// Statline .
type Statline struct {
	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Game     *Game               `json:"game,omitempty" bson:"game,omitempty"`
	Team     *StatlineSide       `json:"team,omitempty" bson:"team,omitempty"`
	Opponent *StatlineSide       `json:"opponent,omitempty" bson:"opponent,omitempty"`
	Player   *PlayerStats        `json:"player,omitempty" bson:"player,omitempty"`
}

// StatlineSide .
type StatlineSide struct {
	Score   float64                `json:"score,omitempty" bson:"score,omitempty"`
	Winner  bool                   `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *Team                  `json:"team,omitempty" bson:"team,omitempty"`
	Stats   *ballchasing.TeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
	Players []*Player              `json:"players,omitempty" bson:"players,omitempty"`
}

type Participant struct {
	Team    *Team     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*Player `json:"players,omitempty" bson:"players,omitempty"`
}

func toEvents(cursor *mongo.Cursor) (interface{}, error) {
	var event Event
	if err := cursor.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}

func toMatches(cursor *mongo.Cursor) (interface{}, error) {
	var match Match
	if err := cursor.Decode(&match); err != nil {
		return nil, err
	}
	return match, nil
}

func toGames(cursor *mongo.Cursor) (interface{}, error) {
	var game Game
	if err := cursor.Decode(&game); err != nil {
		return nil, err
	}
	return game, nil
}

func toPlayers(cursor *mongo.Cursor) (interface{}, error) {
	var player Player
	if err := cursor.Decode(&player); err != nil {
		return nil, err
	}
	return player, nil
}

func toTeams(cursor *mongo.Cursor) (interface{}, error) {
	var team Team
	if err := cursor.Decode(&team); err != nil {
		return nil, err
	}
	return team, nil
}

func toStatlines(cursor *mongo.Cursor) (interface{}, error) {
	var statline Statline
	if err := cursor.Decode(&statline); err != nil {
		return nil, err
	}
	return statline, nil
}
