package main

import "fmt"

var number int

func main() {
	fmt.Scanf("%d", &number)
	
	for i := number; i < number + 12; i++ {
		if i % 2 == 1 {
			fmt.Println(i)
		}
	}
}
