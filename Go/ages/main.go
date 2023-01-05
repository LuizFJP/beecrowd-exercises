package main

import "fmt"

var (i, sum, age float64 = 0.0, 0.0, 0.0)

func main() {
	for {
		fmt.Scanf("%f", &age)
		if age < 0.0 {
			break
		}
		sum += age
		i++
	}
	fmt.Printf("%.2f\n", sum/i)
}
