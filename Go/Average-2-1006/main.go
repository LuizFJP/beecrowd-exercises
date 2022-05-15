package main

import (
	"fmt"
)

var A, B, C float64

func main() {
	_, _ = fmt.Scanf("%b", &A)
	_, _ = fmt.Scanf("%b", &B)
	_, _ = fmt.Scanf("%b", &C)

	expression := ((A * 2 + B * 3 + C * 5) / 10)
	result := fmt.Sprintf("MEDIA = %.1f", expression)
	fmt.Println(result)
}