package main

import (
	"fmt"
	"math"
)

var x int

func main() {
	for {
		var sum int
		fmt.Scanf("%d", &x)
		if x == 0 {
			break
		}

		if int(math.Abs(float64(x))) % 2 == 1 {
			x += 1
		}

		for i := x; i <= x + 8; i+=2 {
			sum +=  i
		}
		fmt.Println(sum)
	}
}
