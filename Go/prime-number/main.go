package main

import "fmt"

var tests int

func main() {
	fmt.Scanf("%d", &tests)
	for i := 0; i < tests; i++ {
		var prime int
		var isPrime bool
		fmt.Scanf("%d", &prime)
		for i := 2; i < prime; i++ {
			if prime % i == 0 {
				isPrime = true
				break
			}
		}
		if isPrime {
			fmt.Printf("%d nao eh primo\n", prime)
		} else {
			fmt.Printf("%d eh primo\n", prime)
		}
	}
}
