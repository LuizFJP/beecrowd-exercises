package main

import "fmt"

var number [20]int

func main() {
	for i := 0; i < 20; i++ {
		fmt.Scanf("%d", &number[i])
	}

	for i := 0; i < 20; i++ {
			fmt.Printf("N[%d] = %d\n", i, number[len(number) - i - 1])
	}
}
