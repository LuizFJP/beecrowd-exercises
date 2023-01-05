package main

import "fmt"

var number int

func main() {
	fmt.Scanf("%d", &number)
	var (
		A int = 0
		B int = 1
		C int
		i int
	)
	for {
		fmt.Printf("%d", A)
		i++
		if i == number {
			break
		}
		fmt.Print(" ")

		C = A + B
		A = B
		B = C
	}
	fmt.Println()
}
