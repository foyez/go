package main

import (
	"fmt"
)

// *****************************************
// BASIC FUNCTION (RETURN NOTHING)
// *****************************************
// func printAge(age int) {
// 	fmt.Println(age)
// }

// *****************************************
// FUNCTION (RETURN)
// *****************************************
// func printAge(age int) int {
// 	return age
// }

// *****************************************
// FUNCTION (RETURN MULTIPLE)
// *****************************************
// func printAge(age int) (string, int) {
// 	return "Foyez", age
// }

// *****************************************
// RETURN MULTIPLE NAMED VALUES
// *****************************************
// func printAge(age1, age2 int) (ageOfBob, ageOfSally int) {
// 	ageOfBob = age1
// 	ageOfSally = age2
// 	return
// }

// *****************************************
// UNKNOWN NUMBER OF ARGS
// *****************************************
func average(ages ...int) float64 {
	total := 0

	for _, age := range ages {
		total += age
	}

	return float64(total) / float64(len(ages))
}

func main() {
	fmt.Println(average(10, 20, 32))
}
