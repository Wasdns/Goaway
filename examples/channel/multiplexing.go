// learngo by examples: multiplexing
package main

import (
	"fmt"
	"math/rand"
)

// return random number
func rand_generator1() int {
	return rand.Int()
}

// return a channel which already receives
// several random number
func rand_generator2() chan int { // hint: not (chan, int)
	out := make(chan int)

	go func() {
		for {
			out <- rand.Int()
		}
	} ()

	return out
}

func rand_generator3() chan int {
	// a number
	rand_generator1 := rand_generator1()
	// a channel
	rand_generator2 := rand_generator2()

	out := make(chan int)

	go func() {
		for {
			out <-rand_generator1
		}
	} ()

	go func() {
		for {
			// <-(<-rand_gen2): get a number from gen2,
			// send this number to channel out
			out <-<-rand_generator2 
		}
	} ()

	return out
}

func main() {
	rand_service_handler := rand_generator3()
	fmt.Printf("%d\n", <-rand_service_handler) 
}