package stats

var (
	// PlayerStats .
	PlayerStatsMapping = map[string]map[string]string{
		"core":        PlayerCoreMapping,
		"boost":       PlayerBoostMapping,
		"movement":    PlayerMovementMapping,
		"positioning": PlayerPositioningMapping,
		"demo":        PlayerDemolitionsMapping,
		"advanced":    PlayerAdvancedMapping,
	}

	// PlayerCore .
	PlayerCoreMapping = map[string]string{
		"score":              "score",
		"goals":              "goals",
		"assists":            "assists",
		"saves":              "saves",
		"shots":              "shots",
		"shootingPercentage": "shooting_percentage",
	}

	// PlayerBoost .
	PlayerBoostMapping = map[string]string{
		"bpm":                       "bpm",
		"bcpm":                      "bcpm",
		"avgAmount":                 "avg_amount",
		"amountCollected":           "amount_collected",
		"amountStolen":              "amount_stolen",
		"amountCollectedBig":        "amount_collected_big",
		"amountStolenBig":           "amount_stolen_big",
		"amountCollectedSmall":      "amount_collected_small",
		"amountStolenSmall":         "amount_stolen_small",
		"countCollectedBig":         "count_collected_big",
		"countStolenBig":            "count_stolen_big",
		"countCollectedSmall":       "count_collected_small",
		"countStolenSmall":          "count_stolen_small",
		"amountOverfill":            "amount_overfill",
		"amountOverfillStolen":      "amount_overfill_stolen",
		"amountUsedWhileSupersonic": "amount_used_while_supersonic",
		"timeZeroBoost":             "time_zero_boost",
		"percentZeroBoost":          "percent_zero_boost",
		"timeFullBoost":             "time_full_boost",
		"percentFullBoost":          "percent_full_boost",
		"timeBoost0To25":            "time_boost_0_25",
		"percentBoost0To25":         "percent_boost_0_25",
		"timeBoost25To50":           "time_boost_25_50",
		"percentBoost25To50":        "percent_boost_25_50",
		"timeBoost50To75":           "time_boost_50_75",
		"percentBoost50To75":        "percent_boost_50_75",
		"timeBoost75To100":          "time_boost_75_100",
		"percentBoost75To100":       "percent_boost_75_100",
	}

	// PlayerMovement .
	PlayerMovementMapping = map[string]string{
		"avgSpeed":               "avg_speed",
		"totalDistance":          "total_distance",
		"timeSupersonicSpeed":    "time_supersonic_speed",
		"timeBoostSpeed":         "time_boost_speed",
		"timeSlowSpeed":          "time_slow_speed",
		"timeGround":             "time_ground",
		"timeLowAir":             "time_low_air",
		"timeHighAir":            "time_high_air",
		"timePowerslide":         "time_powerslide",
		"countPowerslide":        "count_powerslide",
		"avgPowerslideDuration":  "avg_powerslide_duration",
		"avgSpeedPercentage":     "avg_speed_percentage",
		"percentSlowSpeed":       "percent_slow_speed",
		"percentBoostSpeed":      "percent_boost_speed",
		"percentSupersonicSpeed": "percent_supersonic_speed",
		"percentGround":          "percent_ground",
		"percentLowAir":          "percent_low_air",
		"percentHighAir":         "percent_high_air",
	}

	// PlayerPositioning .
	PlayerPositioningMapping = map[string]string{
		"avgDistanceToBall":             "avg_distance_to_ball",
		"avgDistanceToBallPossession":   "avg_distance_to_ball_possession",
		"avgDistanceToBallNoPossession": "avg_distance_to_ball_no_possession",
		"avgDistanceToMates":            "avg_distance_to_mates",
		"timeDefensiveThird":            "time_defensive_third",
		"timeNeutralThird":              "time_neutral_third",
		"timeOffensiveThird":            "time_offensive_third",
		"timeDefensiveHalf":             "time_defensive_half",
		"timeOffensiveHalf":             "time_offensive_half",
		"timeBehindBall":                "time_behind_ball",
		"timeInfrontBall":               "time_infront_ball",
		"timeMostBack":                  "time_most_back",
		"timeMostForward":               "time_most_forward",
		"goalsAgainstWhileLastDefender": "goals_against_while_last_defender",
		"timeClosestToBall":             "time_closest_to_ball",
		"timeFarthestFromBall":          "time_farthest_from_ball",
		"percentDefensiveThird":         "percent_defensive_third",
		"percentOffensiveThird":         "percent_offensive_third",
		"percentNeutralThird":           "percent_neutral_third",
		"percentDefensiveHalf":          "percent_defensive_half",
		"percentOffensiveHalf":          "percent_offensive_half",
		"percentBehindBall":             "percent_behind_ball",
		"percentInfrontBall":            "percent_infront_ball",
		"percentMostBack":               "percent_most_back",
		"percentMostForward":            "percent_most_forward",
		"percentClosestToBall":          "percent_closest_to_ball",
		"percentFarthestFromBall":       "percent_farthest_from_ball",
	}

	// PlayerDemolitions .
	PlayerDemolitionsMapping = map[string]string{
		"inflicted": "inflicted",
		"taken":     "taken",
	}

	// PlayerAdvanced .
	PlayerAdvancedMapping = map[string]string{
		"goalParticipation": "goal_participation",
		"rating":            "rating",
	}

	// TeamStats .
	TeamStatsMapping = map[string]map[string]string{
		"core":        TeamCoreMapping,
		"ball":        TeamBallMapping,
		"boost":       TeamBoostMapping,
		"movement":    TeamMovementMapping,
		"positioning": TeamPositioningMapping,
		"demo":        TeamDemolitionsMapping,
	}

	// TeamBall .
	TeamBallMapping = map[string]string{
		"possessionTime": "possession_time",
		"timeInSide":     "time_in_side",
	}

	// TeamCore .
	TeamCoreMapping = map[string]string{
		"score":              "score",
		"goals":              "goals",
		"assists":            "assists",
		"saves":              "saves",
		"shots":              "shots",
		"shootingPercentage": "shooting_percentage",
	}

	// TeamBoost .
	TeamBoostMapping = map[string]string{
		"bpm":                       "bpm",
		"bcpm":                      "bcpm",
		"avgAmount":                 "avg_amount",
		"amountCollected":           "amount_collected",
		"amountStolen":              "amount_stolen",
		"amountCollectedBig":        "amount_collected_big",
		"amountStolenBig":           "amount_stolen_big",
		"amountCollectedSmall":      "amount_collected_small",
		"amountStolenSmall":         "amount_stolen_small",
		"countCollectedBig":         "count_collected_big",
		"countStolenBig":            "count_stolen_big",
		"countCollectedSmall":       "count_collected_small",
		"countStolenSmall":          "count_stolen_small",
		"amountOverfill":            "amount_overfill",
		"amountOverfillStolen":      "amount_overfill_stolen",
		"amountUsedWhileSupersonic": "amount_used_while_supersonic",
		"timeZeroBoost":             "time_zero_boost",
		"timeFullBoost":             "time_full_boost",
		"timeBoost0To25":            "time_boost_0_25",
		"timeBoost25To50":           "time_boost_25_50",
		"timeBoost50To75":           "time_boost_50_75",
		"timeBoost75To100":          "time_boost_75_100",
	}

	// TeamMovement .
	TeamMovementMapping = map[string]string{
		"totalDistance":       "total_distance",
		"timeSupersonicSpeed": "time_supersonic_speed",
		"timeBoostSpeed":      "time_boost_speed",
		"timeSlowSpeed":       "time_slow_speed",
		"timeGround":          "time_ground",
		"timeLowAir":          "time_low_air",
		"timeHighAir":         "time_high_air",
		"timePowerslide":      "time_powerslide",
		"countPowerslide":     "count_powerslide",
	}

	// TeamPositioning .
	TeamPositioningMapping = map[string]string{
		"timeDefensiveThird": "time_defensive_third",
		"timeNeutralThird":   "time_neutral_third",
		"timeOffensiveThird": "time_offensive_third",
		"timeDefensiveHalf":  "time_defensive_half",
		"timeOffensiveHalf":  "time_offensive_half",
		"timeBehindBall":     "time_behind_ball",
		"timeInfrontBall":    "time_infront_ball",
	}

	// TeamDemolitions .
	TeamDemolitionsMapping = map[string]string{
		"inflicted": "inflicted",
		"taken":     "taken",
	}
)

var FieldsToAverage = []string{
	"shootingPercentage",
	"goalParticipation",
	"rating",
}

var FieldsToAverageOverReplays = []string{
	"bpm",
	"bcpm",
	"avgSpeed",
	"avgSpeedPercentage",
	"percentSlowSpeed",
	"avgDistanceToBall",
	"avgDistanceToBallPossession",
	"avgDistanceToBallNoPossession",
	"avgDistanceToMates",
	"avgPowerslideDuration",
	"percentZeroBoost",
	"percentFullBoost",
	"percentBoost0To25",
	"percentBoost25To50",
	"percentBoost50To75",
	"percentBoost75To100",
	"percentBoostSpeed",
	"percentSupersonicSpeed",
	"percentGround",
	"percentLowAir",
	"percentHighAir",
	"percentDefensiveThird",
	"percentOffensiveThird",
	"percentNeutralThird",
	"percentDefensiveHalf",
	"percentOffensiveHalf",
	"percentBehindBall",
	"percentInfrontBall",
	"percentMostBack",
	"percentMostForward",
	"percentClosestToBall",
	"percentFarthestFromBall",
}
