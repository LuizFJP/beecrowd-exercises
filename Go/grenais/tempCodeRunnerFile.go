package main

import "fmt"

var counterGrenais, interWins, gremioWins, ties int

func main() {
	var keep int = 1
	for keep != 2 {
		var interGoals, gremioGoals int
		fmt.Scanf("%d %d", &interGoals, &gremioGoals)

		counterGrenais++

		if interGoals > gremioGoals {
			interWins++			
		} else if interGoals == gremioGoals {
			ties++
		} else {
			gremioWins++
		}
		fmt.Println("Novo grenal (1-sim 2-nao)")
		fmt.Scanf("%d", &keep)
	}

	fmt.Printf("%d grenais\n", counterGrenais)
	fmt.Printf("Inter:%d\n", interWins)
	fmt.Printf("Gremio:%d\n", gremioWins)
	fmt.Printf("Empates:%d\n", ties)
}
