package main

import (
	"fmt"
)

var numTests, x, y int

func main() {
	fmt.Scanf("%d", &numTests)
	for i := 0; i < numTests; i++ {
		fmt.Scanf("%d %d", &x, &y)

		if x % 2 == 0 {
			x += 1
		}
		var counter, sum = 0, 0

		for i := x; counter < y; i+=2 {
			sum += i
			counter++
		}
	fmt.Println(sum)

	}
}
