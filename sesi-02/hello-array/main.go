package main

import "fmt"

func main() {
	//Array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var strings = [3]string{"Colt", "Porco", "Zeke"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", strings)

	//Array (modify element through index)
	var warriors = [3]string{"pieck", "gabi", "falco"}
	warriors[0] = "pieck"
	warriors[1] = "gabi"
	warriors[2] = "falco"

	fmt.Printf("%#v\n", warriors)

	//Array (loop through elements)
	var soldiers = [3]string{"oluo", "petra", "mike"}
	for i, v := range soldiers {
		fmt.Printf("Index: %d, Value:%s\n", i, v)
	}

	//Array (multidimensional array)
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

}
