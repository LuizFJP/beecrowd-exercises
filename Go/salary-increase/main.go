package main

import "fmt"

var salary float64
var percentage float64

func main() {
	fmt.Scanf("%f", &salary)

	if (salary <= 400.0) {
		percentage = 0.15
	} else if (salary > 400.00 && salary <= 800.00) {
		percentage = 0.12
	} else if (salary > 800.00 && salary <= 1200.00) {
		percentage = 0.1
	} else if (salary > 1200.00 && salary <= 2000.00) {
		percentage = 0.07
	} else {
		percentage = 0.04
	}

	readjustment := salary * percentage
	newSalary := salary + readjustment

	fmt.Printf("Novo salario: %.2f\n", newSalary)
	fmt.Printf("Reajuste ganho: %.2f\n", readjustment)
	fmt.Printf("Em percentual: %d %%\n", int(percentage * 100))
}
