// learngo by examples: interfaces
package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, length float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width*r.length
}

func (r rect) perim() float64 {
	return 2*r.width+2*r.length
}

func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c circle) perim() float64 {
	return 2*math.Pi*c.radius
}

// If a variable has an interface type, then we can call methods 
// that are in the named interface. Here’s a generic measure 
// function taking advantage of this to work on any geometry.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width:3, length:4}
	c := circle{radius:5}

	measure(r)
	measure(c)
}
