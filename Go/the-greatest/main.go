package main

import (
	"fmt"
	"math"
)

var a int
var b int
var c int
func main() {
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	
	biggest := a
	for i := 0; i < 2; i++ {

		biggest = (biggest + b + int(math.Abs(float64(biggest - b)))) / 2
		b = c
	}

	fmt.Printf("%v eh o maior\n", biggest)
}