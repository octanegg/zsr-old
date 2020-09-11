package deprecated

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // sql driver
	"github.com/octanegg/core/octane"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	oldDBURI = "DB_OLD"
)

type deprecated struct {
	DB *sql.DB
}

// Deprecated .
type Deprecated interface {
	GetMatches(*EventLinkage) ([]*Match, error)
	GetGameMap(int) (map[string]map[int]*Game, error)
	Propogate(*octane.Game, map[*primitive.ObjectID]string) error
}

// New .
func New() (Deprecated, error) {
	db, err := sql.Open("mysql", os.Getenv(oldDBURI))
	if err != nil {
		return nil, err
	}

	return &deprecated{
		DB: db,
	}, nil
}
