package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/octanegg/core/deprecated"
	"github.com/octanegg/core/internal/admin"
	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/internal/handler"
	"github.com/octanegg/core/octane"
	"github.com/octanegg/racer"
	"github.com/octanegg/slimline"
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
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	return handler.New(
		octane.New(db),
	)
}

func newAdminHandler() admin.Handler {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	dprctd, err := deprecated.New()
	if err != nil {
		log.Fatal(err)
	}

	return admin.New(
		octane.New(db),
		racer.New(os.Getenv(config.EnvAuthToken)),
		slimline.New("core", "rocketleague"),
		dprctd,
	)
}
