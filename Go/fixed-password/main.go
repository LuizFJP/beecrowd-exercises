package main

import "fmt"

var correctPassword = "2002"

func main() {
	for {
		var password string
		fmt.Scanf("%s", &password)
		if password == correctPassword {
			fmt.Println("Acesso Permitido")
			break
		}
		fmt.Println("Senha Invalida")
	}
}
