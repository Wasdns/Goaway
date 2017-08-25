// learngo by examples: toyChannel
package main

import "fmt"

func main() {
	goChannel := make(chan string)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Sending one request...")
			goChannel <- "Client Request..."
		}
	} ()

	go func() {
		for j := 0; j < 1000; j++ {
			msg := <-goChannel
			if msg == "Client Request..." {
				fmt.Println("Get one request, give a reply...")
			}
		}
	} ()

	var breakWord string
	fmt.Scanln(&breakWord)
	fmt.Println("exit")
}