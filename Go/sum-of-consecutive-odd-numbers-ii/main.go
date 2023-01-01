package main

import "fmt"

var numEntries, greatest, smallest int

func main() {
	fmt.Scanf("%d", &numEntries)
	for i := 0; i < numEntries; i++ {
		var input1, input2, sumOdds int
		fmt.Scanf("%d %d", &input1, &input2)

		if input1 > input2 {
			greatest = input1
			smallest = input2
		} else {
			greatest = input2
			smallest = input1
		}

		for i := smallest + 1; i < greatest; i++ {
			if i % 2 == 1 {
				sumOdds += i
			}
		}
		fmt.Println(sumOdds)
	}
}
