package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// creating an integer channel, which means only 
	// integers can be passed through this channel
	ch := make(chan int)

	wg.Add(2) // We wait for 2 goroutines

	// Channel receive gorountine
	go func() {
		// We are receiving the value from channel `ch` into `i`
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	// Channel send gorountine
	go func() {
		k := 42
		ch <- k // We are sending the value 42 to channel `ch`
		// Since we are sending a copy of the variable `k` to channel `ch`
		// we can locally make changes to `k` without affect the value which is
		// sent over through the channel
		k = 12
		wg.Done()
	}()

	wg.Wait() // Halts the main() routine until all awaited goroutines are finished
}