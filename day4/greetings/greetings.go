package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello return greeting for a person
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value
	// that embeds the name in a greeting
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns randomly selected message formats
func randomFormat() string {
	// Slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return message format with random index
	return formats[rand.Intn(len(formats))]
}
