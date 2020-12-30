package main

import (
	"log"
	"os"
)

func main() {
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		log.Fatal("$PORT must be set")
	}

	userid := os.Getenv("USERID")
	if userid == "" {
		log.Fatal("$USERID must be set")
	}

	password := os.Getenv("PASSWORD")
	if password == "" {
		log.Fatal("$PASSWORD must be set")
	}

	initCache(teams)
	handleRequests(port, userid, password)
}
