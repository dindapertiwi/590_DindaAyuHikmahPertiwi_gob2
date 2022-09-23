package main

import "fmt"

func main() {
	//slice (make function)
	var stuffs = make([]string, 3)

	_ = stuffs

	fmt.Printf("%#v", stuffs)

	//slice (append function)
	var dataStuffs = make([]string, 3)

	dataStuffs = append(dataStuffs, "bag", "book", "bottle")

	fmt.Printf("%#v", dataStuffs)

	//slice (append function with ellipsis)
	var stuffs1 = []string{"bag", "book", "bottle"}

	var stuffs2 = []string{"pen", "note", "earphone"}

	stuffs1 = append(stuffs1, stuffs2...)

	fmt.Printf("%#v", stuffs1)

	//slice (copy function)
	var stuffs3 = []string{"lotion", "powder", "moist"}

	var stuffs4 = []string{"bow", "cowlick", "comb"}

	nn := copy(stuffs3, stuffs4)

	fmt.Println("Stuffs3 =>", stuffs3)
	fmt.Println("Stuffs4 =>", stuffs4)
	fmt.Println("Copied elements =>", nn)

	//slice (slicing)
	var snacks1 = []string{"pie", "biscuit", "cake", "chips", "lollipop"}

	var snacks2 = snacks1[1:4]
	fmt.Printf("%#v\n", snacks2)

	var snacks3 = snacks1[0:]
	fmt.Printf("%#v\n", snacks3)

	var snacks4 = snacks1[:3]
	fmt.Printf("%#v\n", snacks4)

	var snacks5 = snacks1[:]
	fmt.Printf("%#v\n", snacks5)

	//slice (combining slicing and append)
	var food = []string{"soup", "sausage", "meatball", "chicken", "sushi"}

	food = append(food[:3], "noodles")

	fmt.Printf("%#v\n", food)

	//slice (backing array)
	var xanders = []string{"jayden", "lovely", "ethan", "sea", "rion"}

	var xanders1 = xanders[2:4]

	xanders1[0] = "rigel"

	fmt.Println("xanders => ", xanders)
	fmt.Println("xanders1 => ", xanders1)

	//slice (cap function)
	fruits := []string{"mango", "apple", "banana", "peach", "orange"}
	newFruits := []string{}

	newFruits = append(newFruits, fruits[0:2]...)

	fruits[0] = "guava"
	fmt.Println("fruits:", fruits)
	fmt.Println("newFruits:", newFruits)
}
