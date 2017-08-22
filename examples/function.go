// learngo by examples: functions
package main

import "fmt"

func aPlusB(a int, b int) int {
	result := a+b
	return result
}

func main() {
	var a, b int
	fmt.Scanf("%d%d", &a, &b) // note the '&'
	aPlusb := aPlusB(a, b)
	fmt.Println("a:", a, "b:", b, "a+b=", aPlusb)
}