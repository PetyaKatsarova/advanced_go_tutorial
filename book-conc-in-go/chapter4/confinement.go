package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan interface{}, strings <-chan string) <-chan interface{} {
	terminated := make(chan interface{})

	go func() {
		defer fmt.Println("doWork exited.")
		defer close(terminated)
		for {
			select {
			case s := <-strings:
				fmt.Println(s)
			case <-done:
				return
			}
		}
	}()
	return terminated
}

func main() {
	done 		:= make(chan interface{})
	stringsChan := make(chan string)
	terminated 	:= doWork(done, stringsChan)

	sendStrings := func(strings ...string) {
		for _, s := range strings {
			stringsChan <- s
		}
	}

	go func() {
		sendStrings("hello", "world", "tra", "la", "la")
		time.Sleep(1*time.Second)
		fmt.Println("canceling doWork goroutine...")
		close(done)
	}()
	<- terminated
	fmt.Println("Done.")
}