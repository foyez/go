package main

import "fmt"

func main() {
	// *****************************************
	switch city := "Cumilla"; city {
	case "Dhaka", "Cumilla", "Sylhet":
		fmt.Println("You live in", city)
	default:
		fmt.Println("You're not from around here")
	}

	// *****************************************
	var age int = 30

	switch {
	case age < 18:
		fmt.Println("young")
	case age > 18 && age <= 40:
		fmt.Println("adult")
	default:
		fmt.Println("elder")
	}

	// *****************************************
	var num int = 9

	switch {
	case num != 10:
		fmt.Println("Does not equal 10")
		fallthrough // check other case after matching this case
	case num < 10:
		fmt.Println("Less than 10")
	case num > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("Is 10")
	}
}
