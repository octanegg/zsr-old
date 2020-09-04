package octane

// Stats .
type Stats struct {
	Core *CoreStats `json:"core" bson:"core"`
}

// CoreStats .
type CoreStats struct {
	Score              *int     `json:"score" bson:"score"`
	Goals              *int     `json:"goals" bson:"goals"`
	Assists            *int     `json:"assists" bson:"assists"`
	Saves              *int     `json:"saves" bson:"saves"`
	Shots              *int     `json:"shots" bson:"shots"`
	ShootingPercentage *float64 `json:"shooting_percentage" bson:"shooting_percentage"`
	GoalParticipation  *float64 `json:"goal_participation" bson:"goal_participation"`
	Rating             *float64 `json:"rating" bson:"rating"`
	MVP                *bool    `json:"mvp" bson:"mvp"`
}
