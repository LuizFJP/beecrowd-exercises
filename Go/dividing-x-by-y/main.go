package main

import "fmt"

var numberOfEntries int

func main() {
	fmt.Scanf("%d", &numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		var x, y float64
		fmt.Scanf("%f %f", &x, &y)
		
		if y == 0 {
			fmt.Println("divisao impossivel")
		} else {
			fmt.Printf("%.1f\n", x/y)
		}

	}
}
