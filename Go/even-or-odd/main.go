package main

import (
	"fmt"
	"math"
)

var entriesNumber int

func main() {
	scanEntries()
	defineNumber()
}

func scanEntries() {
	fmt.Scanf("%d", &entriesNumber)
}

func defineNumber() {
	for i := 0; i < entriesNumber; i++ {
		var number int
		fmt.Scanf("%d", &number)

		if isZero(number) {
			fmt.Println("NULL")		
		} else if isPositive(number) {
			if isEven(number) {
				fmt.Println("EVEN POSITIVE")
			} else {
				fmt.Println("ODD POSITIVE")
			}
		} else {
			if isEven(number) {
				fmt.Println("EVEN NEGATIVE")
			} else {
				fmt.Println("ODD NEGATIVE")
			}
		}
	}
}

func isEven(number int) bool {
	return int(math.Abs(float64(number))) % 2 == 0
}

func isPositive(number int) bool {
	return number > 0
}

func isZero(number int) bool {
	return number == 0
}
