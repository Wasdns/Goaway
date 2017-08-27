// learngo by examples: rate limiting
package main

import "time"
import "fmt"

func main() {
	// example A: simple rate limiter 

	// create five requests in a channel
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// rate controller, get a value per Millisecond*500
	limiter := time.Tick(time.Second*1)

	// handler requests
	for req := range requests {
		<-limiter // waitting until limter get a value
		fmt.Println("handle request", req, time.Now())
	}

	// example B: bursty rate limiter
	burstyLimiter := make(chan time.Time, 3) // a time channel

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// for maintaining burst capability and controlling the rate
	go func() { 
		for t := range time.Tick(time.Second*2) { 
			burstyLimiter <- t // burstyLimiter up to 3
		}
	} ()

	// Now simulate 5 more incoming requests. 
	// The first 3 of these will benefit from 
	// the burst capability of burstyLimiter.

	// Create 5 requests, 3 of them would got 
	// burst capability, other 2 would be limited.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}