package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	//struct (Pointer to a struct)
	var employee1 = Employee{name: "Jaehyun", age: 24, division: "Singer"}

	var employee2 *Employee = &employee1

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)

	fmt.Println(strings.Repeat("#", 20))

	employee2.name = "Lucas"

	fmt.Println("Employee1 name:", employee1.name)
	fmt.Println("Employee2 name:", employee2.name)

}
