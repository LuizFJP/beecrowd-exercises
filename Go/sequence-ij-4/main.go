package main

import (
	"fmt"
)

var (
	count    int     = 1
	j        float64 = 1
	i        float64 = 0
	decimals float64 = 0.2
)

func main() {
	for {
		if count > 3 {
			count = 1
			j = 1
			j += decimals
			i += 0.2
			decimals += 0.2

		}
		if i > 2 {
			break
		}

		formatedI := fmt.Sprintf("%.1f", i)

		if i == 0.0 || formatedI == "2.0" || formatedI == "1.0" {
			fmt.Printf("I=%d J=%d\n", int(i + 0.5), int(j))
		} else {
			fmt.Printf("I=%.1f J=%.1f\n", i, j)
		}

		j++
		count++
	}
}
