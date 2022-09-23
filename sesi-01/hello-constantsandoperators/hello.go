package main

import "fmt"

func main() {
	//constant
	const full_name string = "Dinda Ayu Hikmah Pertiwi"

	fmt.Printf("Hello %s", full_name)

	//operators (arithmetic operators)
	var value = (2 + 2) * 3
	fmt.Println(value)

	//operators (relational operators)
	var firstCondition bool = 2 < 3
	var secondCondition bool = "carla" == "Carla"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("fourth condition:", fourthCondition)

	//operators (logical operators)
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}
