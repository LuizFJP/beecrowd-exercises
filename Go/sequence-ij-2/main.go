package main

import "fmt"

func main() {
	var counter int = 1
	var j int = 7
	for i := 1; i <= 9; {
		if counter > 3 {
			i += 2
			j = 7
			counter = 1
		}
		if i == 11 {
			break
		}
		fmt.Printf("I=%d J=%d\n", i, j)
		j -= 1
		counter++

	}
}
