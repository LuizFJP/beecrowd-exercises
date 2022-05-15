package main

import (
	"fmt"
	"math"
)

var R float64
var pi float64 = 3.14159

func main() {
	fmt.Scanf("%b", &R)
	result := fmt.Sprintf("VOLUME = %.3f", (4/3.0) * pi * math.Pow(R, 3))
	fmt.Println(result)
}