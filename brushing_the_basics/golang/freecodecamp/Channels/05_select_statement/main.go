package main

import (
	"fmt"
	"time"
)

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
// This is signalling channel. What it means is that it just records
// the fact that whether we recieved or sent something to this channel
// It required zero memory allocation and hence it of type `struct{}`
var doneCh = make(chan struct{}) 

func main() {
	go logger()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}
	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}

	// Once all the messages are passed through the channel `logCh`, a signal 
	// is sent to the signalling channel `doneCh` indicating that we are done sending
	// messages to `logCh` channel . Its an empty struct of type struct{}
	doneCh <- struct{}{}
}

// The following is an inefficient way to do it,
// because if we don't define the closure of `logCh` (which we can do it main gorutine),
// the `logCh` will not be closed gracefully, once the main routine is finished
func loggerInefficient() {
	for entry := range logCh {
		fmt.Printf(
			"%v - [%v]%v\n", 
			entry.time.Format("2006-01-02T15:04:05"), 
			entry.severity,
			entry.message,
		)
	}
}

func logger() {
	for {
		// It's similar to switch, but the cases are usually channels
		// Since, its inside an infinite for loop, it will hit 1st case,
		// until all the values are recieved from main channel, after which
		// it hits the second case, it recieves from signalling channel (Line)
		// because of which it `break` out of the loop
		select {
		case entry := <-logCh:
			fmt.Printf(
				"%v - [%v]%v\n", 
				entry.time.Format("2006-01-02T15:04:05"), 
				entry.severity,
				entry.message,
			)
		case <-doneCh:
			break
		}
	}
}
