package main

import "fmt"

func main() {
	str := "manca"

	for i, s := range str {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}
}
