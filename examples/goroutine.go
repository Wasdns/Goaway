// learngo by examples: goroutines
package main

import "fmt"

func gofunc(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str, "------>", i)
	}
}

func main() {
	gofunc("Hello!")

	go gofunc("Hello Again!")

	// anonymous function
	go func(msg string) {
		for i := 0; i < 10; i++ {
			fmt.Println(msg, "----->", i)
		}
	} ("Hello, I received it.") // cannot stop here?

	var s string
	fmt.Scanln(&s)
	fmt.Println("done.")
}