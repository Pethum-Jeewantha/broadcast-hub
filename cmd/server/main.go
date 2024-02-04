package main

import (
	"broadcast-hub/internal/app"
	"log"
)

func main() {
	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
