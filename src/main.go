package main

import (
	"fmt"
	"log"
)

func main() {
	// Print a welcome message
	fmt.Println("Welcome to CSC-569!")

	// Example of error handling
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// run is the main application logic
func run() error {
	fmt.Println("Application is running...")
	return nil
}
