package main

import "fmt"

var salary, tax, value float64

func main() {
	fmt.Scanf("%f", &salary)
	if salary <= 2000.00 {
		fmt.Println("Isento")
	} else {
		for {
			if salary > 4500.00 {
				value = salary - 4500.00
				tax += value * 0.28 
				salary -= value
			} else if salary > 3000.00 && salary <= 4500.00 {
				value = salary - 3000.00
				tax += value * 0.18
				salary -= value
			} else if salary > 2000.00 && salary <= 3000.00 {
				value = salary - 2000.00
				tax += value * 0.08
				salary -= value
			} else {
				break
			}
		}
		fmt.Printf("R$ %.2f", tax)
	}
}
