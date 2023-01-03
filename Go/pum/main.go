package main

import "fmt"

func main() {
	var (
		firstColumn   int = 1
		secondColumn  int = 2
		thirdColumn   int = 3
		numberOfLines int
	)

	fmt.Scanf("%d", &numberOfLines)

	for i := 0; i < numberOfLines; i++ {
		fmt.Printf("%d %d %d PUM\n", firstColumn, secondColumn, thirdColumn)
		firstColumn += 4
		secondColumn += 4
		thirdColumn += 4
	}
}
