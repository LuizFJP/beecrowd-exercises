package main

import "fmt"

var (
	smallest = 1000000000
 	position int
)

func main() {
	var tests int
	fmt.Scanf("%d", &tests)

	for i := 0; i < tests; i++ {
		var number int
		fmt.Scanf("%d", &number)
		if number < smallest {
			smallest = number
			position = i
		}
	}

	fmt.Printf("Menor valor: %d\n", smallest)
	fmt.Printf("Posicao: %d\n", position)
}
