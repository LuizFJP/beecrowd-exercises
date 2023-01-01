package main

import "fmt"

var A, B int64

func main(){
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	if ((B > A && B % A == 0) || (A > B && A % B == 0)) {
		fmt.Println("Sao Multiplos")
	} else {
		fmt.Println("Nao sao Multiplos")
	}
}