package filter

import (
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Field .
type Field struct {
	Key   string
	Value bson.M
}

// New .
func New(fields ...*Field) bson.M {
	filter := bson.M{}
	for _, field := range fields {
		if field != nil {
			filter[field.Key] = field.Value
		}
	}
	return filter
}

// Strings .
func Strings(key string, vals []string) *Field {
	if len(vals) == 0 {
		return nil
	}

	return &Field{
		Key:   key,
		Value: bson.M{"$in": vals},
	}
}

// BeforeDate .
func BeforeDate(key string, val string) *Field {
	t, err := time.Parse(time.RFC3339Nano, val)
	if err != nil {
		return nil
	}

	return &Field{
		Key:   key,
		Value: bson.M{"$lte": t},
	}
}

// AfterDate .
func AfterDate(key string, val string) *Field {
	t, err := time.Parse(time.RFC3339Nano, val)
	if err != nil {
		return nil
	}

	return &Field{
		Key:   key,
		Value: bson.M{"$gte": t},
	}
}

// Ints .
func Ints(key string, vals []string) *Field {
	if len(vals) == 0 {
		return nil
	}

	var a []int
	for _, val := range vals {
		if i, err := strconv.Atoi(val); err == nil {
			a = append(a, i)
		}
	}

	return &Field{
		Key:   key,
		Value: bson.M{"$in": a},
	}
}

// ObjectIDs .
func ObjectIDs(key string, vals []string) *Field {
	if len(vals) == 0 {
		return nil
	}

	var a []primitive.ObjectID
	for _, val := range vals {
		if i, err := primitive.ObjectIDFromHex(val); err == nil {
			a = append(a, i)
		}
	}

	return &Field{
		Key:   key,
		Value: bson.M{"$in": a},
	}
}
