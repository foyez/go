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

	// slice
	fruits := [5]string{"banana", "pear", "apple", "orange", "peach"}
	splicedFruits := fruits[1:3]              // [pear apple]
	splicedFruits2 := fruits[2:]              // [apple orange peach]
	removeLastFruit := fruits[:len(fruits)-1] // [banana pear apple orange]
	lastFruit := fruits[len(fruits)-1]        // peach
	fmt.Println(splicedFruits, splicedFruits2, removeLastFruit, lastFruit)
	fmt.Println(len(splicedFruits)) // 2
	fmt.Println(cap(splicedFruits)) // 4 (since starts from 1 and end index is 4)

	// append
	fruitsToAdd := append(splicedFruits, "cherry", "pineapple", "guava")
	fmt.Println(splicedFruits, fruitsToAdd)             // [pear apple] [pear apple cherry pineapple guava]
	fmt.Println(len(splicedFruits), cap(splicedFruits)) // 2 4
	fmt.Println(len(fruitsToAdd), cap(fruitsToAdd)) // 5 8 (after crossing the previous capacity, the current capcity is doubled up)

	// multidimensional array
	multi := [2][3]int{{1, 2, 3}, {5, 6, 7}}
	fmt.Println(multi) // [[1 2 3] [5 6 7]]
}
