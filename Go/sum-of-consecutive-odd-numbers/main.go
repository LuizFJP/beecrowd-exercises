package main

import (
	"fmt"
	"math"
)

var input1, input2, end, sum int
var start int = 1

func main() {
	scanEntries()
	defineStartAndEnd()
	sumOddNumbers()
	printSum()
}

func scanEntries() {
	fmt.Scanf("%d", &input1)
	fmt.Scanf("%d", &input2)
}

func defineStartAndEnd() {
	if input1 > input2 {
		start += input2
		end = input1
	} else {
		start += input1
		end = input2
	}
}

func sumOddNumbers() {
	for ; start < end ; start++ {
		if int(math.Abs(float64(start))) % 2 == 1 {
			sum += start
		}
	} 
}

func printSum() {
	fmt.Println(sum)
}
