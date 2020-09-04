package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/octanegg/core/internal/config"
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
	log.Printf("Starting server on port %d\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), r))
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
