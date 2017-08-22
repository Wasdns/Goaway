// learngo by examples: Multiple Return Values
package main

import "fmt"

func simpleCalculate(a int, b int) (int, int) {
	return a+b, a-b
}

func main() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b)
	fmt.Println("a:", a, "b:", b)
	addRes, minusRes := simpleCalculate(a, b)
	fmt.Println("a+b:", addRes, "a-b:", minusRes)
}