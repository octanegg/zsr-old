package octane

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Events .
type Events struct {
	Events []*Event `json:"events"`
	*Pagination
}

// Event .
type Event struct {
	ID        *primitive.ObjectID `json:"_id" bson:"_id"`
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
	ID         int         `json:"_id" bson:"_id"`
	Name       string      `json:"name" bson:"name"`
	Format     string      `json:"format" bson:"format"`
	Region     string      `json:"region,omitempty" bson:"region,omitempty"`
	StartDate  *time.Time  `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate    *time.Time  `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Prize      *Prize      `json:"prize,omitempty" bson:"prize,omitempty"`
	Liquipedia string      `json:"liquipedia,omitempty" bson:"liquipedia,omitempty"`
	Qualifier  bool        `json:"qualifier,omitempty" bson:"qualifier,omitempty"`
	Substages  []*Substage `json:"substages,omitempty" bson:"substages,omitempty"`
}

// Substage .
type Substage struct {
	ID     int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Format string `json:"format" bson:"format"`
}

// Prize .
type Prize struct {
	Amount   float64 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
}

func (c *client) FindEvents(filter bson.M, pagination *Pagination, sort *Sort) (*Events, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionEvents)

	opts := options.Find()
	if pagination != nil {
		opts.SetSkip((pagination.Page - 1) * pagination.PerPage)
		opts.SetLimit(pagination.PerPage)
	}

	if sort != nil {
		opts.SetSort(bson.M{sort.Field: sort.Order})
	}

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var events []*Event
	for cursor.Next(ctx) {
		var event Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	if err != nil {
		return nil, err
	}

	if pagination != nil {
		pagination.PageSize = len(events)
	}

	return &Events{
		events,
		pagination,
	}, nil
}

func (c *client) FindEvent(oid *primitive.ObjectID) (*Event, error) {
	events, err := c.FindEvents(bson.M{"_id": oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(events.Events) == 0 {
		return nil, nil
	}

	return events.Events[0], nil
}

func (c *client) InsertEvent(event interface{}) (*primitive.ObjectID, error) {
	ids, err := c.InsertEvents([]interface{}{event})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
}

func (c *client) InsertEvents(events []interface{}) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionEvents)

	res, err := coll.InsertMany(ctx, events)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *client) DeleteEvent(filter bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionEvents)

	res, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
