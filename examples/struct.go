// learngo by examples: struct
package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println("Welcome to Barcelona, Dimaria!")
	fmt.Println(person{"Dimaria", 27})
	fmt.Println(person{name : "Dimaria", age : 27})
	// error: fmt.Println(person{"Dimaria"})
	fmt.Println(person{name : "Dimaria"})
	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{"Cristiano Ronaldo", 31})

	people := person{"Wasdns", 20} // people := person{name : "Wasdns", age : 20}
	fmt.Println(people.name, people.age)

	peoplePointer := &people
	fmt.Println(peoplePointer.name, peoplePointer.age)
}