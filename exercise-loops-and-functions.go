package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	tmp := float64(1)
	for i := 0; i < 10; i++ {
		fmt.Printf("//times:%2d sqrt:%.30f\n", i+1, z)
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-tmp) < 1e-10 {
			break
		}
		tmp = z
	}
	return z
}

func main() {
	fmt.Println("Sqrt(2)", Sqrt(2))
	fmt.Println("math.Sqrt(2)", math.Sqrt(2))
}
