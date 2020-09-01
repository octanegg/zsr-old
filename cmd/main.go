package main

import (
	"log"
	"net/http"

	"github.com/octanegg/core/internal/handler"
	"github.com/octanegg/core/octane"
)

func main() {
	var (
		c = initialize()
		h = handler.NewHandler(c)
		r = routes(h)
	)

	http.Handle("/", r)
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func initialize() octane.Client {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	c := octane.NewClient(db)
	if err = c.Ping(); err != nil {
		log.Fatal(err)
	}

	return c
}
