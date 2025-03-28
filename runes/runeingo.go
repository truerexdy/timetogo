package main

import (
	"fmt"
)

func main() {
	var s string = "Hello World"
	for ind, char := range s {
		fmt.Printf("Ind: %d, Char: %c and Code: %U\n", ind, char, char)
	}
}
