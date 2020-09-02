package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/octanegg/core/internal/handler"
	"github.com/octanegg/core/octane"
)

const (
	port = 8080
)

func main() {
	var (
		c = initialize()
		h = handler.NewHandler(c)
		r = routes(h)
	)

	http.Handle("/", r)
	log.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
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
