package main

import "fmt"

var x, y int

func main() {
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)

	if x > y {
		aux := y
		y = x
		x = aux
	}

	for i := x + 1; i < y; i++ {
		if i % 5 == 2 || i % 5 == 3  {
			fmt.Println(i)
		}
	}
}
