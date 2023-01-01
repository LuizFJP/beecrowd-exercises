package main

import "fmt"

var ddd = map[int]string {
	61: "Brasilia",
	71: "Salvador",
	11: "Sao Paulo",
	21: "Rio de Janeiro",
	32: "Juiz de Fora",
	19: "Campinas",
	27: "Vitoria",
	31: "Belo Horizonte",
}

var number int

func main() {
	fmt.Scanf("%d", &number)

	city, found := ddd[number]
	if !found {
		fmt.Println("DDD nao cadastrado")
	} else {
		fmt.Println(city)
	}
}
