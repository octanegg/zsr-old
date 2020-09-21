package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/internal/deprecated"
	"github.com/octanegg/core/internal/handler"
	"github.com/octanegg/core/octane"
)

func main() {
	var (
		r = routes(newHandler(), newAdminHandler())
	)

	http.Handle("/", r)
	log.Printf("Starting server on port %d\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), r))
}

func newHandler() handler.Handler {
	o, err := octane.New(os.Getenv(config.EnvURI))
	if err != nil {
		log.Fatal(err)
	}

	return handler.New(o)
}

func newAdminHandler() deprecated.Handler {
	dprctd, err := deprecated.New()
	if err != nil {
		log.Fatal(err)
	}

	o, err := octane.New(os.Getenv(config.EnvURI))
	if err != nil {
		log.Fatal(err)
	}

	return deprecated.NewHandler(dprctd, o)
}
