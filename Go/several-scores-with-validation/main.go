package main

import "fmt"

func main() {
	var keepCalculating int = 1
	for keepCalculating != 2 {
		var gradesSum float64
		if keepCalculating == 1 {
			for i := 0; i < 2; {
				var grade float64
				fmt.Scanf("%f", &grade)
				if grade < 0 || grade > 10 {
					fmt.Println("nota invalida")
					continue
				}
				gradesSum += grade
				i++
			}
			fmt.Printf("media = %.2f\n", gradesSum/2)
		}

		fmt.Println("novo calculo (1-sim 2-nao)")
		fmt.Scanf("%d", &keepCalculating)
	}
}
