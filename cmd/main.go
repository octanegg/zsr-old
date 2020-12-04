package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/internal/deprecated"
	"github.com/octanegg/zsr/internal/handler"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/stats"
)

func main() {
	o, err := octane.New(os.Getenv(config.EnvURI))
	if err != nil {
		log.Fatal(err)
	}

	dprctd, err := deprecated.New()
	if err != nil {
		log.Fatal(err)
	}

	r := routes(handler.New(o, stats.New(o.Statlines())), deprecated.NewHandler(dprctd, o))
	http.Handle("/", r)
	log.Printf("Starting server on port %d\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), r))
}
