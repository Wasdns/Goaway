// learngo by examples: channel2
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

func main() {
	rand_service_handler := rand_generator2()
	fmt.Printf("%d\n", <-rand_service_handler) 
}