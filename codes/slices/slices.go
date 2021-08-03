package main

import (
	"fmt"
)

func main() {
	// SLICE
	// []T
	// A slice type has no specific length

	// declare a slice
	var mySlice []int
	fmt.Println(mySlice) // []

	// mySlice[0] = 1 // occurs an error, since size is unknown

	// create a slice using make function
	// dynamically-sized arrays
	// make([]T, len, cap)
	sliceWithMake := make([]int, 3, 10)
	fmt.Println(sliceWithMake)      // [0 0 0]
	fmt.Println(len(sliceWithMake)) // 3
	fmt.Println(cap(sliceWithMake)) // 5
}
