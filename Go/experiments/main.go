package main

import (
	"fmt"
)

var (
	numOfEntries                                      int
	animalsTotal, numOfRabbits, numOfRats, numOfFrogs float64
)

func main() {
	fmt.Scanf("%d", &numOfEntries)

	for i := 0; i < numOfEntries; i++ {
		var numOfTypes int
		var typeAnimal string
		fmt.Scanf("%d %s", &numOfTypes, &typeAnimal)

		switch typeAnimal {
		case "C":
			numOfRabbits += float64(numOfTypes)
		case "R":
			numOfRats += float64(numOfTypes)
		case "S":
			numOfFrogs += float64(numOfTypes)
		}

		animalsTotal += float64(numOfTypes)
	}

	fmt.Printf("Total: %1.f cobaias\n", animalsTotal)
	fmt.Printf("Total de coelhos: %1.f\n", numOfRabbits)
	fmt.Printf("Total de ratos: %1.f\n", numOfRats)
	fmt.Printf("Total de sapos: %1.f\n", numOfFrogs)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", float64(numOfRabbits/animalsTotal)*100.00)
	fmt.Printf("Percentual de ratos: %.2f %%\n", float64(numOfRats/animalsTotal)*100.00)
	fmt.Printf("Percentual de sapos: %.2f %%\n", float64(numOfFrogs/animalsTotal)*100.00)
}
