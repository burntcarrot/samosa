package main

import (
	"log"

	"github.com/burntcarrot/samosa/command"
)

func main() {
	if err := command.Execute(); err != nil {
		log.Fatalf("failed to start samosa: %v\n", err)
	}
}
