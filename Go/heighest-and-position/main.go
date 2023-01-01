package main

import "fmt"

var (
	highest      = 0
	position int = 0
)

func main() {
	for i := 1; i <= 100; i++ {
		var number int
		fmt.Scanf("%d", &number)

		if number > highest {
			highest = number
			position = i
		}
	}

	fmt.Println(highest)
	fmt.Println(position)
}
