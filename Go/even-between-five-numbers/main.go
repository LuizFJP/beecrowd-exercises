package main

import "fmt"

var evenCounter, n1, n2, n3, n4, n5 int64

func main() {
	fmt.Scanf("%d", &n1)
	fmt.Scanf("%d", &n2)
	fmt.Scanf("%d", &n3)
	fmt.Scanf("%d", &n4)
	fmt.Scanf("%d", &n5)

	numbers := []int64{n1, n2, n3, n4, n5}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] % 2 == 0 {
			evenCounter ++
		}
	}
	fmt.Printf("%d valores pares\n", evenCounter)
}
