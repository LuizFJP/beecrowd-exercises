package main

import ("fmt")

var money int
func main() {
	fmt.Scanf("%d", &money)
	initialMoney := money
	var numbers []int
	notes := [7]int {100, 50, 20, 10, 5, 2, 1}

	for i := 0; i < len(notes); i++ {
		numbers = append(numbers, money / notes[i])
		money -= numbers[i] * notes[i]
	}
	fmt.Println(initialMoney)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d nota(s) de R$ %d,00\n", numbers[i], notes[i])
	}

}