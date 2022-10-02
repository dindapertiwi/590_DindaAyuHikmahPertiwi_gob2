package main

import (
	"fmt"
	"strings"
)

func main() {
	//Pointer (Changing value through pointer)
	var firstPerson string = "Kayleen"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value)  :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value)  :", *secondPerson)
	fmt.Println("secondNumber (memori address) :", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Leo"

	fmt.Println("firstPerson (value)  :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value)  :", secondPerson)
	fmt.Println("secondNumber (memori address) :", secondPerson)

}
