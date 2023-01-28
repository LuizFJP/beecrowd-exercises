package main

import "fmt"

var numTestes, serie int

func main() {
	fmt.Scanf("%d", &numTestes)

	for i := 0; i < 1000; i++ {
		fmt.Printf("N[%d] = %d\n", i, serie)
		serie++
		if serie == numTestes {
			serie = 0
		}
	}
}
