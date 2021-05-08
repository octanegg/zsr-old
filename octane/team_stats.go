package octane

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
