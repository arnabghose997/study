package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// The best practice is to assing goroutines single roles, in terms of channel direction
// It could be either recieve or send

func main() {
	ch := make(chan int)
	wg.Add(2)

	// Recieve only channel is defined
	go func(ch <-chan int) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch)

	// Send only channel is defined
	go func(ch chan<- int) {
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()
}