package ballchasing

import "time"

// Replays .
type Replays struct {
	Count int      `json:"count" bson:"count"`
	List  []Replay `json:"list" bson:"list"`
}

// Replay .
type Replay struct {
	ID              string    `json:"id" bson:"id"`
	Link            string    `json:"link" bson:"link"`
	Created         time.Time `json:"created" bson:"created"`
	Uploader        Uploader  `json:"uploader" bson:"uploader"`
	Status          string    `json:"status" bson:"status"`
	RocketLeagueID  string    `json:"rocket_league_id" bson:"rocket_league_id"`
	MatchGUID       string    `json:"match_guid" bson:"match_guid"`
	Title           string    `json:"title" bson:"title"`
	MapCode         string    `json:"map_code" bson:"map_code"`
	MatchType       string    `json:"match_type" bson:"match_type"`
	TeamSize        int       `json:"team_size" bson:"team_size"`
	PlaylistID      string    `json:"playlist_id" bson:"playlist_id"`
	Duration        int       `json:"duration" bson:"duration"`
	Overtime        bool      `json:"overtime" bson:"overtime"`
	OvertimeSeconds int       `json:"overtime_seconds" bson:"overtime_seconds"`
	Season          int       `json:"season" bson:"season"`
	Date            time.Time `json:"date" bson:"date"`
	Visibility      string    `json:"visibility" bson:"visibility"`
	Groups          []Group   `json:"groups" bson:"groups"`
	Blue            Team      `json:"blue" bson:"blue"`
	Orange          Team      `json:"orange" bson:"orange"`
	PlaylistName    string    `json:"playlist_name" bson:"playlist_name"`
	MapName         string    `json:"map_name" bson:"map_name"`
}

// Uploader .
type Uploader struct {
	SteamID    string `json:"steam_id" bson:"steam_id"`
	Name       string `json:"name" bson:"name"`
	ProfileURL string `json:"profile_url" bson:"profile_url"`
	Avatar     string `json:"avatar" bson:"avatar"`
}

// Group .
type Group struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Link string `json:"link" bson:"link"`
}

// Team .
type Team struct {
	Color   string    `json:"color" bson:"color"`
	Name    string    `json:"name" bson:"name"`
	Players []Player  `json:"players" bson:"players"`
	Stats   TeamStats `json:"stats" bson:"stats"`
}

// Player .
type Player struct {
	StartTime           float64     `json:"start_time" bson:"start_time"`
	EndTime             float64     `json:"end_time" bson:"end_time"`
	Name                string      `json:"name" bson:"name"`
	ID                  ID          `json:"id" bson:"id"`
	Mvp                 bool        `json:"mvp,omitempty" bson:"mvp,omitempty"`
	CarID               int         `json:"car_id" bson:"car_id"`
	CarName             string      `json:"car_name" bson:"car_name"`
	Camera              Camera      `json:"camera" bson:"camera"`
	SteeringSensitivity float64     `json:"steering_sensitivity" bson:"steering_sensitivity"`
	Stats               PlayerStats `json:"stats" bson:"stats"`
}

// ID .
type ID struct {
	Platform string `json:"platform" bson:"platform"`
	ID       string `json:"id" bson:"id"`
}

// Camera .
type Camera struct {
	Fov             int     `json:"fov" bson:"fov"`
	Height          int     `json:"height" bson:"height"`
	Pitch           int     `json:"pitch" bson:"pitch"`
	Distance        int     `json:"distance" bson:"distance"`
	Stiffness       float64 `json:"stiffness" bson:"stiffness"`
	SwivelSpeed     float64 `json:"swivel_speed" bson:"swivel_speed"`
	TransitionSpeed float64 `json:"transition_speed" bson:"transition_speed"`
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
	Shots              int     `json:"shots" bson:"shots"`
	Goals              int     `json:"goals" bson:"goals"`
	Saves              int     `json:"saves" bson:"saves"`
	Assists            int     `json:"assists" bson:"assists"`
	Score              int     `json:"score" bson:"score"`
	ShootingPercentage float64 `json:"shooting_percentage" bson:"shooting_percentage"`
	GoalParticipation  float64 `json:"goal_participation" bson:"goal_participation"`
	Rating             float64 `json:"rating,omitempty" bson:"rating,omitempty"`
	Mvp                bool    `json:"mvp,omitempty" bson:"mvp,omitempty"`
}

// PlayerBoost .
type PlayerBoost struct {
	Bpm                       int     `json:"bpm" bson:"bpm"`
	Bcpm                      float64 `json:"bcpm" bson:"bcpm"`
	AvgAmount                 float64 `json:"avg_amount" bson:"avg_amount"`
	AmountCollected           int     `json:"amount_collected" bson:"amount_collected"`
	AmountStolen              int     `json:"amount_stolen" bson:"amount_stolen"`
	AmountCollectedBig        int     `json:"amount_collected_big" bson:"amount_collected_big"`
	AmountStolenBig           int     `json:"amount_stolen_big" bson:"amount_stolen_big"`
	AmountCollectedSmall      int     `json:"amount_collected_small" bson:"amount_collected_small"`
	AmountStolenSmall         int     `json:"amount_stolen_small" bson:"amount_stolen_small"`
	CountCollectedBig         int     `json:"count_collected_big" bson:"count_collected_big"`
	CountStolenBig            int     `json:"count_stolen_big" bson:"count_stolen_big"`
	CountCollectedSmall       int     `json:"count_collected_small" bson:"count_collected_small"`
	CountStolenSmall          int     `json:"count_stolen_small" bson:"count_stolen_small"`
	AmountOverfill            int     `json:"amount_overfill" bson:"amount_overfill"`
	AmountOverfillStolen      int     `json:"amount_overfill_stolen" bson:"amount_overfill_stolen"`
	AmountUsedWhileSupersonic int     `json:"amount_used_while_supersonic" bson:"amount_used_while_supersonic"`
	TimeZeroBoost             float64 `json:"time_zero_boost" bson:"time_zero_boost"`
	PercentZeroBoost          float64 `json:"percent_zero_boost" bson:"percent_zero_boost"`
	TimeFullBoost             float64 `json:"time_full_boost" bson:"time_full_boost"`
	PercentFullBoost          float64 `json:"percent_full_boost" bson:"percent_full_boost"`
	TimeBoost025              float64 `json:"time_boost_0_25" bson:"time_boost_0_25"`
	TimeBoost2550             float64 `json:"time_boost_25_50" bson:"time_boost_25_50"`
	TimeBoost5075             float64 `json:"time_boost_50_75" bson:"time_boost_50_75"`
	TimeBoost75100            float64 `json:"time_boost_75_100" bson:"time_boost_75_100"`
	PercentBoost025           float64 `json:"percent_boost_0_25" bson:"percent_boost_0_25"`
	PercentBoost2550          float64 `json:"percent_boost_25_50" bson:"percent_boost_25_50"`
	PercentBoost5075          float64 `json:"percent_boost_50_75" bson:"percent_boost_50_75"`
	PercentBoost75100         float64 `json:"percent_boost_75_100" bson:"percent_boost_75_100"`
}

// PlayerMovement .
type PlayerMovement struct {
	AvgSpeed               int     `json:"avg_speed" bson:"avg_speed"`
	TotalDistance          int     `json:"total_distance" bson:"total_distance"`
	TimeSupersonicSpeed    float64 `json:"time_supersonic_speed" bson:"time_supersonic_speed"`
	TimeBoostSpeed         float64 `json:"time_boost_speed" bson:"time_boost_speed"`
	TimeSlowSpeed          float64 `json:"time_slow_speed" bson:"time_slow_speed"`
	TimeGround             float64 `json:"time_ground" bson:"time_ground"`
	TimeLowAir             float64 `json:"time_low_air" bson:"time_low_air"`
	TimeHighAir            float64 `json:"time_high_air" bson:"time_high_air"`
	TimePowerslide         float64 `json:"time_powerslide" bson:"time_powerslide"`
	CountPowerslide        int     `json:"count_powerslide" bson:"count_powerslide"`
	AvgPowerslideDuration  float64 `json:"avg_powerslide_duration" bson:"avg_powerslide_duration"`
	AvgSpeedPercentage     float64 `json:"avg_speed_percentage" bson:"avg_speed_percentage"`
	PercentSlowSpeed       float64 `json:"percent_slow_speed" bson:"percent_slow_speed"`
	PercentBoostSpeed      float64 `json:"percent_boost_speed" bson:"percent_boost_speed"`
	PercentSupersonicSpeed float64 `json:"percent_supersonic_speed" bson:"percent_supersonic_speed"`
	PercentGround          float64 `json:"percent_ground" bson:"percent_ground"`
	PercentLowAir          float64 `json:"percent_low_air" bson:"percent_low_air"`
	PercentHighAir         float64 `json:"percent_high_air" bson:"percent_high_air"`
}

// PlayerPositioning .
type PlayerPositioning struct {
	AvgDistanceToBall             int     `json:"avg_distance_to_ball" bson:"avg_distance_to_ball"`
	AvgDistanceToBallPossession   int     `json:"avg_distance_to_ball_possession" bson:"avg_distance_to_ball_possession"`
	AvgDistanceToBallNoPossession int     `json:"avg_distance_to_ball_no_possession" bson:"avg_distance_to_ball_no_possession"`
	AvgDistanceToMates            int     `json:"avg_distance_to_mates" bson:"avg_distance_to_mates"`
	TimeDefensiveThird            float64 `json:"time_defensive_third" bson:"time_defensive_third"`
	TimeNeutralThird              float64 `json:"time_neutral_third" bson:"time_neutral_third"`
	TimeOffensiveThird            float64 `json:"time_offensive_third" bson:"time_offensive_third"`
	TimeDefensiveHalf             float64 `json:"time_defensive_half" bson:"time_defensive_half"`
	TimeOffensiveHalf             float64 `json:"time_offensive_half" bson:"time_offensive_half"`
	TimeBehindBall                float64 `json:"time_behind_ball" bson:"time_behind_ball"`
	TimeInfrontBall               float64 `json:"time_infront_ball" bson:"time_infront_ball"`
	TimeMostBack                  float64 `json:"time_most_back" bson:"time_most_back"`
	TimeMostForward               float64 `json:"time_most_forward" bson:"time_most_forward"`
	GoalsAgainstWhileLastDefender int     `json:"goals_against_while_last_defender" bson:"goals_against_while_last_defender"`
	TimeClosestToBall             float64 `json:"time_closest_to_ball" bson:"time_closest_to_ball"`
	TimeFarthestFromBall          float64 `json:"time_farthest_from_ball" bson:"time_farthest_from_ball"`
	PercentDefensiveThird         float64 `json:"percent_defensive_third" bson:"percent_defensive_third"`
	PercentOffensiveThird         float64 `json:"percent_offensive_third" bson:"percent_offensive_third"`
	PercentNeutralThird           float64 `json:"percent_neutral_third" bson:"percent_neutral_third"`
	PercentDefensiveHalf          float64 `json:"percent_defensive_half" bson:"percent_defensive_half"`
	PercentOffensiveHalf          float64 `json:"percent_offensive_half" bson:"percent_offensive_half"`
	PercentBehindBall             float64 `json:"percent_behind_ball" bson:"percent_behind_ball"`
	PercentInfrontBall            float64 `json:"percent_infront_ball" bson:"percent_infront_ball"`
	PercentMostBack               float64 `json:"percent_most_back" bson:"percent_most_back"`
	PercentMostForward            float64 `json:"percent_most_forward" bson:"percent_most_forward"`
	PercentClosestToBall          float64 `json:"percent_closest_to_ball" bson:"percent_closest_to_ball"`
	PercentFarthestFromBall       float64 `json:"percent_farthest_from_ball" bson:"percent_farthest_from_ball"`
}

// PlayerDemolitions .
type PlayerDemolitions struct {
	Inflicted int `json:"inflicted" bson:"inflicted"`
	Taken     int `json:"taken" bson:"taken"`
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
	PossessionTime float64 `json:"possession_time" bson:"possession_time"`
	TimeInSide     float64 `json:"time_in_side" bson:"time_in_side"`
}

// TeamCore .
type TeamCore struct {
	Shots              int     `json:"shots" bson:"shots"`
	Goals              int     `json:"goals" bson:"goals"`
	Saves              int     `json:"saves" bson:"saves"`
	Assists            int     `json:"assists" bson:"assists"`
	Score              int     `json:"score" bson:"score"`
	ShootingPercentage float64 `json:"shooting_percentage" bson:"shooting_percentage"`
}

// TeamBoost .
type TeamBoost struct {
	Bpm                       int     `json:"bpm" bson:"bpm"`
	Bcpm                      float64 `json:"bcpm" bson:"bcpm"`
	AvgAmount                 float64 `json:"avg_amount" bson:"avg_amount"`
	AmountCollected           int     `json:"amount_collected" bson:"amount_collected"`
	AmountStolen              int     `json:"amount_stolen" bson:"amount_stolen"`
	AmountCollectedBig        int     `json:"amount_collected_big" bson:"amount_collected_big"`
	AmountStolenBig           int     `json:"amount_stolen_big" bson:"amount_stolen_big"`
	AmountCollectedSmall      int     `json:"amount_collected_small" bson:"amount_collected_small"`
	AmountStolenSmall         int     `json:"amount_stolen_small" bson:"amount_stolen_small"`
	CountCollectedBig         int     `json:"count_collected_big" bson:"count_collected_big"`
	CountStolenBig            int     `json:"count_stolen_big" bson:"count_stolen_big"`
	CountCollectedSmall       int     `json:"count_collected_small" bson:"count_collected_small"`
	CountStolenSmall          int     `json:"count_stolen_small" bson:"count_stolen_small"`
	AmountOverfill            int     `json:"amount_overfill" bson:"amount_overfill"`
	AmountOverfillStolen      int     `json:"amount_overfill_stolen" bson:"amount_overfill_stolen"`
	AmountUsedWhileSupersonic int     `json:"amount_used_while_supersonic" bson:"amount_used_while_supersonic"`
	TimeZeroBoost             float64 `json:"time_zero_boost" bson:"time_zero_boost"`
	TimeFullBoost             float64 `json:"time_full_boost" bson:"time_full_boost"`
	TimeBoost025              float64 `json:"time_boost_0_25" bson:"time_boost_0_25"`
	TimeBoost2550             float64 `json:"time_boost_25_50" bson:"time_boost_25_50"`
	TimeBoost5075             float64 `json:"time_boost_50_75" bson:"time_boost_50_75"`
	TimeBoost75100            float64 `json:"time_boost_75_100" bson:"time_boost_75_100"`
}

// TeamMovement .
type TeamMovement struct {
	TotalDistance       int     `json:"total_distance" bson:"total_distance"`
	TimeSupersonicSpeed float64 `json:"time_supersonic_speed" bson:"time_supersonic_speed"`
	TimeBoostSpeed      float64 `json:"time_boost_speed" bson:"time_boost_speed"`
	TimeSlowSpeed       float64 `json:"time_slow_speed" bson:"time_slow_speed"`
	TimeGround          float64 `json:"time_ground" bson:"time_ground"`
	TimeLowAir          float64 `json:"time_low_air" bson:"time_low_air"`
	TimeHighAir         float64 `json:"time_high_air" bson:"time_high_air"`
	TimePowerslide      float64 `json:"time_powerslide" bson:"time_powerslide"`
	CountPowerslide     int     `json:"count_powerslide" bson:"count_powerslide"`
}

// TeamPositioning .
type TeamPositioning struct {
	TimeDefensiveThird float64 `json:"time_defensive_third" bson:"time_defensive_third"`
	TimeNeutralThird   float64 `json:"time_neutral_third" bson:"time_neutral_third"`
	TimeOffensiveThird float64 `json:"time_offensive_third" bson:"time_offensive_third"`
	TimeDefensiveHalf  float64 `json:"time_defensive_half" bson:"time_defensive_half"`
	TimeOffensiveHalf  float64 `json:"time_offensive_half" bson:"time_offensive_half"`
	TimeBehindBall     float64 `json:"time_behind_ball" bson:"time_behind_ball"`
	TimeInfrontBall    float64 `json:"time_infront_ball" bson:"time_infront_ball"`
}

// TeamDemolitions .
type TeamDemolitions struct {
	Inflicted int `json:"inflicted" bson:"inflicted"`
	Taken     int `json:"taken" bson:"taken"`
}
