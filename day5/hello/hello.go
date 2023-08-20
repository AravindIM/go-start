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

	// A slice of names
	names := []string{"Athos", "Porthos", "Aramis", "d'Artagnan"}

	// Request a greeting message
	messages, err := greetings.Hellos(names)

	// Print error into log on error
	if err != nil {
		log.Fatal(err)
	}

	// Print the greeting if no error
	for _, name := range names {
		message := messages[name]
		fmt.Println(message)
	}
}
