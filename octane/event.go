package octane

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"time"

	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Event .
type Event struct {
	ID        *primitive.ObjectID `json:"id" bson:"_id"`
	Name      string              `json:"name" bson:"name"`
	StartDate *time.Time          `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate   *time.Time          `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Region    string              `json:"region" bson:"region"`
	Mode      int                 `json:"mode" bson:"mode"`
	Prize     *Prize              `json:"prize,omitempty" bson:"prize,omitempty"`
	Tier      string              `json:"tier,omitempty" bson:"tier,omitempty"`
	Stages    []*Stage            `json:"stages,omitempty" bson:"stages,omitempty"`
}

// Stage .
type Stage struct {
	Name       string      `json:"name" bson:"name"`
	Format     string      `json:"format" bson:"format"`
	Region     string      `json:"region" bson:"region"`
	StartDate  *time.Time  `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate    *time.Time  `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Prize      *Prize      `json:"prize,omitempty" bson:"prize,omitempty"`
	Liquipedia string      `json:"liquipedia" bson:"liquipedia"`
	Qualifier  bool        `json:"qualifier,omitempty" bson:"qualifier,omitempty"`
	Substages  []*Substage `json:"substages,omitempty" bson:"substages,omitempty"`
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

func (c *client) FindEvents(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	events, err := c.Find(config.CollectionEvents, filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
		var event Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		return event, nil
	})

	if err != nil {
		return nil, err
	}

	if events == nil {
		events = make([]interface{}, 0)
	}

	if pagination != nil {
		pagination.PageSize = len(events)
	}

	return &Data{
		events,
		pagination,
	}, nil
}

func (c *client) FindEvent(oid *primitive.ObjectID) (interface{}, error) {
	events, err := c.FindEvents(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(events.Data) == 0 {
		return nil, nil
	}

	return events.Data[0], nil
}

func (c *client) InsertEventWithReader(body io.ReadCloser) (*ObjectID, error) {
	var event Event
	if err := json.NewDecoder(body).Decode(&event); err != nil {
		return nil, err
	}

	id := primitive.NewObjectID()
	event.ID = &id
	oid, err := c.Insert(config.CollectionEvents, event)
	if err != nil {
		return nil, err
	}

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdateEventWithReader(oid *primitive.ObjectID, body io.ReadCloser) (*ObjectID, error) {
	data, err := c.FindEvent(oid)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	var fields Event
	if err := json.NewDecoder(body).Decode(&fields); err != nil {
		return nil, err
	}

	event := data.(Event)
	update := updateFields(reflect.ValueOf(&event).Elem(), reflect.ValueOf(&fields).Elem()).(Event)
	update.ID = oid

	id, err := c.Replace(config.CollectionEvents, oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}

	return &ObjectID{oid.Hex()}, nil
}

func (c *client) InsertEvent(event *Event) (*ObjectID, error) {
	id := primitive.NewObjectID()
	event.ID = &id
	oid, err := c.Insert(config.CollectionEvents, event)
	if err != nil {
		return nil, err
	}

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdateEvent(oid *primitive.ObjectID, fields *Event) (*ObjectID, error) {
	data, err := c.FindEvent(oid)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	event := data.(Event)
	update := updateFields(reflect.ValueOf(&event).Elem(), reflect.ValueOf(fields).Elem()).(Event)
	update.ID = oid

	id, err := c.Replace(config.CollectionEvents, oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}

	return &ObjectID{oid.Hex()}, nil
}

func (c *client) DeleteEvent(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(config.CollectionEvents, oid)
}
