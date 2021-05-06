package helper

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/octanegg/zsr/octane"
)

const idLength = 4

func EventSlug(event *octane.Event) string {
	id := event.ID.Hex()
	return slug.Make(fmt.Sprintf("%s-%s", id[len(id)-idLength:], event.Name))
}

func MatchSlug(match *octane.Match) string {
	id := match.ID.Hex()
	blue, orange := "TBD", "TBD"

	if match.Blue != nil {
		blue = match.Blue.Team.Team.Name
	}

	if match.Orange != nil {
		orange = match.Orange.Team.Team.Name
	}

	return slug.Make(fmt.Sprintf("%s-%s-vs-%s", id[len(id)-idLength:], blue, orange))
}

func PlayerSlug(player *octane.Player) string {
	id := player.ID.Hex()
	return slug.Make(fmt.Sprintf("%s-%s", id[len(id)-idLength:], player.Tag))
}

func TeamSlug(team *octane.Team) string {
	id := team.ID.Hex()
	return slug.Make(fmt.Sprintf("%s-%s", id[len(id)-idLength:], team.Name))
}
