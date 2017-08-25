// learngo by examples: Channel Buffering
package main 

import "fmt"

func main() {
	msg := make(chan string, 2)
	msg <- "SIGCOMM"
	msg <- "NSDI"
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}