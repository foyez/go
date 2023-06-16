package main

import (
	"fmt"
	"os"
)

type Student struct {
	ID   uint
	Name string
}

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

	fmt.Print("Type your name: ")
	fmt.Scan(&nameInput)
	fmt.Print("Type your age: ")
	fmt.Scan(&ageInput)
	fmt.Print("Are you a legend (true/false)? ")
	fmt.Scan(&legend)
	fmt.Printf("I'm %s. I'm %d. I'm a legend: %t\n", nameInput, ageInput, legend)

	s := Student{
		ID:   1,
		Name: "John Doe",
	}
	fmt.Printf("%s\n", "Hello")              // string
	fmt.Printf("%d\n", -34)                  // decimal
	fmt.Printf("%+d\n", 4)                   // positive decimal
	fmt.Printf("%t\n", false)                // boolean
	fmt.Printf("%f, %.2f\n", 3.1416, 3.1416) // float
	fmt.Printf("%v\n", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%T\n", s)
}
