package main

import "fmt"

var (
	n1, n2, n3, n4, n5, n6 float64
	posCounter             int64
	numbers, posNumbers    []float64
	average                float64
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
	calculateAverage()

	fmt.Printf("%d valores positivos\n", posCounter)
	fmt.Printf("%.1f\n", average)

}

func positiverCounter() {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			posCounter++
			posNumbers = append(posNumbers, numbers[i])
		}
	}
}

func calculateAverage() {
	var sum float64
	for i := 0; i < len(posNumbers); i++ {
		sum += posNumbers[i]
	}
	average = sum / float64(len(posNumbers))
}
