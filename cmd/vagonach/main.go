package main

import (
	"log"

	"github.com/vladimirfedunov/2chan-clone/internal/app"
)

func main() {
	application := app.NewApp()
	if err := application.Run(); err != nil {
		log.Fatalf("app stopped: %v", err)
	}
}
