package main

import "fmt"

func main() {
	//number (integer)
	var first = 89
	var second = -12

	fmt.Printf("tipe data first : %T\n", first)
	fmt.Printf("bilangan second: %T\n", second)

	//number (integer)
	var dataFirst uint8 = 89
	var dataSecond int8 = -12

	fmt.Printf("tipe data first : %T\n", dataFirst)
	fmt.Printf("tipe data second: %T\n", dataSecond)

	//number (float)
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal Number: %f\n", decimalNumber)
	fmt.Printf("decimal Number: %.3f\n", decimalNumber)

	//bool
	var condition bool = true
	fmt.Printf("is it permitted? %t \n", condition)

	//string
	var message = `Selamat datang di "Hacktiv8".
Salam kenal :).
Mari belajar "Scalable Web Service With Go".`

	fmt.Println(message)
}
