// This package records my first go lib: learngo
// init.go
package learngo

import "fmt"

// the variable must be declared before being called
var CopyRight string

// The reason why the compiler would give error if I directly declare 
// CopyRight in init() is that this special function is only used to 
// give values to declared variables, not used to declare somthing.
func init() {
	// Error: CopyRight := "Wasdns"
	CopyRight = "Wasdns"
	// CopyRight must be called here:
	fmt.Printf("learngo/init.go: terrible copyright - @%s\n", CopyRight)
	fmt.Printf("\nI have experienced an interesting error.\n")
	fmt.Printf("See learngo/init.go for more information.\n\n")
}

func PrintCopyRight(copyright string) {
	fmt.Printf("\nCopyright @%s\n", copyright)
}