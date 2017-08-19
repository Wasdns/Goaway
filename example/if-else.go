// learngo by examples: if-else
package main

import "fmt"

func main() {
	if j := 1; j != 2 {
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}

	i := 2
	if j := 1; i==2 && j==1 { // only enable one ';'
		fmt.Println("if")
	} else {
		fmt.Println("else")
	}
}