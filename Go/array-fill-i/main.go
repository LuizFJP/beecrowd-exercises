package main

import "fmt"

var number int

func main() {
	fmt.Scanf("%d", &number)
	for i := 0; i < 10; i++ {
		fmt.Printf("N[%d] = %d\n", i, number)
		number *= 2
	}
}
