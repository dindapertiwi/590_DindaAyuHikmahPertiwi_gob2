package main

import "fmt"

func main() {
	//function 1
	greet("Chasey", 22)
}

func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}
