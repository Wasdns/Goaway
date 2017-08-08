package main

import (
	"./helloworld"
	"fmt"
)

const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

func main() {
	// function: PrintHelloGo, in file helloworld.go
	learngo.PrintHelloGo()
	fmt.Println("1+2=")
	fmt.Println(learngo.Sum(1, 2))
	fmt.Println("1-2=")
	fmt.Println(learngo.Minus(1, 2))
}
