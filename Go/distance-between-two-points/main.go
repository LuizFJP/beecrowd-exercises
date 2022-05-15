package main

import (
	"fmt"
	"math"
)

var x1 float64
var x2 float64
var y1 float64
var y2 float64

func main() {
	fmt.Scanf("%b", &x1)
	fmt.Scanf("%b", &y1)
	fmt.Scanf("%b", &x2)
	fmt.Scanf("%b", &y2)
	distance := math.Sqrt(math.Pow((x2 - x1), 2) + math.Pow((y2 - y1), 2))

	fmt.Printf("%.4f\n", distance)
}