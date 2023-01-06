package main

import (
	"fmt"
	"math"
)

var sum float64

func main() {
	for i := 0.0; i < 19; i++ {
		sum += ((2 * i) + 1) / math.Pow(2, i) 
	}
	fmt.Printf("%.2f\n", sum)
}
