package main

import "fmt"

var X, Z, sum, counter int

func main() {
	fmt.Scanf("%d", &X)

	for Z <= X {
		fmt.Scanf("%d", &Z)
	}

	for i := X; ; i++ {
		if sum > Z {
			break
		}
		sum += i
		counter++
	}
	fmt.Println(counter)
}
