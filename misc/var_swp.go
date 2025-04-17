package main

import "fmt"

func main(){	
	var a, b int32 = 7, 2
	fmt.Printf("a = %d, b = %d\n", a, b)
	a = a^b
	b = a^b
	a = a^b
	fmt.Printf("a = %d, b = %d\n", a, b)
}
