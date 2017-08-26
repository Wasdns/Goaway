// learngo by examples: closing channels
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Printf("R:Get one job: J%d\n", j)
			} else {
				fmt.Println("R:No more jobs...")
				done <- true
				return
			}
		}
	} ()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("S:Send one job:", i)
	}

	close(jobs)
	fmt.Println("S:Stop send jobs and using '<-done' to wait for another goroutine")

	// We await the worker using the synchronization 
	// approach we saw earlier.
	<-done
	fmt.Println("S:End of waiting.")
}