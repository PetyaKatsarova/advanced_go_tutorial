package main
// book: concurrency in go

import (
	"sync"
	"testing"
)


// create 2 goroutines and send msg between them,a benchmark test designed to measure the time taken for context switching between two goroutines using channs
// NB !!! go test -bench=. in terminal

//cd src/gos-concurrency-building-blocks/the-sync-package/pool/ && \
// go test -benchtime=10s -bench=.

/*
This benchmark test was focused on measuring the performance of a
 context switch within the specified Go package. The main takeaways
  are the average time per operation (232.4 nanoseconds) and the
   fact that this operation does not involve any memory allocation.
    This kind of benchmarking is crucial for understanding 
	performance characteristics and optimization opportunities in
	 concurrent Go applications.
*/

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin	:= make(chan struct{})
	c 		:= make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<- begin
		for i :=0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}



// taskset -c 0 perf bench sched pipe -T
// for windows do: wsl command

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	fmt.Println("hello wo")
// 	memConsumed := func () uint64 {
// 		runtime.GC() // triggers garbage collection
// 		var s runtime.MemStats // MemStats is a struct provided by the Go runtime that holds various memory allocation statistics.
// 		runtime.ReadMemStats(&s)
// 		return s.Sys
// 	}
// 	var c <-chan interface{}
// 	var wg sync.WaitGroup
// 	noop := func ()  { wg.Done(); <-c } // never exit go routine

// 	const numGoroutines = 1e4
// 	wg.Add(numGoroutines)
// 	b4 := memConsumed()
// 	for i := numGoroutines; i > 0; i-- {
// 		go noop()
// 	}
// 	wg.Wait()
// 	after := memConsumed()
// 	fmt.Printf("%.3fkb", float64(after-b4)/numGoroutines/1000)
// }

/*
Garbage collection is a process that automatically reclaims memory 
occupied by objects that are no longer in use by the program.
 In Go, garbage collection is usually managed by the runtime,
  but it can be manually triggered using runtime.GC().
*/