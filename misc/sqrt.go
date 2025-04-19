package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	if x<1{
		return -1
	}
	l, h := 1, int(x)
 	var m int

	for {
		x_int := int(x)
		m = l + ((h - l) / 2)
		m_sq := m * m
		m1_sq := (m + 1) * (m + 1)

		if m_sq <= x_int && m1_sq > x_int {
			break
		} else if m_sq > x_int {
			h = m
		} else if m1_sq < x_int {
			l = m
		}
	}

	z := float64(m)

	e := float64(1)
	for e > 0.000001 {
		e = (z*z - x) / (2 * z)
		z -= e
	}
	return z
}

func main() {
	var num float64
	fmt.Println("Enter a number:")
	fmt.Scanf("%f", &num)
	fmt.Println(Sqrt(num))
}
