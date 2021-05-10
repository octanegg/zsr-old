package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/octanegg/zsr/internal/cache"
	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/internal/handler"
	"github.com/octanegg/zsr/octane"
)

func main() {
	o, err := octane.New(os.Getenv(config.EnvURI))
	if err != nil {
		log.Fatal(err)
	}

	c := cache.New(os.Getenv(config.EnvCacheURI))
	r := routes(handler.New(o, c))

	http.Handle("/", r)
	log.Printf("Starting server on port %d\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), r))
}
