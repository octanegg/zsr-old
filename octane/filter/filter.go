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
	Value interface{}
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

// Dates .
func Dates(key string, before, after string) *Field {
	f := bson.M{}
	if before != "" {
		b, err := time.Parse("2006-01-02", before)
		if err != nil {
			return nil
		}
		f["$lte"] = b
	}

	if after != "" {
		a, err := time.Parse("2006-01-02", after)
		if err != nil {
			return nil
		}
		f["$gte"] = a
	}

	if len(f) == 0 {
		return nil
	}

	return &Field{
		Key:   key,
		Value: f,
	}
}

// BeforeDate .
func BeforeDate(key string, val string) *Field {
	t, err := time.Parse(time.RFC3339, val)
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
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil
	}

	return &Field{
		Key:   key,
		Value: bson.M{"$gte": t},
	}
}

// Bool .
func Bool(key string, val string) *Field {
	b, err := strconv.ParseBool(val)
	if err != nil {
		return nil
	}

	op := "$eq"
	if !b {
		op = "$ne"
	}

	return &Field{
		Key:   key,
		Value: bson.M{op: true},
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

// Or .
func Or(fields ...*Field) *Field {
	f := []bson.M{}
	for _, field := range fields {
		if field != nil {
			f = append(f, bson.M{field.Key: field.Value})
		}
	}

	if len(f) == 0 {
		return nil
	}

	return &Field{
		Key:   "$or",
		Value: f,
	}
}
