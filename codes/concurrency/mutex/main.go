package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex

func inc(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go inc(&wg)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
