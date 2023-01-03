package main

import "fmt"

var A, N, sumConsecutive int

func main() {
	fmt.Scanf("%d %d", &A, &N)

	for N <= 0 {
		fmt.Scanf("%d", &N)
	}
	
	for i := 0; i <= N-1; i++ {
		sumConsecutive += A + i
	}
	fmt.Println(sumConsecutive)
}
