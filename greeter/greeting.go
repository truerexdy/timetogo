package main

import "fmt"

func main() {
	var s string
	fmt.Println("Enter your name:")
	_, err := fmt.Scan(&s)
	if err != nil {
		fmt.Println("Invalid Input")
	} else {
		fmt.Println("Hello ", s)
	}
}
