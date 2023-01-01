package main

import "fmt"

var (
	count int = 1
	j     int = 7
	i int = 1
)

func main() {
	for {
		if count > 3 {
			count = 1
			j += 5
			i += 2
		}
		if i == 11 {
			break
		}

		fmt.Printf("I=%d J=%d\n", i, j)
		j--
		count++
	}
}
