package main

import "fmt"

var numEntries, inCount, outCount int

func main() {
	fmt.Scanf("%d", &numEntries)
	for i := 0; i < numEntries; i++ {
		var number int
		fmt.Scanf("%d", &number)
		if number >= 10 && number <= 20 {
			inCount ++
		} else {
			outCount ++
		}
	}

	fmt.Printf("%d in\n", inCount)
	fmt.Printf("%d out\n", outCount)
}
