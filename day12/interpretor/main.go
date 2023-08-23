package main

import "fmt"

func main() {
	var cmd string
	var args string
Loop:
	for {
		fmt.Printf("> ")
		fmt.Scanf("%s %s", &cmd, &args)
		switch cmd {
		case "exit":
			break Loop
		case "echo":
			fmt.Println(args)
			break
		default:
			fmt.Printf("error: %s is not a valid command!\n", cmd)
			break
		}
	}
}
