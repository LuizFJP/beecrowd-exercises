package main

import (
	"fmt"
	"math"
)

var A, B, C float64

func main() {
	fmt.Scanf("%f", &A)
	fmt.Scanf("%f", &B)
	fmt.Scanf("%f", &C)
	var numbers = []float64{A, B, C}

	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				aux := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = aux
			}
		}
	}

	if numbers[0] >= numbers[1] + numbers[2] {
		fmt.Println("NAO FORMA TRIANGULO")
	} else {
		if math.Pow(numbers[0], 2) == math.Pow(numbers[1], 2) + math.Pow(numbers[2], 2) {
			fmt.Println("TRIANGULO RETANGULO")
		} 
		if math.Pow(numbers[0], 2) > math.Pow(numbers[1], 2) + math.Pow(numbers[2], 2) {
			fmt.Println("TRIANGULO OBTUSANGULO")
		}
		if math.Pow(numbers[0], 2) < math.Pow(numbers[1], 2) + math.Pow(numbers[2], 2){
			fmt.Println("TRIANGULO ACUTANGULO")
		}
		if numbers[0] == numbers[1] && numbers[1] == numbers[2] {
			fmt.Println("TRIANGULO EQUILATERO")
		}
		if numbers[0] == numbers[1] && numbers[1] != numbers[2] || numbers[0] == numbers[2] && numbers[0] != numbers[1] || numbers[1] == numbers[2] && numbers[2] != numbers[0] {
			fmt.Println("TRIANGULO ISOSCELES")
		}
 	}
}