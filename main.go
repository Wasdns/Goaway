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

	fmt.Printf("\n")

	// function: int add and minus,
	// in file learngo/learngo_calculate_func.go
	var calculateRes int = 0
	calculateRes = learngo.Sum(1, 2)
	fmt.Printf("1+2=%d\n", calculateRes)
	calculateRes = learngo.Minus(1, 2)
	fmt.Printf("1-2=%d\n", calculateRes)

	fmt.Printf("\n")

	// print global variable,
	// in file learngo/learngo_var.go
	fmt.Println(learngo.StringGoaway)

	fmt.Printf("\n")

	// print goos,
	// in file example/goos.go
	example.PrintOperatingSystem()

	fmt.Printf("\n")

	// variable swap
	swapTestA, swapTestB := learngo.Sum(1, 3), learngo.Minus(1, 3)
	fmt.Printf("Before swaping: swapTestA=%d, swapTestB=%d\n", swapTestA, swapTestB)
	swapTestA, swapTestB = swapTestB, swapTestA
	fmt.Printf("After swaping: swapTestA=%d, swapTestB=%d\n", swapTestA, swapTestB)

	fmt.Printf("\n")

	// print "terrible" copyright @Wasdns,
	// in file learngo/init.go
	learngo.PrintCopyRight(learngo.CopyRight)

	fmt.Printf("\n")

	// print hypervisor type,
	// in file learngo/learngo_basic_type.go
	fmt.Printf("isFlowVisor Judging...\n")
	learngo.HypervisorJudging(true)
	fmt.Printf("isnotFlowVisor Judging...\n")
	learngo.HypervisorJudging(false)

	fmt.Printf("\n")

	// int16 to int32 casting
	var a16bitIntNumber int16 = 2017
	fmt.Printf("16bit Int Number is: %16d\n", a16bitIntNumber)
	var a32bitIntNumber int32 = learngo.Int16toInt32(a16bitIntNumber)
	fmt.Printf("32bit Int Number is: %32d\n", a32bitIntNumber)

	fmt.Printf("\n")

	// complex value calculating
	// Note that the default type of complex value is complex128
	complexValueA := 5+10i
	complexValueB := 6-3i
	fmt.Printf("complexValueA:%v, complexValueB:%v\n", complexValueA, complexValueB)
	// Question: get "use of builtin complex not in function call" if declare complex here:
	// var complexAddResult complex
	// var complexMinusResult complex
	fmt.Printf("complexValueA+complexValueB=%v\n", learngo.ComplexAdd(complexValueA, complexValueB))
	fmt.Printf("complexValueA-complexValueB=%v\n", learngo.ComplexMinus(complexValueA, complexValueB))
	complexValueC := 11+23i
	fmt.Printf("complexValueC:%v\n", complexValueC)
	valueCRealPart, valueCImgPart := learngo.DivideComplex128(complexValueC)
	fmt.Printf("complexValueC_real:%f, complexValueC_img:%f\n", valueCRealPart, valueCImgPart)

	fmt.Printf("\n")
}
