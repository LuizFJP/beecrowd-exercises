package main

import "fmt"

var sum float64

func main() {
	for i := 1.0; i <= 100; i++ {
		sum += 1 / i
	}
	fmt.Printf("%.2f\n", sum)
}
