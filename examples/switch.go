// learngo by examples: switch
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { // current time
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	t := time.Now()
	switch { // no switch entity
	case t.Hour() < 12: // logic
		fmt.Println("beforenoon")
	case t.Hour() >= 12:
		fmt.Println("afternoon")
	}

	whatAmI := func(i interface{}) { // function has the highest priority
		fmt.Println("switch case example")
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool.")
		case int:
			fmt.Println("I'm an int.")
		default:
			fmt.Println("Interesting! %T", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("happy birthday")
}