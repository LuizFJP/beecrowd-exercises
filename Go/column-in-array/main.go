package main

import "fmt"

var (
	numbers   [12][12]float64
	column      int
	operation string
)

func main() {
	fmt.Scanf("%d", &column)
	fmt.Scanf("%s", &operation)

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			var number float64
			fmt.Scanf("%f", &number)
			numbers[i][j] = number
		}
	}

	if operation == "S" {
		fmt.Printf("%.1f\n", sumLine())
	} else {
		fmt.Printf("%.1f\n", sumLine() / float64(len(numbers)))
	}

}

func sumLine() float64 {
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i][column]
	}
	return sum
}
