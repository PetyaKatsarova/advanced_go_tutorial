package main

import (
	"sync"
	"time"
	"fmt"
)
// context = scope

// example of: Starvation is any situation where a concurrent process cannot get all the resources it needs to perform work.

func main() {
	var wg sync.WaitGroup
	var sharedLock sync.Mutex
	const runtime = 1*time.Second

	greedyWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(3*time.Nanosecond)
			sharedLock.Unlock()
			count++
		}

		fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
	}

	politeWorker := func() {
		defer wg.Done()

		var count int
		for begin := time.Now(); time.Since(begin) <= runtime; {
			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			sharedLock.Lock()
			time.Sleep(1*time.Nanosecond)
			sharedLock.Unlock()

			count++
		}

		fmt.Printf("Polite worker was able to execute %v work loops.\n", count)
	}

	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
/*
FINDING A BALANCE
It is worth mentioning that the previous code example can also serve as an example of the performance ramifications of memory access synchronization.
 Because synchronizing access to the memory is expensive, it might be advantageous to broaden our lock beyond our critical sections. On the other hand,
  by doing so—as we saw—we run the risk of starving other concurrent processes.
If you utilize memory access synchronization, you’ll have to find a balance between preferring coarse-grained synchronization for performance, 
and fine-grained synchronization for fairness. When it comes time to performance tune your application, to start with, I highly recommend you constrain 
memory access synchronization only to critical sections; if the synchronization becomes a performance problem, you can always broaden the scope. It’s much
 harder to go the other way.
 So starvation can cause your program to behave inefficiently or incorrectly.
  The prior example demonstrates an inefficiency, but if you have a concurrent
   process that is so greedy as to completely prevent another concurrent process
    from accomplishing work, you have a larger problem on your hands.

We should also consider the case where the starvation is coming from outside the
 Go process. Keep in mind that starvation can also apply to CPU, memory, file
  handles, database connections: any resource that must be shared is a candidate 
  for starvation.
*/

// import (
// 	"sync/atomic"
// 	"sync"
// 	"time"
// 	"fmt"
// 	"bytes"
// )
// livelocks are a subset of a larger set of problems called starvation.
// func main() {
// 	cadence := sync.NewCond(&sync.Mutex{}) // rate of speed
// 	go func() {
// 		for range time.Tick(1*time.Millisecond) { //returns chan on which the time is sent every ms
// 			cadence.Broadcast() 
// 			/*Broadcast is signaling all waiting goroutines that the condition they are waiting 
// 			for might be true, so they should wake up and check the condition.*/
// 		}
// 	}()

// 	takeStep := func() {
// 		cadence.L.Lock()
// 		cadence.Wait()
// 		cadence.L.Unlock()
// 	}

// /*bytes.Buffer implements the io.Reader, io.Writer, io.ByteWriter, and io.RuneWriter interfaces, making it
//  a versatile tool for both reading from and writing to byte slices. dir is direction */
// 	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
// 		fmt.Fprintf(out, "%v", dirName)
// 		atomic.AddInt32(dir, 1)
// 		takeStep()
// 		if atomic.LoadInt32(dir) == 1 {
// 			fmt.Fprint(out, ". Sucess!")
// 			return true
// 		}
// 		takeStep()
// 		atomic.AddInt32(dir, -1)
// 		return false
// 	}

// 	var left, right int32
// 	tryLeft := func(out *bytes.Buffer) bool {  return tryDir(" left", &left, out)   }
// 	tryRight := func(out *bytes.Buffer) bool { return tryDir(" right", &right, out) }

// 	walk := func (walking *sync.WaitGroup, name string)  {
// 		var out bytes.Buffer
// 		defer func () { fmt.Println(out.String()) }()
// 		defer walking.Done()
// 		fmt.Fprintf(&out, "%v is trying to scoot:", name)
// 		for i := 0; i < 5; i++ {
// 			if tryLeft(&out) || tryRight(&out) { return }
// 		}
// 		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
// 	}
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go walk(&wg, "Alice")
// 	go walk(&wg, "Barbie")
// 	go walk(&wg, "TinkieWinkie")
// 	wg.Wait()
// }

// deadlock example: 
// import (
// 	"sync"
// 	"time"
// 	"fmt"
// )
// type val struct {
//     mu    sync.Mutex
//     value int
// }

// func main() {
// 	var wg sync.WaitGroup
// 	printSum := func(v1, v2 *val) {
// 		defer wg.Done()
// 		v1.mu.Lock()
// 		defer v1.mu.Unlock()

// 		time.Sleep(2*time.Second) // triggers the deadlock
// 		v2.mu.Lock()
// 		defer v2.mu.Unlock()

// 		fmt.Printf("sum=%v\n", v1.value + v2.value)
// 	}

// 	var a, b val
// 	wg.Add(2)
// 	go printSum(&a, &b) // deadlock: both a and b are locked
// 	go printSum(&b, &a)
// 	wg.Wait()
// }
