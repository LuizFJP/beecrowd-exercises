package main

import "fmt"

var (
	n1, n2, n3, n4, n5, n6 float64
	posCounter             int64
	numbers    []float64
)

func main() {
	fmt.Scanf("%f", &n1)
	fmt.Scanf("%f", &n2)
	fmt.Scanf("%f", &n3)
	fmt.Scanf("%f", &n4)
	fmt.Scanf("%f", &n5)
	fmt.Scanf("%f", &n6)

	numbers = []float64{n1, n2, n3, n4, n5, n6}

	positiverCounter()

	fmt.Printf("%d valores positivos\n", posCounter)
}

func positiverCounter() {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			posCounter++
		}
	}
}
