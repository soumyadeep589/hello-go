package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:=1.0
	for i:=1; i <=10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("real : ", math.Sqrt(2))
}