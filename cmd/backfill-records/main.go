package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/helper"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	o, err := octane.New(os.Getenv(config.EnvURI))
	if err != nil {
		log.Fatal(err)
	}

	events, err := o.Events().Find(nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	for i, event := range events {
		start := time.Now()
		slug := event.(octane.Event).Slug
		raw, err := o.Statlines().Find(bson.M{
			"game.match.event.slug": slug,
		}, nil, nil)
		if err != nil {
			log.Fatal(err)
		}

		var statlines []*octane.Statline
		var records []interface{}
		for _, statline := range raw {
			s := statline.(octane.Statline)
			statlines = append(statlines, &s)
		}
		for _, record := range helper.StatlinesToRecords(statlines) {
			records = append(records, record)
		}

		if len(records) == 0 {
			continue
		}

		fmt.Println("FMT", i, event.(octane.Event).Slug, len(records), time.Since(start))
		start = time.Now()

		// if _, err := o.Records().Delete(bson.M{"game.match.event.slug": slug}); err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println("DLT", i, event.(octane.Event).Slug, len(records), time.Since(start))
		// start = time.Now()

		if _, err := o.Records().Insert(records); err != nil {
			log.Fatal(err)
		}

		fmt.Println("WRT", i, event.(octane.Event).Slug, len(records), time.Since(start))

	}

}
