package main

import (
	"fmt"

)

var X int
var Y float64

func main() {
	fmt.Scanf("%d", &X)
	fmt.Scanf("%b", &Y)

	fmt.Printf("%.3f km/l\n", float64(X) / Y)
}