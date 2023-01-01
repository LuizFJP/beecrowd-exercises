package main

import "fmt"

var subphylum, animal, group string

func main() {
	fmt.Scanf("%s", &subphylum)
	fmt.Scanf("%s", &animal)
	fmt.Scanf("%s", &group)

	if subphylum == "vertebrado" {
		if animal == "ave" {
			if group == "carnivoro" {
				fmt.Println("aguia")
			} else {
				fmt.Println("pomba")
			}
		} else {
			if group == "onivoro" {
				fmt.Println("homem")
			} else {
				fmt.Println("vaca")
			}
		}
	} else {
		if animal == "inseto" {
			if group == "hematofago" {
				fmt.Println("pulga")
			} else {
				fmt.Println("lagarta")
			}
		} else {
			if group == "hematofago" {
				fmt.Println("sanguessuga")
			} else {
				fmt.Println("minhoca")
			}
		}
	}
}
