// learngo by examples: select
package main

import (
	"fmt" 
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	} ()
	go func() {
		time.Sleep(time.Second)
		c2 <- "second"
	} ()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("channel c1 receives:", msg1)
		case msg2 := <-c2:
			fmt.Println("channel c2 receives:", msg2)
		}
	}
}