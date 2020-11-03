package octane

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Team .
type Team struct {
	ID   *primitive.ObjectID `json:"_id" bson:"_id"`
	Name string              `json:"name" bson:"name"`
}
