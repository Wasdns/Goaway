// learngo by examples: timeouts
package main

import "time"
import "fmt"

func main() {
	chan1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second*2)
		chan1 <- "Add two second..."
	} ()

	select {
	case res := <-chan1:
		fmt.Println(res)
	case <-time.After(time.Second*1):
		fmt.Println("I'm angry!")
	}

	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second*1)
		chan2 <- "Add one second..."
	} ()

	select {
	case res := <-chan2:
		fmt.Println(res)
	case <-time.After(time.Second*2):
		fmt.Println("I'm angry!")
	}
}