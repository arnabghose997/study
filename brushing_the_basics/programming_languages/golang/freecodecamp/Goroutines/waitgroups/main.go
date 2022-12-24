package main

import (
	"fmt"
	"sync"
)

// Read about "Green Threads"

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2) // It basically tells Golang that there are two Go routines to wait for
		go sayHello()
		go increment()
	}
	wg.Wait() // This is make sure that `main()` function doesn't exit quickly and waits for go routines to finish
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
