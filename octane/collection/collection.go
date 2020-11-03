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
	Insert([]interface{}) ([]interface{}, error)
	InsertOne(interface{}) (*primitive.ObjectID, error)
	Delete(bson.M) (int64, error)
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
func New(coll *mongo.Collection, decode func(*mongo.Cursor) (interface{}, error)) Collection {
	return &collection{
		Collection: coll,
		Decode:     decode,
	}
}

func (c *collection) Find(filter bson.M, sort bson.M, pagination *Pagination) ([]interface{}, error) {
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
