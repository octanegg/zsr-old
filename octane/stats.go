package octane

// Stats .
type Stats struct {
	MVP  *bool      `json:"mvp" bson:"mvp"`
	Core *CoreStats `json:"core" bson:"core"`
}

// CoreStats .
type CoreStats struct {
	Score              *int     `json:"score" bson:"score"`
	Goals              *int     `json:"goals" bson:"goals"`
	Assists            *int     `json:"assists" bson:"assists"`
	Saves              *int     `json:"saves" bson:"saves"`
	Shots              *int     `json:"shots" bson:"shots"`
	ShootingPercentage *float64 `json:"shootingPercentage" bson:"shootingPercentage"`
	GoalParticipation  *float64 `json:"goalParticipation" bson:"goalParticipation"`
	Rating             *float64 `json:"rating" bson:"rating"`
}
