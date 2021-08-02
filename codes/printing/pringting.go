package main

import (
	"fmt"
	"os"
)

func main() {
	// Print
	fmt.Println("Hello, World")
	fmt.Print("Hello without newline")
	fmt.Print("Hello again\n")
	fmt.Printf("I'm %s. I'm %d. I'm a %t legend\n", "Zohan", 1, true)

	// Fprint
	const name, age = "Farah", 12
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}

	fmt.Println(n, "bytes written")

	// Scan
	var nameInput string
	var ageInput int
	var legend bool

	fmt.Println("Type your name:")
	fmt.Scan(&nameInput)
	fmt.Println("Type your age:")
	fmt.Scan(&ageInput)
	fmt.Println("Are you a legend (true/false)?")
	fmt.Scan(&legend)
	fmt.Printf("I'm %s. I'm %d. I'm a legend: %t", nameInput, ageInput, legend)
}
