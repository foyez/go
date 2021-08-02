package main

import (
	"fmt"
)

func main() {
	// *****************************************
	// BASIC FOR LOOP
	// *****************************************
	fmt.Println("Basic for loop")
	for i := 1; i <= 5; i++ {
		fmt.Print(i)
	}

	// *****************************************
	// SIMILAR TO WHILE LOOP
	// *****************************************
	fmt.Println("\nSimilar to while loop")
	j := 1

	for j <= 5 {
		fmt.Print(j)
		j++
	}

	// *****************************************
	// INFINITE LOOP
	// *****************************************
	fmt.Println("\nInfinite loop")
	num := 1

	for {
		num = num + 2

		if num == 7 {
			continue
		}

		fmt.Print(num)

		if num == 11 {
			break
		}
	}

	// ********************************************
	// BASIC FOR LOOP ITERATION (STRING, ARRAY,...)
	// ********************************************
	fmt.Println("\nBasic for loop iteration")
	var name = "Farah"

	for i := 0; i < len(name); i++ {
		fmt.Println("Letter:", string(name[i]))
	}

	// *****************************************
	// STRING ITERATION
	// *****************************************
	fmt.Println("\nString iteration")
	var myCity = "কুমিল্লা"

	for index, letter := range myCity {
		if index%2 == 0 {
			fmt.Printf("Index: %d, Letter:%#U\n", index, letter)
		}
	}

	// *****************************************
	// SLICE OR ARRAY ITERATION
	// *****************************************
	fmt.Println("\nSlice or Array iteration")
	cities := []string{"Dhaka", "Cumilla"}

	for _, city := range cities {
		fmt.Printf("%s ", city)
	}

	// *****************************************
	// MAP ITERATION
	// *****************************************
	fmt.Println("\nMap iteration")
	results := map[string]float64{
		"Farah":   3.4,
		"Laaibah": 3.3,
		"Zayan":   3.5,
	}

	for key, value := range results {
		fmt.Println(key, value)
	}

	// *****************************************
	// CHANNEL ITERATION
	// *****************************************
	fmt.Println("\nChannel iteration")

	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
