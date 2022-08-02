package main

import (
	"log"

	"github.com/ehrktia/samosa"
)

func main() {
	if err := samosa.Execute(); err != nil {
		log.Fatalf("failed to start samosa: %v\n", err)
	}
}
