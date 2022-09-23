// looping over string (rune-bye-rune)
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str3 := "minum"

	str4 := "air"

	fmt.Printf("str3 character length => %d\n", utf8.RuneCountInString(str3))
	fmt.Printf("str4 character length => %d\n", utf8.RuneCountInString(str4))
}
