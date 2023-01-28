package main

import "fmt"

var even, odd []int

func main() {
	for i := 0; i < 15; i++ {
		var number int
		fmt.Scanf("%d", &number)
		placeNumber(number)
		array := checkArraysAreFull()
		if array != nil {
			cleanArray(array)
		}
	}
	printArray(odd, "impar")
	printArray(even, "par")
}

func placeNumber(number int) {
	if number%2 == 0 {
		even = append(even, number)
	} else {
		odd = append(odd, number)
	}
}

func checkArraysAreFull() *[]int {
	if len(even) == 5 {
		printArray(even, "par")
		return &even
	} else if len(odd) == 5 {
		printArray(odd, "impar")
		return &odd
	} else {
		return nil
	}
}

func printArray(array []int, text string) {
	for i := 0; i < len(array); i++ {
		fmt.Printf("%s[%d] = %d\n", text, i, array[i])
	}
}

func cleanArray(array *[]int) {
	*array = nil
}
