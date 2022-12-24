package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wgMutex = sync.WaitGroup{}
var counterMutex = 0
var m = sync.RWMutex{}


// This right here is a bad example of sync because, we are esentially destroying it
// by locking these go routines from outside. Notice, that they run sequentially, as if
// we never applied the concept of Go Coroutine
func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wgMutex.Add(2)
		m.RLock()
		go sayHelloMutex()
		m.Lock()
		go incrementMutex()
	}
	wgMutex.Wait()
}

// What we are doing is, basically applying a read lock on `counterMutex` (Line 21)
// which will not let any goroutine to ever make changes to it, until that lock is present
func sayHelloMutex() {
	fmt.Printf("Hello %v\n", counterMutex)
	m.RUnlock()
	wgMutex.Done()
}

// What we are doing is, basically applying a write lock on `counterMutex` (Line 23)
// which will not let any goroutine to ever read from that variable, until that lock is present
func incrementMutex() {
	counterMutex++
	m.Unlock()
	wgMutex.Done()
}
