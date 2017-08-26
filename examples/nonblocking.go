// learngo by examples: non-blocking channel operation
// Basic sends and receives on channels are blocking. 
package main

import "fmt"

func main() {
	msgs := make(chan string)
	signals := make(chan bool)

	select {
	case <-msgs:
		fmt.Println("Happy, I get one gift from Hawaii!")
	default:
		fmt.Println("I'm angry!")
	}

	msg := "Hawaii Gittar"
	// Get "fatal error: all goroutines are asleep - deadlock!" if
	// msgs <- msg
	select {
	case msgs <- msg:
		fmt.Println("Hawaii sends one gift to ...")
	default:
		fmt.Println("No, you must send one gift to him, otherwise he would be angry!")
	}

	select {
	case <-msgs:
		fmt.Println("Happy, I get one gift from Hawaii!")
	case <-signals:
		fmt.Println("Fork your guys!")
	default:
		fmt.Println("I'm angry!")
	}
}