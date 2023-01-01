package main

import (
	"fmt"
	"math"
)

var number int

func main() {
	fmt.Scanf("%d", &number)

	for i := 1; i <= number; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d^2 = %1.f\n", i, math.Pow(float64(i), 2))
		}
	}
}
