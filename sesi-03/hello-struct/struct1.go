package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	//struct (Initializing struct)
	var employee1 = Employee{}
	employee1.name = "Hange Zoe"
	employee1.age = 27
	employee1.division = "Head Commander Survey Corps"

	var employee2 = Employee{name: "Erwin Smith", age: 29, division: "Commander Survey Corps"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}
