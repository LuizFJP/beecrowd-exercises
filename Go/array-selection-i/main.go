package main

import "fmt"

var number float64 = 100 

func main() {
	for i := 0; i < 100; i++ {
		fmt.Scanf("%f", &number)
		if number <= 10 {
			fmt.Printf("A[%d] = %.1f\n", i, number)
		}
	}
}
