package main

import "fmt"

var x, y, sumNotMultiplesOf13 int

func main() {
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	if x > y {
		aux := y
		y = x
		x = aux
	}

	for i := x; i <= y ; i++ {
		if i % 13 != 0 {
			sumNotMultiplesOf13 += i
		}
	}

	fmt.Println(sumNotMultiplesOf13)
}
