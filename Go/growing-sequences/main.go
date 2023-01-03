package main

import (
	"fmt"
)

var x int = 1

func main() {
	for x != 0 {
		fmt.Scanf("%d", &x)
		for i := 1; x != 0 ; i++ {
			if i == x {
				fmt.Printf("%d\n", i)
				break
			}
			fmt.Printf("%d ", i)
		}
	}
}
