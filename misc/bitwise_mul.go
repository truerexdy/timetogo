package main

import (
	"fmt"
)

func main(){
	var a int32 = 8
	var b int32 = 2
	
	fmt.Printf("a = %d\nb = %d\n", a, b)

	var result int32 = 0
	for (b != 0){
		if(b & 1 != 0){
			result += a
		}
		a <<= 1
		b >>= 1
	}
	
	fmt.Printf("Result = %d\n",result)
}
