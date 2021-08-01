package main

import "fmt"

// var variableName type = value
// can declare outside and inside of a function
var name string = "Zayan"

func main() {
	// Infer variable type
	var age = 20

	// variables without assinging value return default value
	// int: 0, float: 0.0, string: "", bool: false
	var salary int

	// variables in only function
	funcVar := "can't declare outside of a function"

	// multiple variables
	one, two := 1, "two"

	fmt.Println(name, age, salary)
	fmt.Println(funcVar)
	fmt.Println(one, two)
}
