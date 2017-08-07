package main

import (
	"./helloworld"
	"fmt"
)

func main() {
	// function: PrintHelloGo, in file helloworld.go
	learngo.PrintHelloGo()
	fmt.Println("1+2=")
	fmt.Println(learngo.Sum(1, 2))
	fmt.Println("1-2=")
	fmt.Println(learngo.Minus(1, 2))
}
