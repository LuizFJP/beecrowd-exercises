package main

import "fmt"

var number int

func main() {
	fmt.Scanf("%d", &number)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, number, i * number)
	}
}
