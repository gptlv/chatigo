package main

import (
	"log"

	"github.com/gptlv/chatigo/server/db"
)

func main() {

	_, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("DB connection error: %s", err)
	}

}
