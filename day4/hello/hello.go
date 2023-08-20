package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("AIM")

	// Print error into log on error
	if err != nil {
		log.Fatal(err)
	}

	// Print the greeting if no error
	fmt.Println(message)
}
