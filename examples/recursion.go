// learngo by examples: recursion
package main

import "fmt"

func fact(n int) int {
	fmt.Println("Now n is -----> ", n)
	if n == 0 {
		return 1
	} else {
		return n*fact(n-1)
	}
}

func main() {
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println("recursion start...")
	result := fact(input)
	fmt.Println("recursion stop, the result is", result)
}