package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, y int

	fmt.Scanf("%d %d", &x, &y)

	var initializeForJ, limitForJ, currentNumber int

	for i := 1; ; i++ {
		var output string
		for j := initializeForJ; j < x + limitForJ && currentNumber != y; j++ {
			currentNumber = i + j
			output += fmt.Sprintf("%d ", currentNumber)
			initializeForJ = j
		}
		limitForJ = initializeForJ
		fmt.Println(strings.TrimSpace(output))
		if(currentNumber == y) {
			break
		}
	}
}
