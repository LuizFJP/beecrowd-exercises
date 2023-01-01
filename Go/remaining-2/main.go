package main

import "fmt"

var number int

func main() {
	fmt.Scanf("%d", &number)

	for i := 1; i <= 10000; i++ {
		if i % number == 2 {
			fmt.Println(i)
		}
	}
}	
