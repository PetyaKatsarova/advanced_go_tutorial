package main

// import (
// 	"sync"
// 	"fmt"
// )

// type Btn struct {
// 	Clicked *sync.Cond
// }

// func subscribe(c *sync.Cond, fn func()) {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func() {
// 		wg.Done()
// 		c.L.Lock()
// 		defer c.L.Unlock()
// 		c.Wait() // wait for broadcast or signal
// 		fn()
// 	}()
// 	wg.Wait()
// }

// func main() {
// 	btn := Btn{ Clicked: sync.NewCond(&sync.Mutex{}) }
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	subscribe(btn.Clicked, func() {
// 		fmt.Println("Maximizing window.")
// 		wg.Done()
// 	})
// 	subscribe(btn.Clicked, func() {
// 		fmt.Println("Displaying annoying dialog box.")
// 		wg.Done()
// 	})
// 	subscribe(btn.Clicked, func() {
// 		fmt.Println("Mouse clikced.")
// 		wg.Done()
// 	})
// 	btn.Clicked.Broadcast()
// 	wg.Wait()
// }