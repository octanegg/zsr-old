package deprecated

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // sql driver
)

const (
	oldDBURI = "DB_OLD"
)

type deprecated struct {
	DB *sql.DB
}

// Deprecated .
type Deprecated interface {
	UpdateMatches([]*UpdateMatchContext) error
	GetMatches(*GetMatchesContext) ([]*Match, error)
	GetMatch(*GetMatchContext) (*Match, error)
	DeleteGame(*DeleteGameContext) error
	GetGames(*GetGamesContext) ([]*Game, error)
	InsertGame(*Game) (error)

	getLinkages([]int) ([]*EventLinkage, error)
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
