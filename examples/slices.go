// learngo by examples: slices
package main

import "fmt"

func main() {
	// initialize a slice
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0], s[1], s[2] = "alice", "bob", "cindy"
	fmt.Println("not empty:", s)
	fmt.Println("slice length:", len(s))

	// append entity
	s = append(s, "dogge")
	s = append(s, "enemy", "f**k")
	fmt.Println("append:", s)

	// copy
	copys := make([]string, len(s))
	copy(copys, s)
	fmt.Println("copy:", copys)

	// slice the slices
	s1, s2, s3 := s[2:5], s[2:], s[:5]
	fmt.Println("s[2:5]:", s1)
	fmt.Println("s[2:]:", s2)
	fmt.Println("s[:5]:", s3)

	// directly create slice
	t := []string{"alice", "bob", "cindy"}
	fmt.Println("directly create slice:", t)

	// 2d
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		midlen := i+1
		twoD[i] = make([]int, midlen)
		for j := 0; j < midlen; j++ { // error: < i
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d:", twoD)
}