package main

import "fmt"

func main() {
	for {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)

		if x == 0 || y == 0 {
			break
		}
		if x > 0 && y > 0 {
			fmt.Println("primeiro")
		} else if x > 0 && y < 0 {
			fmt.Println("quarto")
		} else if x < 0 && y < 0 {
			fmt.Println("terceiro")
		} else {
			fmt.Println("segundo")
		}
	}
}
