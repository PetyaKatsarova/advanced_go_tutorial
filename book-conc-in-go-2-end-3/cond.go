package main

import (
	"fmt"
	"sync"
	"time"
)

/*
c.Wait() in the loop is a blocking operation that prevents the queue from growing beyond 2 items. When the queue has 2 items, the loop's goroutine waits
 until an item is removed (removeFromQueue function). This mechanism ensures that the queue does  not exceed its intended maximum size, and it
  demonstrates a producer-consumer pattern where the production of new items is paused until there is space in the queue.
*/

func main() {
	c 		:= sync.NewCond(&sync.Mutex{})
	queue 	:= make([]interface{}, 0, 10)

	removeFromQueue := func (delay time.Duration)  {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		/*c.Wait: atomically unlock the mutex and suspend the execution of the goroutine. It will wait until some other goroutine calls c.Signal()
		 or c.Broadcast() on the same condition variable. Upon receiving the signal, c.Wait() automatically locks the mutex again and returns,
		  allowing the goroutine to proceed.*/
		for len(queue) == 2 { c.Wait() } // wait for the c.Signal()
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second)
		c.L.Unlock()
	}
}
/*
We also have a new method in this example, Signal. This is one of
 two methods that the Cond type provides for notifying goroutines
  blocked on a Wait call that the condition has been triggered. 
  The other is a method called Broadcast. Internally, the runtime
   maintains a FIFO list of goroutines waiting to be signaled;
    Signal finds the goroutine that’s been waiting the longest
	 and notifies that, whereas Broadcast sends a signal to al
	 l goroutines that are waiting. Broadcast is arguably the more
	  interesting of the two methods as it provides a way to
	   communicate with multiple goroutines at once. We can 
	   trivially reproduce Signal with channels (as we’ll see 
		in the section “Channels”), but reproducing the behavior 
		of repeated calls to Broadcast would be more difficult. 
		In addition, the Cond type is much more performant than 
		utilizing channels.
*/