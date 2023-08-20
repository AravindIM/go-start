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

// Hellos returns map associating names of people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {
	// create map
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
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
