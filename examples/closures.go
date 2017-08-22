// learngo by examples: closures
package main

import "fmt"

func learnClosures() func() int { // return a closure in terms of function
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	func1 := learnClosures()
	for i := 0; i < 10; i++ {
		fmt.Println("No.", i+1, "-----> function result:", func1())
	}

	func2 := learnClosures()
	fmt.Println("New start -----> function result:", func2())
}