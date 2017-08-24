// learngo by examples: error
package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cannot work with 42!")
	} else {
		return arg+3, nil
	}
}

type argError struct {
	arg int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s\n", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot work with 42!"}
	} else {
		return arg+3, nil
	}
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e == nil {
			fmt.Println("f1 worked:", r)
		} else {
			fmt.Println("f1 failed:", e)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e == nil {
			fmt.Println("f2 worked:", r)
		} else {
			fmt.Println("f2 failed:", e.Error())
		}
	}

	// If you want to programmatically use the data in a custom error,
	// you’ll need to get the error as an instance of the custom error 
	// type via type assertion.
	// In other words, using a error format to parse the error.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}