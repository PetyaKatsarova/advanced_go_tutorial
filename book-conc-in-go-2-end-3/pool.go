package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance")
			return struct{}{}
		},
	}

	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()

	/* -------- EXAMPLE 2 ---------
	So why use a pool and not just instantiate objects as you go? Go has a garbage collector, so the instantiated objects 
	will be automatically cleaned up. What’s the point? Consider this example:
	
	sync.Pool is a pool of objects that can be reused to manage and allocate memory efficiently in concurrent programs.
	
	*/
	fmt.Println("----- EXAMPLE 2 ---------")
	var numCalcsCreated int
	calcPool := &sync.Pool { // is on object
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}
	// seed the pool with 4kb
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024*1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func ()  {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}()
	}
	wg.Wait()
	fmt.Printf("%d calculators were created\n", numCalcsCreated)
}