// https://go.dev/blog/pipelines

package main

import (
	"fmt"
	"sync"
)

// stage one
func gen(nums ...int) <-chan int { // return send only chan
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// stage 2, accepts receive only chan, returns send only chan
func sq(done <-chan struct{}, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            select {
            case out <- n * n:
            case <-done:
                return
            }
        }
    }()
    return out
}

/*
The main function sets up the pipeline and runs the final stage: it receives values from the second stage and prints each one,
 until the channel is closed:
*/
func main() {
	// set up the pipeline and consume the output
		done	:= make(chan struct{})
		defer close(done)
		in 		:= gen(4,2,3,7)
		// distribute the sq work across 2 go routines that both read from in
		c1 := sq(done, in)
		c2 := sq(done, in)
		// for n := range(merge(c1, c2)) {
		// 	fmt.Println(n)
		// }
		out := merge(done, c1, c2)
		fmt.Println(<- out)
}

/*
Fan-out, fan-in
Multiple functions can read from the same channel until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group
 of workers to parallelize CPU use and I/O. A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto
  a single channel that’s closed when all the inputs are closed. This is called fan-in. We can change our pipeline to run two instances of sq, each 
  reading from the same input channel. We introduce a new function, merge, to fan in the results:

  
*/

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done: return
			}
		}
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func ()  {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
 if a stage fails to consume all the inbound values, the goroutines attempting to send those values will block indefinitely:: a resource leak: goroutines consume memory and runtime resources, and heap references in goroutine stacks keep data from being garbage collected.
 Goroutines are not garbage collected; they must exit on their own.
*/

