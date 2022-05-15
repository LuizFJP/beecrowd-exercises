package main

import (
	"fmt"
)

var p1, unit1, p2, unit2 int
var price1, price2 float64

func main() {
	_, _ = fmt.Scanf("%d", &p1)
	_, _ = fmt.Scanf("%d", &unit1)
	_, _ = fmt.Scanf("%b", &price1)

	_, _ = fmt.Scanf("%d", &p2)
	_, _ = fmt.Scanf("%d", &unit2)
	_, _ = fmt.Scanf("%b", &price2)

	fmt.Printf("VALOR A  PAGAR: R$ %.2f\n", (float64(unit1)*price1) + (float64(unit2)*price2))
}