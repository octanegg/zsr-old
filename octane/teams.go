package octane

import "go.mongodb.org/mongo-driver/bson/primitive"

// Team .
type Team struct {
	ID     *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug   string              `json:"slug,omitempty" bson:"slug,omitempty"`
	Name   string              `json:"name,omitempty" bson:"name,omitempty"`
	Region string              `json:"region,omitempty" bson:"region,omitempty"`
	Image  string              `json:"image,omitempty" bson:"image,omitempty"`
}
