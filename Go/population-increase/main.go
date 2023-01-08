package main

import "fmt"

var (
	A, B   int
	PA, PB float64
)

func main() {
	var test int
	fmt.Scanf("%d", &test)
	for i := 0; i < test; i++ {
		var counterYears int
		fmt.Scanf("%d %d %f %f", &A, &B, &PA, &PB)

		for {
			if A > B {
				break
			}
			A = int(float64(A) * (PA / 100)) + A
			B =int(float64(B) * (PB / 100)) + B
			counterYears++
		}
		if counterYears > 100 {
			fmt.Println("Mais de 1 seculo.")
		} else {
			fmt.Printf("%d anos.\n", counterYears)
		}
	}
}
