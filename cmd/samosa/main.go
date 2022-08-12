package main

import (
	"log"

	"github.com/burntcarrot/samosa"
)

func main() {
	if err := samosa.Execute(); err != nil {
		log.Fatalf("failed to start samosa: %v\n", err)
	}
}
