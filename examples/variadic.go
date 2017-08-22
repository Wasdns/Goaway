// learngo by examples: Variadic Functions
package main

import "fmt"

func variadicAdd(numbers ...int) int {
	fmt.Println("numbers:", numbers)
	result := 0
	for idx, number := range numbers {
		fmt.Println("idx:", idx, "-----> number:", number)
		result += number
	}
	return result
}

func main() {
	fmt.Println("\n===first===")
	result := variadicAdd(1, 2, 3)
	fmt.Println("calculate result:", result)

	fmt.Println("\n===second===")
	result = variadicAdd(1, 4, 5)
	fmt.Println("calculate result:", result)

	fmt.Println("\n===third===")
	numbers := []int{9, 5, 2, 6, 9, 3, 3, 5, 8}
	result = variadicAdd(numbers...)
	fmt.Println("calculate result:", result)
}