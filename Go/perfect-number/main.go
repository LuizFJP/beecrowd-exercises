package main

import "fmt"

var tests int

func main() {
	fmt.Scanf("%d", &tests)	
	for i := 0; i < tests; i++ {
		var number, sum int
		fmt.Scanf("%d", &number)
		for i := 1; i < number; i++ {
			if number % i == 0 {
				sum += i
			}
		}

		if sum == number {
			fmt.Printf("%d eh perfeito\n", number)
		} else {
			fmt.Printf("%d nao eh perfeito\n", number)

		}
	}
}
