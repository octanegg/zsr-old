package octane

import (
	"errors"
	"reflect"
	"time"

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
	StartDate time.Time          `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   time.Time          `json:"endDate,omitempty" bson:"endDate,omitempty"`
	Region    string             `json:"region,omitempty" bson:"region,omitempty"`
	Mode      int                `json:"mode,omitempty" bson:"mode,omitempty"`
	Prize     Prize              `json:"prize,omitempty" bson:"prize,omitempty"`
	Tier      string             `json:"tier,omitempty" bson:"tier,omitempty"`
	Stages    []Stage            `json:"stages,omitempty" bson:"stages,omitempty"`
}

// Stage .
type Stage struct {
	Name       string     `json:"name" bson:"name"`
	Format     string     `json:"format,omitempty" bson:"format,omitempty"`
	Region     string     `json:"region,omitempty" bson:"region,omitempty"`
	StartDate  time.Time  `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate    time.Time  `json:"endDate,omitempty" bson:"endDate,omitempty"`
	Prize      Prize      `json:"prize,omitempty" bson:"prize,omitempty"`
	Liquipedia string     `json:"liquipedia,omitempty" bson:"liquipedia,omitempty"`
	Qualifier  bool       `json:"qualifier,omitempty" bson:"qualifier,omitempty"`
	Substages  []Substage `json:"substages,omitempty" bson:"substages,omitempty"`
}

// Substage .
type Substage struct {
	Name   string `json:"name" bson:"name"`
	Format string `json:"format" bson:"format"`
}

// Prize .
type Prize struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
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

func (c *client) FindEventByID(oid *primitive.ObjectID) (*Event, error) {
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

func (c *client) InsertEvent(event *Event) (*ObjectID, error) {
	event.ID = primitive.NewObjectID()
	id, err := c.Insert("events", event)
	if err != nil {
		return nil, err
	}

	return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdateEvent(oid *primitive.ObjectID, fields *Event) (*ObjectID, error) {
	event, err := c.FindEventByID(oid)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("No event found for ID")
	}

	filter := bson.M{"_id": oid}
	update := updateFields(reflect.ValueOf(event).Elem(), reflect.ValueOf(fields).Elem()).(Event)
	update.ID = *oid

	id, err := c.Replace("events", filter, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}

	return &ObjectID{oid.Hex()}, nil
}
