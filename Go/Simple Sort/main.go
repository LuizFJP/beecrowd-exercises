package main

import (
	"fmt"
)

var a, b, c int

// func main() {
// 	fmt.Scanf("%d", &a)
// 	fmt.Scanf("%d", &b)
// 	fmt.Scanf("%d", &c)
// 	numbers := []int{a,b,c}
// 	sort.Ints(numbers)
// 	for _, n := range numbers {
// 		fmt.Println(n)
// 	}
// 	fmt.Println("\n",a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

func main() {
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)

	var numbers = []int{a, b, c}

	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				aux := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = aux
			}
		}
	}
	for _,n := range numbers {
		fmt.Println(n)
	}
	fmt.Println()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
