package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	division string
	person   Person
}

func main() {
	//struct (Embedded struct)
	var employee1 = Employee{}

	employee1.person.name = "Yoona"
	employee1.person.age = 25
	employee1.division = "Lead Dancer"

	fmt.Printf("%+v", employee1)
}
