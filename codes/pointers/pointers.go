package main

import (
	"fmt"
	"strings"
)

type Coordinates struct {
	X, Y float64
}

func upperName(name *string) {
	*name = strings.ToUpper(*name)
}

func updateCoordinates(c Coordinates) {
	c.X = 200
}

func updateCoordinatesWithPtr(c *Coordinates) {
	c.X = 200
}

func main() {
	var name string
	var namePtr *string // var pointerVar *type

	fmt.Println(name)    // ""
	fmt.Println(namePtr) // <nil>

	name = "Cumilla"
	namePtr = &name          // read variable address - &pointerVar
	var nameValue = *namePtr // read variable value - *pointerVar

	fmt.Println(name)      // Cumilla
	fmt.Println(namePtr)   // 0xc00009c050
	fmt.Println(nameValue) // Cumilla

	// ******************************************
	// Pass by Reference
	// ******************************************
	n := "Chayon"
	upperName(&n)
	fmt.Println(n)

	// ******************************************
	// Pointer with Structs
	// ******************************************
	var c = Coordinates{X: 10, Y: 20}

	updateCoordinates(c)
	fmt.Println(c) // {10 20}

	updateCoordinatesWithPtr(&c)
	fmt.Println(c) // {200 20}
}
