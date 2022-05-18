package main

import (
	"fmt"
	"math"
)
var money float64
var notesExchange []float64
var coinsExchange []float64

func exchange(values []float64, arrayExchange *[]float64, text string) {
	for i := 0; i < len(values); i++ {
		*arrayExchange = append(*arrayExchange, math.Floor(money / values[i]))
		money -= values[i] * (*arrayExchange)[i]
		fmt.Printf("%v %v(s) de R$ %.2f\n", (*arrayExchange)[i], text, values[i])
	}
}

func main() {
	fmt.Scanf("%b", &money)

	notes := []float64 { 100.00, 50.00, 20.00, 10.00, 5.00, 2.00 }
	coins := []float64 { 1.00, 0.50, 0.25, 0.10, 0.05, 0.01 }

	fmt.Println("NOTAS:")
	exchange(notes, &notesExchange, "nota");
	fmt.Println("MOEDAS:")
	exchange(coins, &coinsExchange, "moeda");

}