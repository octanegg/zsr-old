package collection

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection .
type Collection interface {
	Find(bson.M, bson.M, *Pagination) ([]interface{}, error)
	FindOne(bson.M) (interface{}, error)
	UpdateOne(interface{}, interface{}) (interface{}, error)
	Insert([]interface{}) ([]interface{}, error)
	InsertOne(interface{}) (*primitive.ObjectID, error)
	Delete(bson.M) (int64, error)
	Pipeline([]bson.M, func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error)
}

type collection struct {
	Collection *mongo.Collection
	Decode     func(*mongo.Cursor) (interface{}, error)
}

// Pagination .
type Pagination struct {
	Page     int64 `json:"page"`
	PerPage  int64 `json:"perPage"`
	PageSize int   `json:"pageSize"`
}

// New .
func New(c *mongo.Collection, d func(*mongo.Cursor) (interface{}, error)) Collection {
	return &collection{c, d}
}

func (c *collection) Find(filter, sort bson.M, pagination *Pagination) ([]interface{}, error) {
	opts := options.Find()
	if pagination != nil {
		opts.SetSkip((pagination.Page - 1) * pagination.PerPage)
		opts.SetLimit(pagination.PerPage)
	}

	opts.SetSort(sort)
	cursor, err := c.Collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var data []interface{}
	if c.Decode == nil {
		if err := cursor.All(context.TODO(), &data); err != nil {
			return nil, err
		}
		return data, nil
	}

	for cursor.Next(context.TODO()) {
		i, err := c.Decode(cursor)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}

func (c *collection) FindOne(filter bson.M) (interface{}, error) {
	data, err := c.Find(filter, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("no data found")
	}

	return data[0], nil
}

func (c *collection) UpdateOne(filter, data interface{}) (interface{}, error) {
	res, err := c.Collection.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return nil, err
	}

	return res.UpsertedID, nil
}

func (c *collection) Insert(data []interface{}) ([]interface{}, error) {
	res, err := c.Collection.InsertMany(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *collection) InsertOne(data interface{}) (*primitive.ObjectID, error) {
	ids, err := c.Insert([]interface{}{data})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
}

func (c *collection) Delete(filter bson.M) (int64, error) {
	res, err := c.Collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func (c *collection) Pipeline(pipeline []bson.M, decode func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error) {
	cursor, err := c.Collection.Aggregate(context.TODO(), pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		return nil, err
	}

	var data []interface{}
	if decode == nil {
		if err := cursor.All(context.TODO(), &data); err != nil {
			return nil, err
		}
		return data, nil
	}

	for cursor.Next(context.TODO()) {
		i, err := decode(cursor)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}
