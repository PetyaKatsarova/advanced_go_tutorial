package main

import (
	"fmt"
	"sync"
)

func main() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v\n", id)
	}

	const numGreetings = 5
	var wg sync.WaitGroup
	wg.Add(numGreetings)
	for i := 0; i < numGreetings; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
