// learngo by examples: timer
package main

import "fmt"
import "time"

func main() {
	timer1 := time.NewTimer(time.Second*2)
	<-timer1.C
	fmt.Println("Timer1 expired.")

	timer2 := time.NewTimer(time.Second*5)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired.")
	} ()
	stop2 := timer2.Stop() // stop the timer2
	if stop2 {
		fmt.Println("stop Timer2.")
	}
}