// learngo by examples: Atomic Counters

// It is very interesting and funny.

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {
	var ops uint64 = 0
	var tmp int = 0

	// 50*700+ ~= 35000+
	for i := 0; i < 5; i++ { 
		fmt.Println("Add new i:", i) // #1st
		
		go func() { 
			id := i
			for { // no for: 50; has for: 700+
				atomic.AddUint64(&ops, 1) 
				time.Sleep(time.Millisecond)
				tmp = tmp+1
				fmt.Print("Goroutine ", id, " add 1 to tmp\n") // #4th
			}
			// <--- will never get here --->
		} () 

		// run and see what will happen
		time.Sleep(time.Second)

		fmt.Print("\n-----> ins2 - goroutine ", i, " tmp:", tmp, "\n") // #2nd
	}

	fmt.Println("<===== tmp before sleeping:", tmp, " =====>") // #3rd
	time.Sleep(time.Second) // sleep for running the goroutine
	fmt.Println("<===== tmp after sleeping:", tmp, " =====>") // #5th

	opsfinal := atomic.LoadUint64(&ops)
	fmt.Println("Execution times:", opsfinal)
}