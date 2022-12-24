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

	n_loop := 5
	wg.Add(1+n_loop) // We wait for the recive routine

	// Channel receive gorountine. Note that it expects to recieve
	// only a single message
	go func() {
		// We are receiving the value from channel `ch` into `i` "once"
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	// Channel send gorountine
	// We are attempting to send multiple messages 
	// What will happen is: in the first iteration would be successful in sending 42
	// Once the channel recieves it, the channel is basically done right there and
	// it doesn't expect any more messages to recieve
	// But since, we are trying to send multiple messages (more than once) it will end up
	// in deadlock
	for j := 0; j < n_loop; j++ {
		go func() {
			k := 42
			ch <- k // We are sending the value 42 to channel `ch`
			wg.Done()
		}()
	}

	wg.Wait() // Halts the main() routine until all awaited goroutines are finished
}