package octane

import "go.mongodb.org/mongo-driver/bson/primitive"

// Player .
type Player struct {
	ID  primitive.ObjectID `json:"id" bson:"_id"`
	Tag string             `json:"tag" bson:"tag"`
}
