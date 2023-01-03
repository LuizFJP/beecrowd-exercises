package main

import (
	"fmt"
	"math"
)

func main() {
	var numberOfLines float64

	fmt.Scanf("%f", &numberOfLines)

	for i := 1.0; i <= numberOfLines; i++ {
		fmt.Printf("%1.f %1.f %1.f\n", i, math.Pow(i, 2), math.Pow(i, 3))
		fmt.Printf("%1.f %1.f %1.f\n", i, math.Pow(i, 2) + 1, math.Pow(i, 3) + 1)
	}
}
