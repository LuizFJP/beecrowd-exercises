package main

import "fmt"

var tests int
var fibonnaci []int
var numbers []int

func main() {
	fmt.Scanf("%d", &tests)
	var (
		A int = 0
		B int = 1
		C int
	)

	fibonnaci = append(fibonnaci, A)
	fibonnaci = append(fibonnaci, B)

	for i := 0; i < 60; i++ {
		C = A + B
		fibonnaci = append(fibonnaci, C)
		A = B
		B = C
	}

	for i := 0; i < tests; i++ {
		var number int
		fmt.Scanf("%d", &number)
		numbers = append(numbers, number)
	}

	for i := 0; i < tests; i++ {
		fmt.Printf("Fib(%d) = %d\n", numbers[i], fibonnaci[numbers[i]])
	}
}
