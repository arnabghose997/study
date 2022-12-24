package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// The best practice is to assing goroutines single roles, in terms of channel direction
// It could be either recieve or send

func main() {
	ch := make(chan int, 3) // Second argument here is the channel buffer
	wg.Add(2)

	// Recieve only channel is defined
	go func(ch <-chan int) {
		// Now Range does know when a channel is closed from the sending end
		for i := range ch {
			fmt.Println(i)
		}

		// Below is an alternate way to think about the `range` over channels
		/*
		for {
			// `ok` checks if the channel is open
			if i, ok := <- ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		*/
		wg.Done()
	}(ch)

	// Send only channel is defined
	go func(ch chan<- int) {
		ch <- 42
		ch <- 23
		ch <- 67
		// This is done to basically indicate an EOF for the channels
		// This means that channel will be close only after 67
		// This is done to avoid deadlock on the recieving end
		// If we hadn't done this, channel would closing after each of the sends shown above
		close(ch) 
		wg.Done()
	}(ch)

	wg.Wait()
}