package main

import "fmt"

var A, B int64

func main() {
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)

	if A >= B {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", (24 - A) + B)
	} else {
		fmt.Printf("O JOGO DUROU %d HORA(S)\n", B - A)
	}
}