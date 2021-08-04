package main

import (
	"fmt"
	"github.com/foyez/golang/codes/toolkit/utils"
)

func main() {
	fmt.Println("Packages")
	total := utils.Add(1, 2, 3)
	fmt.Println(total)
	utils.Greeting("Mithu")
}
