package octane

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Events .
type Events struct {
	Events []interface{} `json:"events"`
}

// Event .
type Event struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	StartDate string             `json:"startDate" bson:"startDate"`
	EndDate   string             `json:"endDate" bson:"endDate"`
	Region    string             `json:"region" bson:"region"`
	Mode      int                `json:"mode" bson:"mode"`
	Prize     Prize              `json:"prize" bson:"prize"`
	Tier      string             `json:"tier" bson:"tier"`
	Stages    []Stage            `json:"stages" bson:"stages"`
}

// Stage .
type Stage struct {
	Name       string     `json:"name" bson:"name"`
	Format     string     `json:"format" bson:"format"`
	Region     string     `json:"region" bson:"region"`
	StartDate  string     `json:"startDate" bson:"startDate"`
	EndDate    string     `json:"endDate" bson:"endDate"`
	Prize      Prize      `json:"prize,omitempty" bson:"prize"`
	Liquipedia string     `json:"liquipedia" bson:"liquipedia"`
	Qualifier  bool       `json:"qualifier,omitempty" bson:"qualifier"`
	Substages  []Substage `json:"substages,omitempty" bson:"substages"`
}

// Substage .
type Substage struct {
	Name   string `json:"name,omitempty" bson:"name"`
	Format string `json:"format,omitempty" bson:"format"`
}

// Prize .
type Prize struct {
	Amount   float64 `json:"amount,omitempty" bson:"amount"`
	Currency string  `json:"currency,omitempty" bson:"currency"`
}

func (c *client) FindEvents(filter bson.M) (*Events, error) {
	events, err := c.Find("events", filter, func(cursor *mongo.Cursor) (interface{}, error) {
		var event Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		return event, nil
	})

	if err != nil {
		return nil, err
	}

	return &Events{events}, nil
}

func (c *client) FindEventByID(id string) (*Event, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	events, err := c.FindEvents(bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}

	if len(events.Events) == 0 {
		return nil, nil
	}

	event := events.Events[0].(Event)
	return &event, nil
}
