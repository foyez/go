package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x = 10
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(float64(x) + 5.5)
}
