package main

import "fmt"

func main() {
	var day string
	var week string
	fmt.Scanf("%s", &day)

	switch day {
	case "mon", "monday":
		week = "week day"
		break
	case "tue", "tueday":
		week = "week day"
		break
	case "wed", "wednesday":
		week = "week day"
		break
	case "thu", "thursday":
		week = "week day"
		break
	case "fri", "friday":
		week = "week day"
		break
	case "sat", "saturday":
		week = "weekend"
		break
	case "sun", "sunday":
		week = "weekend"
		break
	default:
		fmt.Printf("Error: invalid weekend\n")
		return
	}

	fmt.Printf("%v\n", week)
}
