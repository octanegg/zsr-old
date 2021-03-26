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
	RocketLeagueID  string    `json:"rocketLeagueId" bson:"rocket_league_id"`
	MatchGUID       string    `json:"matchGuid" bson:"match_guid"`
	Title           string    `json:"title" bson:"title"`
	MapCode         string    `json:"mapCode" bson:"map_code"`
	MatchType       string    `json:"matchType" bson:"match_type"`
	TeamSize        int       `json:"teamSize" bson:"team_size"`
	PlaylistID      string    `json:"playlistId" bson:"playlist_id"`
	Duration        int       `json:"duration" bson:"duration"`
	Overtime        bool      `json:"overtime" bson:"overtime"`
	OvertimeSeconds int       `json:"overtimeSeconds" bson:"overtime_seconds"`
	Season          int       `json:"season" bson:"season"`
	Date            time.Time `json:"date" bson:"date"`
	Visibility      string    `json:"visibility" bson:"visibility"`
	Groups          []Group   `json:"groups" bson:"groups"`
	Blue            Team      `json:"blue" bson:"blue"`
	Orange          Team      `json:"orange" bson:"orange"`
	PlaylistName    string    `json:"playlistName" bson:"playlist_name"`
	MapName         string    `json:"mapName" bson:"map_name"`
}

// Uploader .
type Uploader struct {
	SteamID    string `json:"steamId" bson:"steam_id"`
	Name       string `json:"name" bson:"name"`
	ProfileURL string `json:"profileUrl" bson:"profile_url"`
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
	StartTime           float64     `json:"startTime" bson:"start_time"`
	EndTime             float64     `json:"endTime" bson:"end_time"`
	Name                string      `json:"name" bson:"name"`
	ID                  ID          `json:"id" bson:"id"`
	Mvp                 bool        `json:"mvp,omitempty" bson:"mvp,omitempty"`
	CarID               int         `json:"carId" bson:"car_id"`
	CarName             string      `json:"carName" bson:"car_name"`
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
