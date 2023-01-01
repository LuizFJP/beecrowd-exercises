package main

import "fmt"

var greatest, smallest int

func main() {
	for {
		var input1, input2, sumNumbers int
		fmt.Scanf("%d %d", &input1, &input2)
		if input1 <= 0 || input2 <= 0 {
			break
		}
		if input1 > input2 {
			greatest = input1
			smallest = input2
		} else {
			greatest = input2
			smallest = input1
		}
	
		for i := smallest; i <= greatest; i++ {
			sumNumbers += i
			fmt.Printf("%d ", i)
		}
	
		fmt.Printf("Sum=%d\n", sumNumbers)
	}
}
