package main

import (
	"fmt"
)

var A, B float64

func main() {
	fmt.Scanf("%b", &A)
	fmt.Scanf("%b", &B)
	expression := (A*3.5  + B*7.5) / 11
	result := fmt.Sprintf("MEDIA = %.5f", expression)
	fmt.Println(result)
}