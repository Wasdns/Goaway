// learngo by examples: for
package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i+1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("It's a for loop.")
		break
	}

	for i = 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Wasdns")
			continue
		} else {
			fmt.Println(i)
		}
	}
}