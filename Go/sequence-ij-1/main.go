package main

import "fmt"

func main() {
	var i = 1
	for j := 60; j >= 0; j-=5 {
		fmt.Printf("I=%d J=%d\n", i, j)
		i += 3
	}
}
