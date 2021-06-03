package octane

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event .
type Event struct {
	ID        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug      string              `json:"slug,omitempty" bson:"slug,omitempty"`
	Name      string              `json:"name,omitempty" bson:"name,omitempty"`
	StartDate *time.Time          `json:"startDate,omitempty" bson:"start_date,omitempty"`
	EndDate   *time.Time          `json:"endDate,omitempty" bson:"end_date,omitempty"`
	Region    string              `json:"region,omitempty" bson:"region,omitempty"`
	Mode      int                 `json:"mode,omitempty" bson:"mode,omitempty"`
	Prize     *Prize              `json:"prize,omitempty" bson:"prize,omitempty"`
	Tier      string              `json:"tier,omitempty" bson:"tier,omitempty"`
	Image     string              `json:"image,omitempty" bson:"image,omitempty"`
	Groups    []string            `json:"groups,omitempty" bson:"groups,omitempty"`
	Socials   []Social            `json:"socials,omitempty" bson:"socials,omitempty"`
	Stages    []*Stage            `json:"stages,omitempty" bson:"stages,omitempty"`
}

// Stage .
type Stage struct {
	ID         int         `json:"_id" bson:"_id"`
	Name       string      `json:"name,omitempty" bson:"name,omitempty"`
	Format     string      `json:"format,omitempty" bson:"format,omitempty"`
	Region     string      `json:"region,omitempty" bson:"region,omitempty"`
	StartDate  *time.Time  `json:"startDate,omitempty" bson:"start_date,omitempty"`
	EndDate    *time.Time  `json:"endDate,omitempty" bson:"end_date,omitempty"`
	Prize      *Prize      `json:"prize,omitempty" bson:"prize,omitempty"`
	Liquipedia string      `json:"liquipedia,omitempty" bson:"liquipedia,omitempty"`
	Qualifier  bool        `json:"qualifier,omitempty" bson:"qualifier,omitempty"`
	LAN        bool        `json:"lan,omitempty" bson:"lan,omitempty"`
	Location   *Location   `json:"location,omitempty" bson:"location,omitempty"`
	Substages  []*Substage `json:"substages,omitempty" bson:"substages,omitempty"`
}

// Location .
type Location struct {
	Venue   string `json:"venue,omitempty" bson:"venue,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
}

// Substage .
type Substage struct {
	ID     int    `json:"_id" bson:"_id"`
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Format string `json:"format,omitempty" bson:"format,omitempty"`
}

// Prize .
type Prize struct {
	Amount   float64 `json:"amount,omitempty" bson:"amount,omitempty"`
	Currency string  `json:"currency,omitempty" bson:"currency,omitempty"`
}

type Participant struct {
	Team    *Team     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*Player `json:"players,omitempty" bson:"players,omitempty"`
}
