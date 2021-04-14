package ballchasing

import "time"

// Replays .
type Replays struct {
	Count int      `json:"count"`
	List  []Replay `json:"list"`
}

// Replay .
type Replay struct {
	ID              string    `json:"id"`
	Link            string    `json:"link"`
	Created         time.Time `json:"created"`
	Uploader        Uploader  `json:"uploader"`
	Status          string    `json:"status"`
	RocketLeagueID  string    `json:"rocket_league_id"`
	MatchGUID       string    `json:"match_guid"`
	Title           string    `json:"title"`
	MapCode         string    `json:"map_code"`
	MatchType       string    `json:"match_type"`
	TeamSize        int       `json:"team_size"`
	PlaylistID      string    `json:"playlist_id"`
	Duration        int       `json:"duration"`
	Overtime        bool      `json:"overtime"`
	OvertimeSeconds int       `json:"overtime_seconds"`
	Season          int       `json:"season"`
	Date            time.Time `json:"date"`
	Visibility      string    `json:"visibility"`
	Groups          []Group   `json:"groups"`
	Blue            Team      `json:"blue"`
	Orange          Team      `json:"orange"`
	PlaylistName    string    `json:"playlist_name"`
	MapName         string    `json:"map_name"`
}

// Uploader .
type Uploader struct {
	SteamID    string `json:"steam_id"`
	Name       string `json:"name"`
	ProfileURL string `json:"profile_url"`
	Avatar     string `json:"avatar"`
}

// Group .
type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Team .
type Team struct {
	Color   string    `json:"color"`
	Name    string    `json:"name"`
	Players []Player  `json:"players"`
	Stats   TeamStats `json:"stats"`
}

// Player .
type Player struct {
	StartTime           float64     `json:"start_time"`
	EndTime             float64     `json:"end_time"`
	Name                string      `json:"name"`
	ID                  ID          `json:"id"`
	Mvp                 bool        `json:"mvp,omitempty"`
	CarID               int         `json:"car_id"`
	CarName             string      `json:"car_name"`
	Camera              Camera      `json:"camera"`
	SteeringSensitivity float64     `json:"steering_sensitivity"`
	Stats               PlayerStats `json:"stats"`
}

// ID .
type ID struct {
	Platform string `json:"platform"`
	ID       string `json:"id"`
}

// Camera .
type Camera struct {
	Fov             int     `json:"fov"`
	Height          int     `json:"height"`
	Pitch           int     `json:"pitch"`
	Distance        int     `json:"distance"`
	Stiffness       float64 `json:"stiffness"`
	SwivelSpeed     float64 `json:"swivel_speed"`
	TransitionSpeed float64 `json:"transition_speed"`
}

// PlayerStats .
type PlayerStats struct {
	Core        *PlayerCore        `json:"core,omitempty"`
	Boost       *PlayerBoost       `json:"boost,omitempty"`
	Movement    *PlayerMovement    `json:"movement,omitempty"`
	Positioning *PlayerPositioning `json:"positioning,omitempty"`
	Demolitions *PlayerDemolitions `json:"demo,omitempty"`
}

// PlayerCore .
type PlayerCore struct {
	Shots              float64 `json:"shots"`
	Goals              float64 `json:"goals"`
	Saves              float64 `json:"saves"`
	Assists            float64 `json:"assists"`
	Score              float64 `json:"score"`
	ShootingPercentage float64 `json:"shooting_percentage"`
}

// PlayerBoost .
type PlayerBoost struct {
	Bpm                       float64 `json:"bpm"`
	Bcpm                      float64 `json:"bcpm"`
	AvgAmount                 float64 `json:"avg_amount"`
	AmountCollected           float64 `json:"amount_collected"`
	AmountStolen              float64 `json:"amount_stolen"`
	AmountCollectedBig        float64 `json:"amount_collected_big"`
	AmountStolenBig           float64 `json:"amount_stolen_big"`
	AmountCollectedSmall      float64 `json:"amount_collected_small"`
	AmountStolenSmall         float64 `json:"amount_stolen_small"`
	CountCollectedBig         float64 `json:"count_collected_big"`
	CountStolenBig            float64 `json:"count_stolen_big"`
	CountCollectedSmall       float64 `json:"count_collected_small"`
	CountStolenSmall          float64 `json:"count_stolen_small"`
	AmountOverfill            float64 `json:"amount_overfill"`
	AmountOverfillStolen      float64 `json:"amount_overfill_stolen"`
	AmountUsedWhileSupersonic float64 `json:"amount_used_while_supersonic"`
	TimeZeroBoost             float64 `json:"time_zero_boost"`
	PercentZeroBoost          float64 `json:"percent_zero_boost"`
	TimeFullBoost             float64 `json:"time_full_boost"`
	PercentFullBoost          float64 `json:"percent_full_boost"`
	TimeBoost025              float64 `json:"time_boost_0_25"`
	TimeBoost2550             float64 `json:"time_boost_25_50"`
	TimeBoost5075             float64 `json:"time_boost_50_75"`
	TimeBoost75100            float64 `json:"time_boost_75_100"`
	PercentBoost025           float64 `json:"percent_boost_0_25"`
	PercentBoost2550          float64 `json:"percent_boost_25_50"`
	PercentBoost5075          float64 `json:"percent_boost_50_75"`
	PercentBoost75100         float64 `json:"percent_boost_75_100"`
}

// PlayerMovement .
type PlayerMovement struct {
	AvgSpeed               float64 `json:"avg_speed"`
	TotalDistance          float64 `json:"total_distance"`
	TimeSupersonicSpeed    float64 `json:"time_supersonic_speed"`
	TimeBoostSpeed         float64 `json:"time_boost_speed"`
	TimeSlowSpeed          float64 `json:"time_slow_speed"`
	TimeGround             float64 `json:"time_ground"`
	TimeLowAir             float64 `json:"time_low_air"`
	TimeHighAir            float64 `json:"time_high_air"`
	TimePowerslide         float64 `json:"time_powerslide"`
	CountPowerslide        float64 `json:"count_powerslide"`
	AvgPowerslideDuration  float64 `json:"avg_powerslide_duration"`
	AvgSpeedPercentage     float64 `json:"avg_speed_percentage"`
	PercentSlowSpeed       float64 `json:"percent_slow_speed"`
	PercentBoostSpeed      float64 `json:"percent_boost_speed"`
	PercentSupersonicSpeed float64 `json:"percent_supersonic_speed"`
	PercentGround          float64 `json:"percent_ground"`
	PercentLowAir          float64 `json:"percent_low_air"`
	PercentHighAir         float64 `json:"percent_high_air"`
}

// PlayerPositioning .
type PlayerPositioning struct {
	AvgDistanceToBall             float64 `json:"avg_distance_to_ball"`
	AvgDistanceToBallPossession   float64 `json:"avg_distance_to_ball_possession"`
	AvgDistanceToBallNoPossession float64 `json:"avg_distance_to_ball_no_possession"`
	AvgDistanceToMates            float64 `json:"avg_distance_to_mates"`
	TimeDefensiveThird            float64 `json:"time_defensive_third"`
	TimeNeutralThird              float64 `json:"time_neutral_third"`
	TimeOffensiveThird            float64 `json:"time_offensive_third"`
	TimeDefensiveHalf             float64 `json:"time_defensive_half"`
	TimeOffensiveHalf             float64 `json:"time_offensive_half"`
	TimeBehindBall                float64 `json:"time_behind_ball"`
	TimeInfrontBall               float64 `json:"time_infront_ball"`
	TimeMostBack                  float64 `json:"time_most_back"`
	TimeMostForward               float64 `json:"time_most_forward"`
	GoalsAgainstWhileLastDefender float64 `json:"goals_against_while_last_defender"`
	TimeClosestToBall             float64 `json:"time_closest_to_ball"`
	TimeFarthestFromBall          float64 `json:"time_farthest_from_ball"`
	PercentDefensiveThird         float64 `json:"percent_defensive_third"`
	PercentOffensiveThird         float64 `json:"percent_offensive_third"`
	PercentNeutralThird           float64 `json:"percent_neutral_third"`
	PercentDefensiveHalf          float64 `json:"percent_defensive_half"`
	PercentOffensiveHalf          float64 `json:"percent_offensive_half"`
	PercentBehindBall             float64 `json:"percent_behind_ball"`
	PercentInfrontBall            float64 `json:"percent_infront_ball"`
	PercentMostBack               float64 `json:"percent_most_back"`
	PercentMostForward            float64 `json:"percent_most_forward"`
	PercentClosestToBall          float64 `json:"percent_closest_to_ball"`
	PercentFarthestFromBall       float64 `json:"percent_farthest_from_ball"`
}

// PlayerDemolitions .
type PlayerDemolitions struct {
	Inflicted float64 `json:"inflicted"`
	Taken     float64 `json:"taken"`
}

// TeamStats .
type TeamStats struct {
	Ball        *TeamBall        `json:"ball,omitempty"`
	Core        *TeamCore        `json:"core,omitempty"`
	Boost       *TeamBoost       `json:"boost,omitempty"`
	Movement    *TeamMovement    `json:"movement,omitempty"`
	Positioning *TeamPositioning `json:"positioning,omitempty"`
	Demolitions *TeamDemolitions `json:"demo,omitempty"`
}

// TeamBall .
type TeamBall struct {
	PossessionTime float64 `json:"possession_time"`
	TimeInSide     float64 `json:"time_in_side"`
}

// TeamCore .
type TeamCore struct {
	Shots              float64 `json:"shots"`
	Goals              float64 `json:"goals"`
	Saves              float64 `json:"saves"`
	Assists            float64 `json:"assists"`
	Score              float64 `json:"score"`
	ShootingPercentage float64 `json:"shooting_percentage"`
}

// TeamBoost .
type TeamBoost struct {
	Bpm                       float64 `json:"bpm"`
	Bcpm                      float64 `json:"bcpm"`
	AvgAmount                 float64 `json:"avg_amount"`
	AmountCollected           float64 `json:"amount_collected"`
	AmountStolen              float64 `json:"amount_stolen"`
	AmountCollectedBig        float64 `json:"amount_collected_big"`
	AmountStolenBig           float64 `json:"amount_stolen_big"`
	AmountCollectedSmall      float64 `json:"amount_collected_small"`
	AmountStolenSmall         float64 `json:"amount_stolen_small"`
	CountCollectedBig         float64 `json:"count_collected_big"`
	CountStolenBig            float64 `json:"count_stolen_big"`
	CountCollectedSmall       float64 `json:"count_collected_small"`
	CountStolenSmall          float64 `json:"count_stolen_small"`
	AmountOverfill            float64 `json:"amount_overfill"`
	AmountOverfillStolen      float64 `json:"amount_overfill_stolen"`
	AmountUsedWhileSupersonic float64 `json:"amount_used_while_supersonic"`
	TimeZeroBoost             float64 `json:"time_zero_boost"`
	TimeFullBoost             float64 `json:"time_full_boost"`
	TimeBoost025              float64 `json:"time_boost_0_25"`
	TimeBoost2550             float64 `json:"time_boost_25_50"`
	TimeBoost5075             float64 `json:"time_boost_50_75"`
	TimeBoost75100            float64 `json:"time_boost_75_100"`
}

// TeamMovement .
type TeamMovement struct {
	TotalDistance       float64 `json:"total_distance"`
	TimeSupersonicSpeed float64 `json:"time_supersonic_speed"`
	TimeBoostSpeed      float64 `json:"time_boost_speed"`
	TimeSlowSpeed       float64 `json:"time_slow_speed"`
	TimeGround          float64 `json:"time_ground"`
	TimeLowAir          float64 `json:"time_low_air"`
	TimeHighAir         float64 `json:"time_high_air"`
	TimePowerslide      float64 `json:"time_powerslide"`
	CountPowerslide     float64 `json:"count_powerslide"`
}

// TeamPositioning .
type TeamPositioning struct {
	TimeDefensiveThird float64 `json:"time_defensive_third"`
	TimeNeutralThird   float64 `json:"time_neutral_third"`
	TimeOffensiveThird float64 `json:"time_offensive_third"`
	TimeDefensiveHalf  float64 `json:"time_defensive_half"`
	TimeOffensiveHalf  float64 `json:"time_offensive_half"`
	TimeBehindBall     float64 `json:"time_behind_ball"`
	TimeInfrontBall    float64 `json:"time_infront_ball"`
}

// TeamDemolitions .
type TeamDemolitions struct {
	Inflicted float64 `json:"inflicted"`
	Taken     float64 `json:"taken"`
}
