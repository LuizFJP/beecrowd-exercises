package main

import "fmt"

var n int

func main() {
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		if n % i == 0 {
			fmt.Println(i)
		}
	}
}
