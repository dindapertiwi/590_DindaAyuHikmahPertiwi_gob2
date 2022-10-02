package main

import (
	"fmt"
	"math"
)

func main() {
	//Function (predefined return value)
	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)
}

func calculate(d float64) (area float64, circumference float64) {
	//Menghitung luas
	area = math.Pi * math.Pow(d/2, 2)
	//Menghitung keliling

	circumference = math.Pi * d

	return
}
