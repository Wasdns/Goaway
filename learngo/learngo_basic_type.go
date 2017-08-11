// This package records my first go lib: learngo
// learngo_basic_type.go: practicing basic types in Go
package learngo

import (
    "fmt"
)

// practicing bool format
func HypervisorJudging(isFlowVisor bool) {
	if isFlowVisor {
		fmt.Printf("The hypervisor is FlowVisor, Open Networking Lab.\n")
	} else {
		fmt.Printf("Other hypervisor.\n")
	}
}

// practicing format casting
func Int16toInt32(a16bitIntNumber int16) int32 {
	// casting int16 to int32
	a32bitIntNumber := int32(a16bitIntNumber)
	return a32bitIntNumber
}

/* complex value related functions */

// function: complex value add
func ComplexAdd(complexValueA complex128, complexValueB complex128) complex128 {
	complexAddResult := complexValueA+complexValueB
	return complexAddResult
}

// function: complex value minus
func ComplexMinus(complexValueA complex128, complexValueB complex128) complex128 {
	complexMinusResult := complexValueA-complexValueB
	return complexMinusResult
}

// function: return complex real part and virtual part
func DivideComplex128(complexValue complex128) (float64, float64) {
	complexRealPart, complexImgPart := real(complexValue), imag(complexValue)
	// Note that the type of return value is float32
	return complexRealPart, complexImgPart
} 
