package main

import (
	"go-api-boilerplate/utils/database"
	"go-api-boilerplate/utils/server"
	"log"
)

func main() {
	_, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatal("failed to migragte database schema: %w", err)
	}

	server.LaunchWebServer()
}
