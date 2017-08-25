// learngo by examples: Channel Directions
package main

import "fmt"

// pings is a channel used to receive msg
func ping(pings chan<- string, msg string) { 
	pings <- msg
	fmt.Println("channel pings <-----", msg)
}

// In this function(pong)...
// pings is a channel used to send msg,
// pongs is a channel used to receive msg.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	fmt.Println("channel pongs <-----", msg)
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "fork you")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}