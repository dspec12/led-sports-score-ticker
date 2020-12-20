package main

import (
	"log"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	handleRequests(port)
}
