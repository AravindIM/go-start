package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("AIM")
	fmt.Println(message)
}
