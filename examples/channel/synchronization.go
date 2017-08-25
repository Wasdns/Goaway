// learngo by examples: Channel Synchronization
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working for several seconds...")
	time.Sleep(5*time.Second)
	fmt.Println("done")
	// send a msg to notice that the work has been done
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done) 

	// If you removed the <- done line from this program, 
	// the program would exit before the worker even started.
	if <-done {
		fmt.Println("Synchronization")
	}
}