package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/internal/handler"
	"github.com/octanegg/core/octane"
)

func main() {
	var (
		c = initialize()
		h = handler.NewHandler(c)
		r = mux.NewRouter()
	)

	r.HandleFunc("/events", h.GetEvents).Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", h.GetEvent).Methods(http.MethodGet)

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
