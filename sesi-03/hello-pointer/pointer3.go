package main

import "fmt"

func main() {
	//Pointer (Pointer as a parameter)
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
