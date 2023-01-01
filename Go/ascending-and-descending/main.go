package main

import "fmt"

func main() {
	for {
		var input1, input2  int
		fmt.Scanf("%d %d", &input1, &input2)
		if input1 == input2 {
			break
		}
		if input1 > input2 {
			fmt.Println("Decrescente")
		} else {
			fmt.Println("Crescente")
		}
	}
}
