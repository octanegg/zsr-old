package octane

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
