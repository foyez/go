package main

import (
	"errors"
	"fmt"
	"os"
)

func isGreaterThanTen(num int) error {
	if num < 10 {
		return errors.New("something bad happened")
	}

	return nil
}

func openFile() error {
	f, err := os.Open("missingFile.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func main() {
	num := 9

	if err := isGreaterThanTen(num); err != nil {
		fmt.Println(fmt.Errorf("%d is not greater than ten", num))
		// panic(err)
		// log.Fatalln(err)
	}

	if err := openFile(); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	// RECOVER
	defer recoverFromPanic()
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 2 {
			panic("PANIC!")
		}
	}
}
