package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//struct (slice of struct)
	var people = []Person{
		{name: "Dokyeom", age: 21},
		{name: "Seungkwan", age: 21},
		{name: "Hoshi", age: 21},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
