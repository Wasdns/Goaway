// learngo by examples: channel
package main

import "fmt"

func main() {
	msgChannel := make(chan string)
	go func() {
		fmt.Println("Sending one 'hello'...")
		msgChannel <- "hello"
	} ()

	go func() {
		msg := <- msgChannel
		if msg == "hello" {
			fmt.Println("Receiving one 'hello'...")
		}
	} ()

	var breakWord string
	fmt.Scanln(&breakWord)
	fmt.Println("exit")
}