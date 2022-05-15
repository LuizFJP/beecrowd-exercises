package main

import (
	"fmt"
	"math"
)

var A, B, C float64

func triangle(x, y float64) string {
	result := fmt.Sprintf("TRIANGULO: %.3f", (x * y) / 2)
	return result
}

func circle(r float64) string {
	pi := 3.14159
	result := fmt.Sprintf("CIRCULO: %.3f", pi * math.Pow(C, 2) )
	return result
}

func trapezium(B, b, h float64) string {
	result := fmt.Sprintf("TRAPEZIO: %.3f", (B + b)* h / 2)
	return result
}

func square(s float64) string {
	result := fmt.Sprintf("QUADRADO: %.3f", math.Pow(s, 2))
	return result
}

func rectangle(x, y float64) string {
	result := fmt.Sprintf("RETANGULO: %.3f", x * y)
	return result
}

func main() {
	fmt.Scanf("%b", &A)
	fmt.Scanf("%b", &B)
	fmt.Scanf("%b", &C)

	fmt.Println(triangle(A, C))
	fmt.Println(circle(C))
	fmt.Println(trapezium(A, B, C))
	fmt.Println(square(B))
	fmt.Println(rectangle(A, B))
}