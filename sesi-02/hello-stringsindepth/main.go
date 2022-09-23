package main

import "fmt"

func main() {
	//looping over string (byte-by-byte)
	city := "Yogyakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])

		var city []byte = []byte{89, 111, 103, 121, 97, 107, 97, 114, 116, 97}

		fmt.Println(string(city))
	}

	//looping over string (rune-by-rune)
	str1 := "makan"

	str2 := "mainca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))
}
