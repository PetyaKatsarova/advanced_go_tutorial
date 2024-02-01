package main

// synchronization primitives
//can be used to synchronize access of the memory, they are best
//used to communicate information between goroutines.

/* -------- ONLY READ -----------
var dataStream <-chan interface{}
dataStream := make(<-chan interface{})
-------- ONLY SEND -----------
var dataStream chan<- interface{}
dataStream := make(chan<- interface{})
---------- UNIDIRECTIONAL -----------
var dataStream chan interface{}
dataStream = make(chan interface{})

var receiveChan <-chan interface{}
var sendChan chan<- interface{}
dataStream := make(chan interface{})
 interface{} variable, which means that we can place any kind of data onto it,

// Valid statements:
receiveChan = dataStream
sendChan = dataStream
*/

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begun\n", i)
		}(i)
	}

	fmt.Println("Unblocking goroutines...")
	close(begin)
	wg.Wait()

	fmt.Println("--- EXAMPLE 2 --------")
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}

	fmt.Println("-------- EXAMPLE 3 ------------")
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")

	fmt.Println("------- EXAMPLE 4 -------")
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}

	fmt.Println("-------- EXAMPLE 5 ------")
    /*
    !! NB: To summarize, when you receive from a closed channel in Go, the operation immediately returns the zero value for the channel's type.
     In a select statement with multiple closed channels, one of the cases will be selected at random in each iteration, leading to the behavior 
     cd you observed in your code.*/
	c1 := make(chan interface{})
	close(c1) // if chans r not closed
	// they stay blocked and deadlock 
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)

}

/*Remember in “The sync Package” we discussed using the sync.Cond type to perform the same behavior. You can certainly use that, but as we’ve discussed, channels are composable, so this is my favorite way to unblock multiple goroutines at the same time.*/
