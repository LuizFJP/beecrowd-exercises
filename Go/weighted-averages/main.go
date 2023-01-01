package main

import "fmt"

var numberTests int

const (
	WEIGHT_1 = 2
	WEIGHT_2 = 3
	WEIGHT_3 = 5

	SUM_OF_ALL_WEIDHTS = WEIGHT_1 + WEIGHT_2 + WEIGHT_3
)

func main() {
	fmt.Scanf("%d", &numberTests)

	for i := 0; i < numberTests; i++ {
		var number1, number2, number3 float64

		fmt.Scanf("%f", &number1)
		fmt.Scanf("%f", &number2)
		fmt.Scanf("%f", &number3)

		fmt.Printf("%.1f\n",
			((number1 * WEIGHT_1) +
				(number2 * WEIGHT_2) +
				(number3 * WEIGHT_3)) /
				SUM_OF_ALL_WEIDHTS)
	}
}
