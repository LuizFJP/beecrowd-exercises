package main

import ("fmt")

var h int
var kmh int

func main() {
	fmt.Scanf("%d", &h)
	fmt.Scanf("%d", &kmh)
	distance := float64(h * kmh)
	fmt.Printf("%.3f\n",distance / 12.0)
}