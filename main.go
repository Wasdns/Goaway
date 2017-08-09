package main

import (
	"./example"
	"./learngo"
	"fmt"
)

// WeekTime: const value
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

func main() {
	// function: PrintHelloGo,
	// in file learngo/helloworld.go
	learngo.PrintHelloGo()

	// function: int add and minus,
	// in file learngo/learngo_calculate_func.go
	var calculateRes int = 0
	calculateRes = learngo.Sum(1, 2)
	fmt.Printf("1+2=%d\n", calculateRes)
	calculateRes = learngo.Minus(1, 2)
	fmt.Printf("1-2=%d\n", calculateRes)

	// print global variable,
	// in file learngo/learngo_var.go
	fmt.Println(learngo.StringGoaway)

	// print goos,
	// in file example/goos.go
	example.PrintOperatingSystem()
}
