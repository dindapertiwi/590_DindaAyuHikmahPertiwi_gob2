package main

import "fmt"

func main() {
	//function 2
	greet("Chasen", "Jalan Taman Kota")
}

func greet(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}
