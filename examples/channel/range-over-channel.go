// learngo by examples: range over channels
package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "Hello"
	queue <- "Girl"
	
	close(queue)

	for item := range queue {
		fmt.Println(item)
	}
}