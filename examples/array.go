// learngo by examples: arrays
package main

import "fmt"

func main() {
	var array [5]int
	array[4] = 10
	fmt.Println("Array:", array)
	fmt.Println("Length:", len(array))

	brray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("NewArray:", brray)

	var doubleArray [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			doubleArray[i][j] = i+j
		}
	}
	fmt.Println(doubleArray)
}