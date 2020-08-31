package main

import "log"

import "github.com/octanegg/core/octane"

func main() {
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	client := octane.NewClient(db)

	if err = client.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected!")
}
