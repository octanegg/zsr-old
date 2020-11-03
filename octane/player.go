package octane

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Player .
type Player struct {
	ID      *primitive.ObjectID `json:"_id" bson:"_id"`
	Tag     string              `json:"tag" bson:"tag"`
	Name    string              `json:"name,omitempty" bson:"name,omitempty"`
	Country string              `json:"country,omitempty" bson:"country,omitempty"`
	Team    string              `json:"team,omitempty" bson:"team,omitempty"`
	Account *Account            `json:"account,omitempty" bson:"account,omitempty"`
}

// Account .
type Account struct {
	Platform string `json:"platform" bson:"platform"`
	ID       string `json:"id" bson:"id"`
}
