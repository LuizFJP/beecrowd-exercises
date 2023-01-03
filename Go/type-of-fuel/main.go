package main

import "fmt"

var fuels = map[int]int {
	1: 0,
	2: 0,
	3: 0,
}

var inputData int

func main() {
	for inputData != 4 {
		fmt.Scanf("%d", &inputData)
		fuels[inputData] += 1
	}

	fmt.Println("MUITO OBRIGADO")
	fmt.Printf("Alcool: %d\n", fuels[1])
	fmt.Printf("Gasolina: %d\n", fuels[2])
	fmt.Printf("Diesel: %d\n", fuels[3])
}
