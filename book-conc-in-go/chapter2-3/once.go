package main

// import (
// 	"fmt"
// 	"sync"
// )
// // grep -ir sync.Once $(go env GOROOT)/src |wc -l
// // to check how often stadlib go uses sync.Once
// // in powershell: (Get-ChildItem -Recurse -Path (go env GOROOT) -Include *.go | Select-String -Pattern "sync.Once" -SimpleMatch).Count


// func main() {
// 	var count int

// 	increment := func() {
// 		count++
// 	}
// 	var once sync.Once
// 	var wg sync.WaitGroup
// 	wg.Add(100)
// 	for i := 0; i < 100; i++ {
// 		go func ()  {
// 			defer wg.Done()
// 			once.Do(increment) // will only do 1 call to the func increment ever, even in a loop!
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("Count is ", count)

// 	fmt.Println("example 2: ")
// 	var onceA, onceB sync.Once
// 	var initB func()
// 	initA := func ()  {	onceB.Do(initB) }
// 	initB  = func ()  {	onceA.Do(initA) }
// 	onceA.Do(initA)
// }
// // ------------- SYNC.ONCE --------------------
// /*
// sync.Once is a type that utilizes some sync primitives internally to ensure that only one call to Do ever calls the function passed in—even on different
//  goroutines. This is indeed because we wrap the call to increment in a sync.Once Do method.
// */