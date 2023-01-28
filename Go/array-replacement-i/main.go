package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		var number int
		fmt.Scanf("%d", &number)
		if number <= 0 {
			fmt.Printf("X[%d] = 1\n", i)
		} else {
			fmt.Printf("X[%d] = %d\n", i, number)
		}
	}
}
