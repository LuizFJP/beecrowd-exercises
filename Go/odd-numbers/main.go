package main

import "fmt"

var number int

func main() {
	fmt.Scanf("%d", &number)

	for i := 1; i <= number; i++ {
		if i % 2 == 1 {
			fmt.Println(i)
		}
	}
}
