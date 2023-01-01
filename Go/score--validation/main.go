package main

import "fmt"

var gradesSum float64

func main() {
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
