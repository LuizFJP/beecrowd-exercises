package main

import "fmt"

var number int

func factorial(number int) int{
		if number == 1 {
			return 1
		} else {
			return number * factorial(number - 1)
		} 
}

func main() {
	fmt.Scanf("%d", &number)

	result := factorial(number)
	fmt.Println(result)
}
