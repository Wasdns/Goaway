// learngo by examples: worker pool
package main

import "fmt"
import "time"

func work(worderID int, jobs <-chan int, results chan<- int) {
	// if a job is available, the worker would start doing this job.
	// But if the worker is busy, i.e.the worker is working on a job,
	// he will not carry out the new job until his job has been done.
	for j := range jobs {
		fmt.Println("<=== worker", worderID, "starts work", j, "===>")

		// simulate worker works
		for t := 0; t < 3; t++ {
			time.Sleep(time.Second*1)
			fmt.Println("worker", worderID, "fork this program...")
		}
		
		fmt.Println("<=== worker", worderID, "stops work", j, "===>")
		results <- j // signals
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// three workers
	for w := 0; w < 3; w++ {
		// wait for jobs
		go work(w, jobs, results)
	}

	// five jobs
	for j := 0; j < 5; j++ {
		jobs <- j
	} 
	close(jobs)

	for i := 0; i < 5; i++ {
		<-results
	}
	close(results)
}