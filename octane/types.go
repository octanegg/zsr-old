package octane

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Event .
type Event struct {
	ID        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
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
	Location   *Location   `json:"location,omitempty" bson:"location,omitempty"`
	Substages  []*Substage `json:"substages,omitempty" bson:"substages,omitempty"`
}

// Location .
type Location struct {
	Venue   string `json:"venue,omitempty" bson:"venue,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
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
	Score   int           `json:"score,omitempty" bson:"score,omitempty"`
	Winner  bool          `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *TeamInfo     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*PlayerInfo `json:"players,omitempty" bson:"players,omitempty"`
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
	Winner  bool          `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *TeamInfo     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*PlayerInfo `json:"players,omitempty" bson:"players,omitempty"`
}

// PlayerInfo .
type PlayerInfo struct {
	Player   *Player        `json:"player,omitempty" bson:"player,omitempty"`
	Car      *Car           `json:"car,omitempty" bson:"car,omitempty"`
	Camera   *Camera        `json:"camera,omitempty" bson:"camera,omitempty"`
	Stats    *PlayerStats   `json:"stats,omitempty" bson:"stats,omitempty"`
	Advanced *AdvancedStats `json:"advanced,omitempty" bson:"advanced,omitempty"`
}

// AdvancedStats .
type AdvancedStats struct {
	GoalParticipation float64 `json:"goalParticipation" bson:"goal_participation"`
	Rating            float64 `json:"rating,omitempty" bson:"rating,omitempty"`
	MVP               bool    `json:"mvp,omitempty" bson:"mvp,omitempty"`
}

// TeamInfo .
type TeamInfo struct {
	Team  *Team      `json:"team,omitempty" bson:"team,omitempty"`
	Stats *TeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
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
	ID         *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Tag        string              `json:"tag,omitempty" bson:"tag,omitempty"`
	Name       string              `json:"name,omitempty" bson:"name,omitempty"`
	Country    string              `json:"country,omitempty" bson:"country,omitempty"`
	Team       *Team               `json:"team,omitempty" bson:"team,omitempty"`
	Accounts   []*Account          `json:"accounts,omitempty" bson:"accounts,omitempty"`
	Substitute bool                `json:"substitute,omitempty" bson:"substitute,omitempty"`
	Coach      bool                `json:"coach,omitempty" bson:"coach,omitempty"`
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
	Player   *PlayerInfo         `json:"player,omitempty" bson:"player,omitempty"`
}

// StatlineSide .
type StatlineSide struct {
	Score   float64    `json:"score,omitempty" bson:"score,omitempty"`
	Winner  bool       `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *Team      `json:"team,omitempty" bson:"team,omitempty"`
	Stats   *TeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
	Players []*Player  `json:"players,omitempty" bson:"players,omitempty"`
}

type Participant struct {
	Team    *Team     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*Player `json:"players,omitempty" bson:"players,omitempty"`
}

// Camera .
type Camera struct {
	Fov             int     `json:"fov" bson:"fov"`
	Height          int     `json:"height" bson:"height"`
	Pitch           int     `json:"pitch" bson:"pitch"`
	Distance        int     `json:"distance" bson:"distance"`
	Stiffness       float64 `json:"stiffness" bson:"stiffness"`
	SwivelSpeed     float64 `json:"swivelSpeed" bson:"swivel_speed"`
	TransitionSpeed float64 `json:"transitionSpeed" bson:"transition_speed"`
}

// PlayerStats .
type PlayerStats struct {
	Core        *PlayerCore        `json:"core,omitempty" bson:"core,omitempty"`
	Boost       *PlayerBoost       `json:"boost,omitempty" bson:"boost,omitempty"`
	Movement    *PlayerMovement    `json:"movement,omitempty" bson:"movement,omitempty"`
	Positioning *PlayerPositioning `json:"positioning,omitempty" bson:"positioning,omitempty"`
	Demolitions *PlayerDemolitions `json:"demo,omitempty" bson:"demo,omitempty"`
}

// PlayerCore .
type PlayerCore struct {
	Shots              float64 `json:"shots" bson:"shots"`
	Goals              float64 `json:"goals" bson:"goals"`
	Saves              float64 `json:"saves" bson:"saves"`
	Assists            float64 `json:"assists" bson:"assists"`
	Score              float64 `json:"score" bson:"score"`
	ShootingPercentage float64 `json:"shootingPercentage" bson:"shooting_percentage"`
}

// PlayerBoost .
type PlayerBoost struct {
	Bpm                       float64 `json:"bpm" bson:"bpm"`
	Bcpm                      float64 `json:"bcpm" bson:"bcpm"`
	AvgAmount                 float64 `json:"avgAmount" bson:"avg_amount"`
	AmountCollected           float64 `json:"amountCollected" bson:"amount_collected"`
	AmountStolen              float64 `json:"amountStolen" bson:"amount_stolen"`
	AmountCollectedBig        float64 `json:"amountCollectedBig" bson:"amount_collected_big"`
	AmountStolenBig           float64 `json:"amountStolenBig" bson:"amount_stolen_big"`
	AmountCollectedSmall      float64 `json:"amountCollectedSmall" bson:"amount_collected_small"`
	AmountStolenSmall         float64 `json:"amountStolenSmall" bson:"amount_stolen_small"`
	CountCollectedBig         float64 `json:"countCollectedBig" bson:"count_collected_big"`
	CountStolenBig            float64 `json:"countStolenBig" bson:"count_stolen_big"`
	CountCollectedSmall       float64 `json:"countCollectedSmall" bson:"count_collected_small"`
	CountStolenSmall          float64 `json:"countStolenSmall" bson:"count_stolen_small"`
	AmountOverfill            float64 `json:"amountOverfill" bson:"amount_overfill"`
	AmountOverfillStolen      float64 `json:"amountOverfillStolen" bson:"amount_overfill_stolen"`
	AmountUsedWhileSupersonic float64 `json:"amountUsedWhileSupersonic" bson:"amount_used_while_supersonic"`
	TimeZeroBoost             float64 `json:"timeZeroBoost" bson:"time_zero_boost"`
	PercentZeroBoost          float64 `json:"percentZeroBoost" bson:"percent_zero_boost"`
	TimeFullBoost             float64 `json:"timeFullBoost" bson:"time_full_boost"`
	PercentFullBoost          float64 `json:"percentFullBoost" bson:"percent_full_boost"`
	TimeBoost025              float64 `json:"timeBoost0To25" bson:"time_boost_0_25"`
	TimeBoost2550             float64 `json:"timeBoost25To50" bson:"time_boost_25_50"`
	TimeBoost5075             float64 `json:"timeBoost50To75" bson:"time_boost_50_75"`
	TimeBoost75100            float64 `json:"timeBoost75To100" bson:"time_boost_75_100"`
	PercentBoost025           float64 `json:"percentBoost0To25" bson:"percent_boost_0_25"`
	PercentBoost2550          float64 `json:"percentBoost25To50" bson:"percent_boost_25_50"`
	PercentBoost5075          float64 `json:"percentBoost50To75" bson:"percent_boost_50_75"`
	PercentBoost75100         float64 `json:"percentBoost75To100" bson:"percent_boost_75_100"`
}

// PlayerMovement .
type PlayerMovement struct {
	AvgSpeed               float64 `json:"avgSpeed" bson:"avg_speed"`
	TotalDistance          float64 `json:"totalDistance" bson:"total_distance"`
	TimeSupersonicSpeed    float64 `json:"timeSupersonicSpeed" bson:"time_supersonic_speed"`
	TimeBoostSpeed         float64 `json:"timeBoostSpeed" bson:"time_boost_speed"`
	TimeSlowSpeed          float64 `json:"timeSlowSpeed" bson:"time_slow_speed"`
	TimeGround             float64 `json:"timeGround" bson:"time_ground"`
	TimeLowAir             float64 `json:"timeLowAir" bson:"time_low_air"`
	TimeHighAir            float64 `json:"timeHighAir" bson:"time_high_air"`
	TimePowerslide         float64 `json:"timePowerslide" bson:"time_powerslide"`
	CountPowerslide        float64 `json:"countPowerslide" bson:"count_powerslide"`
	AvgPowerslideDuration  float64 `json:"avgPowerslideDuration" bson:"avg_powerslide_duration"`
	AvgSpeedPercentage     float64 `json:"avgSpeedPercentage" bson:"avg_speed_percentage"`
	PercentSlowSpeed       float64 `json:"percentSlowSpeed" bson:"percent_slow_speed"`
	PercentBoostSpeed      float64 `json:"percentBoostSpeed" bson:"percent_boost_speed"`
	PercentSupersonicSpeed float64 `json:"percentSupersonicSpeed" bson:"percent_supersonic_speed"`
	PercentGround          float64 `json:"percentGround" bson:"percent_ground"`
	PercentLowAir          float64 `json:"percentLowAir" bson:"percent_low_air"`
	PercentHighAir         float64 `json:"percentHighAir" bson:"percent_high_air"`
}

// PlayerPositioning .
type PlayerPositioning struct {
	AvgDistanceToBall             float64 `json:"avgDistanceToBall" bson:"avg_distance_to_ball"`
	AvgDistanceToBallPossession   float64 `json:"avgDistanceToBallPossession" bson:"avg_distance_to_ball_possession"`
	AvgDistanceToBallNoPossession float64 `json:"avgDistanceToBallNoPossession" bson:"avg_distance_to_ball_no_possession"`
	AvgDistanceToMates            float64 `json:"avgDistanceToMates" bson:"avg_distance_to_mates"`
	TimeDefensiveThird            float64 `json:"timeDefensiveThird" bson:"time_defensive_third"`
	TimeNeutralThird              float64 `json:"timeNeutralThird" bson:"time_neutral_third"`
	TimeOffensiveThird            float64 `json:"timeOffensiveThird" bson:"time_offensive_third"`
	TimeDefensiveHalf             float64 `json:"timeDefensiveHalf" bson:"time_defensive_half"`
	TimeOffensiveHalf             float64 `json:"timeOffensiveHalf" bson:"time_offensive_half"`
	TimeBehindBall                float64 `json:"timeBehindBall" bson:"time_behind_ball"`
	TimeInfrontBall               float64 `json:"timeInfrontBall" bson:"time_infront_ball"`
	TimeMostBack                  float64 `json:"timeMostBack" bson:"time_most_back"`
	TimeMostForward               float64 `json:"timeMostForward" bson:"time_most_forward"`
	GoalsAgainstWhileLastDefender float64 `json:"goalsAgainstWhileLastDefender" bson:"goals_against_while_last_defender"`
	TimeClosestToBall             float64 `json:"timeClosestToBall" bson:"time_closest_to_ball"`
	TimeFarthestFromBall          float64 `json:"timeFarthestFromBall" bson:"time_farthest_from_ball"`
	PercentDefensiveThird         float64 `json:"percentDefensiveThird" bson:"percent_defensive_third"`
	PercentOffensiveThird         float64 `json:"percentOffensiveThird" bson:"percent_offensive_third"`
	PercentNeutralThird           float64 `json:"percentNeutralThird" bson:"percent_neutral_third"`
	PercentDefensiveHalf          float64 `json:"percentDefensiveHalf" bson:"percent_defensive_half"`
	PercentOffensiveHalf          float64 `json:"percentOffensiveHalf" bson:"percent_offensive_half"`
	PercentBehindBall             float64 `json:"percentBehindBall" bson:"percent_behind_ball"`
	PercentInfrontBall            float64 `json:"percentInfrontBall" bson:"percent_infront_ball"`
	PercentMostBack               float64 `json:"percentMostBack" bson:"percent_most_back"`
	PercentMostForward            float64 `json:"percentMostForward" bson:"percent_most_forward"`
	PercentClosestToBall          float64 `json:"percentClosestToBall" bson:"percent_closest_to_ball"`
	PercentFarthestFromBall       float64 `json:"percentFarthestFromBall" bson:"percent_farthest_from_ball"`
}

// PlayerDemolitions .
type PlayerDemolitions struct {
	Inflicted float64 `json:"inflicted" bson:"inflicted"`
	Taken     float64 `json:"taken" bson:"taken"`
}

// TeamStats .
type TeamStats struct {
	Ball        *TeamBall        `json:"ball,omitempty" bson:"ball,omitempty"`
	Core        *TeamCore        `json:"core,omitempty" bson:"core,omitempty"`
	Boost       *TeamBoost       `json:"boost,omitempty" bson:"boost,omitempty"`
	Movement    *TeamMovement    `json:"movement,omitempty" bson:"movement,omitempty"`
	Positioning *TeamPositioning `json:"positioning,omitempty" bson:"positioning,omitempty"`
	Demolitions *TeamDemolitions `json:"demo,omitempty" bson:"demo,omitempty"`
}

// TeamBall .
type TeamBall struct {
	PossessionTime float64 `json:"possessionTime" bson:"possession_time"`
	TimeInSide     float64 `json:"timeInSide" bson:"time_in_side"`
}

// TeamCore .
type TeamCore struct {
	Shots              float64 `json:"shots" bson:"shots"`
	Goals              float64 `json:"goals" bson:"goals"`
	Saves              float64 `json:"saves" bson:"saves"`
	Assists            float64 `json:"assists" bson:"assists"`
	Score              float64 `json:"score" bson:"score"`
	ShootingPercentage float64 `json:"shootingPercentage" bson:"shooting_percentage"`
}

// TeamBoost .
type TeamBoost struct {
	Bpm                       float64 `json:"bpm" bson:"bpm"`
	Bcpm                      float64 `json:"bcpm" bson:"bcpm"`
	AvgAmount                 float64 `json:"avgAmount" bson:"avg_amount"`
	AmountCollected           float64 `json:"amountCollected" bson:"amount_collected"`
	AmountStolen              float64 `json:"amountStolen" bson:"amount_stolen"`
	AmountCollectedBig        float64 `json:"amountCollectedBig" bson:"amount_collected_big"`
	AmountStolenBig           float64 `json:"amountStolenBig" bson:"amount_stolen_big"`
	AmountCollectedSmall      float64 `json:"amountCollectedSmall" bson:"amount_collected_small"`
	AmountStolenSmall         float64 `json:"amountStolenSmall" bson:"amount_stolen_small"`
	CountCollectedBig         float64 `json:"countCollectedBig" bson:"count_collected_big"`
	CountStolenBig            float64 `json:"countStolenBig" bson:"count_stolen_big"`
	CountCollectedSmall       float64 `json:"countCollectedSmall" bson:"count_collected_small"`
	CountStolenSmall          float64 `json:"countStolenSmall" bson:"count_stolen_small"`
	AmountOverfill            float64 `json:"amountOverfill" bson:"amount_overfill"`
	AmountOverfillStolen      float64 `json:"amountOverfillStolen" bson:"amount_overfill_stolen"`
	AmountUsedWhileSupersonic float64 `json:"amountUsedWhileSupersonic" bson:"amount_used_while_supersonic"`
	TimeZeroBoost             float64 `json:"timeZeroBoost" bson:"time_zero_boost"`
	TimeFullBoost             float64 `json:"timeFullBoost" bson:"time_full_boost"`
	TimeBoost025              float64 `json:"timeBoost0To25" bson:"time_boost_0_25"`
	TimeBoost2550             float64 `json:"timeBoost25To50" bson:"time_boost_25_50"`
	TimeBoost5075             float64 `json:"timeBoost50To75" bson:"time_boost_50_75"`
	TimeBoost75100            float64 `json:"timeBoost75To100" bson:"time_boost_75_100"`
}

// TeamMovement .
type TeamMovement struct {
	TotalDistance       float64 `json:"totalDistance" bson:"total_distance"`
	TimeSupersonicSpeed float64 `json:"timeSupersonicSpeed" bson:"time_supersonic_speed"`
	TimeBoostSpeed      float64 `json:"timeBoostSpeed" bson:"time_boost_speed"`
	TimeSlowSpeed       float64 `json:"timeSlowSpeed" bson:"time_slow_speed"`
	TimeGround          float64 `json:"timeGround" bson:"time_ground"`
	TimeLowAir          float64 `json:"timeLowAir" bson:"time_low_air"`
	TimeHighAir         float64 `json:"timeHighAir" bson:"time_high_air"`
	TimePowerslide      float64 `json:"timePowerslide" bson:"time_powerslide"`
	CountPowerslide     float64 `json:"countPowerslide" bson:"count_powerslide"`
}

// TeamPositioning .
type TeamPositioning struct {
	TimeDefensiveThird float64 `json:"timeDefensiveThird" bson:"time_defensive_third"`
	TimeNeutralThird   float64 `json:"timeNeutralThird" bson:"time_neutral_third"`
	TimeOffensiveThird float64 `json:"timeOffensiveThird" bson:"time_offensive_third"`
	TimeDefensiveHalf  float64 `json:"timeDefensiveHalf" bson:"time_defensive_half"`
	TimeOffensiveHalf  float64 `json:"timeOffensiveHalf" bson:"time_offensive_half"`
	TimeBehindBall     float64 `json:"timeBehindBall" bson:"time_behind_ball"`
	TimeInfrontBall    float64 `json:"timeInfrontBall" bson:"time_infront_ball"`
}

// TeamDemolitions .
type TeamDemolitions struct {
	Inflicted float64 `json:"inflicted" bson:"inflicted"`
	Taken     float64 `json:"taken" bson:"taken"`
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
