package main

import (
	"fmt"
)

func gcd(a int64, b int64) int64{
	var ans int64
	switch{
	case a==0 && b==0:
		 ans=0
	case a==0:
		ans = b
	case b==0:
		ans = a
	}
	for b!=0{
		var temp int64 = a%b
		a = b
		b = temp
	}
	ans = a
	return ans
}

func main(){
	var a, b int64
	fmt.Println("Enter the first number:")
	fmt.Scanf("%d", &a)
	fmt.Println("Enter the first number:")
	fmt.Scanf("%d", &b)
	fmt.Printf("gcd(%d, %d) = %d\n", a, b, gcd(a,b))
}
