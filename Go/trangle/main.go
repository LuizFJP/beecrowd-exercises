package main

import "fmt"

var (
	A, B, C float64
)

func main(){
	fmt.Scanf("%f",&A)
	fmt.Scanf("%f",&B)
	fmt.Scanf("%f",&C)

	if 
	((A + B > C) && 
	(A + C > B) && 
	(B + C > A)) {
		fmt.Printf("Perimetro = %.1f\n", A + B + C)
		} else {
			fmt.Printf("Area = %.1f\n", ((A + B) * C) / 2)
		}
}