package deprecated

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/octanegg/core/octane"
	"github.com/octanegg/racer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	insertURL = "INSERT_URL"
)

func (d *deprecated) Propogate(game *octane.Game, lookup map[*primitive.ObjectID]string) error {
	v := url.Values{}
	v.Add("murl", game.OctaneID)
	v.Add("Game", strconv.Itoa(game.Number))
	v.Add("Map", game.Map)
	v.Add("Datee", game.Date.Format("2006-01-02"))
	v.Add("Team", lookup[game.Blue.Team.ID])
	v.Add("Vs", lookup[game.Orange.Team.ID])
	v.Add("Length", strconv.Itoa(game.Duration))

	if game.Blue.Goals > game.Orange.Goals {
		v.Add("Result", lookup[game.Blue.Team.ID])
	} else {
		v.Add("Result", lookup[game.Orange.Team.ID])
	}

	i := 1
	for _, p := range append(game.Blue.Players, game.Orange.Players...) {
		stats := p.Stats.(racer.PlayerStats)
		v.Add(fmt.Sprintf("Player%d", i), lookup[p.Player])
		v.Add(fmt.Sprintf("Score%d", i), strconv.Itoa(stats.Core.Score))
		v.Add(fmt.Sprintf("Goals%d", i), strconv.Itoa(stats.Core.Goals))
		v.Add(fmt.Sprintf("Assists%d", i), strconv.Itoa(stats.Core.Assists))
		v.Add(fmt.Sprintf("Saves%d", i), strconv.Itoa(stats.Core.Saves))
		v.Add(fmt.Sprintf("Shots%d", i), strconv.Itoa(stats.Core.Shots))
		i++
	}

	if _, err := http.PostForm(os.Getenv(insertURL), v); err != nil {
		return err
	}

	return nil
}
