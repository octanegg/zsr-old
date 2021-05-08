package octane

import "go.mongodb.org/mongo-driver/bson/primitive"

// Player .
type Player struct {
	ID         *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug       string              `json:"slug,omitempty" bson:"slug,omitempty"`
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
