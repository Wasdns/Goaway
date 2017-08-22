// learngo by examples: range
package main

import "fmt"

func main() {
	fmt.Println("=====list=====")
	nums, sum := []int{9, 9, 6}, 0
	for idx1, num := range nums {
		fmt.Println("idx:", idx1, "num:", num)
		sum += num
	}
	fmt.Println("sum:", sum)
	fmt.Println("=====end======")

	fmt.Println("=====map=====")
	maps := map[string]int{"10.0.0.1/32":1, "10.0.0.2/32":2}
	for dstIP, dstPort := range maps { // hint: no index here
		fmt.Println("dstIP:", dstIP, "-----> dstPort", dstPort)
	}
	fmt.Println("=====end=====")
}