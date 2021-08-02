package main

import "fmt"

func main() {
	// declare array
	var arr [3]float64
	fmt.Println(arr) // [0 0 0]

	arr[1] = 23               // set element
	element := arr[1]         // read element
	fmt.Println(arr, element) // [0 23 0] 23

	// declare and initialize
	scores := [3]float64{9, 1.5, 2.2}
	fmt.Println(scores)

	// compiler figure out array length
	arrNotMax := [...]int{2, 3, 4}
	fmt.Println(arrNotMax, len(arrNotMax)) // [2 3 4] 3

	// multidimensional array
	multi := [2][3]int{{1, 2, 3}, {5, 6, 7}}
	fmt.Println(multi) // [[1 2 3] [5 6 7]]
}
