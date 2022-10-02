package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//struct (Anonymous struct tanpa pengisian field)
	var employee1 = struct {
		person   Person
		division string
	}{}
	employee1.person = Person{name: "Doyoung", age: 19}
	employee1.division = "Professional Singer"

	//struct (Anonymous struct dengan pengisian field)
	var employee2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Haechan", age: 17},
		division: "Lead Singer",
	}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee1: %+v\n", employee2)

}
