// learngo by examples: pointers
package main

import "fmt"

func fakeUpdatePointValue(pointValue int, modifyValue int) {
	pointValue = modifyValue
}

func updatePointerValue(point *int, modifyValue int) {
	*point = modifyValue
}

func main() {
	value := 100
	fmt.Println("Oringinal value is", value)

	fakeUpdatePointValue(value, value*100)
	fmt.Println("After fake updating, value is", value)

	updatePointerValue(&value, value*100)
	fmt.Println("After updating, value is", value)

	fmt.Println("\n===============================================")
	fmt.Println("Note that when you modify the value, you should")
	fmt.Println("pass the address of value to the function, such")
	fmt.Println("as the address of 'value' here is", &value)
	fmt.Println("and this address is stored in a pointer: &value")
	fmt.Println("===============================================\n")
}