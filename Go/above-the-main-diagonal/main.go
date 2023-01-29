package main

import "fmt"

var (
	numbers   [12][12]float64
	operation string
)

func main() {
	fmt.Scanf("%s", &operation)

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			var number float64
			fmt.Scanf("%f", &number)
			numbers[i][j] = number
		}
	}

	if operation == "S" {
		fmt.Printf("%.1f\n", sumArea())
	} else {
		fmt.Printf("%.1f\n", sumArea() / quantityArea(len(numbers) - 1))
	}
}

func sumArea() float64 {
	var sum float64
	for i := 0; i < len(numbers) - 1; i++ {
		for j := 1 + i; j < len(numbers[i]); j++ {
			sum += numbers[i][j]
		}
	}
	return sum
}

func quantityArea(size int) float64 {
	if size == 1 {
		return 1
	}

	return quantityArea(size - 1) + float64(size)
}
