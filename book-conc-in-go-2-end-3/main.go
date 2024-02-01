package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	var wg sync.WaitGroup
// 	for _, val := range []string{"hi", "greetings", "g'day"} {
// 		wg.Add(1)
// 		go func ()  {
// 			defer wg.Done()
// 			fmt.Println(val)	
// 		}()
// 	}
// 	wg.Wait()
// }

/*
Usually on my machine, the loop exits before any goroutines begin
 running, so salutation is transferred to the heap holding a 
 reference to the last value in my string slice, “good day.” 
 And so I usually see “good day” printed three times. The proper
  way to write this loop is to pass a copy of salutation into the
   closure so that by the time the goroutine is run, it will be
    operating on the data from its iteration of the loop:
*/
func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) //pass in the current iteration’s variable to the closure. A copy of the string struct is made, thereby ensuring that when the goroutine is run, we refer to the proper string.
	}
	wg.Wait()
}
