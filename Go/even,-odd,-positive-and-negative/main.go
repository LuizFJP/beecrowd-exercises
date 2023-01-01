package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	oddCount, evenCount, posCount, negCount int
	numbers                                 []float64
)

func main() {
	scanEntries()
	counter()
	printResult()
}

func scanEntries() {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 5; i++ {
		scanner.Scan()
		number, _ := strconv.ParseFloat(scanner.Text(), 64)

		numbers = append(numbers, number)
	}
}

func counter() {
	for _, number := range numbers {
		if int(math.Abs(number))%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
		if number > 0 {
			posCount++
		} else if number == 0 {
			continue
		} else {
			negCount++
		}
	}
}

func printResult() {
	fmt.Printf("%d valor(es) par(es)\n", evenCount)
	fmt.Printf("%d valor(es) impar(es)\n", oddCount)
	fmt.Printf("%d valor(es) positivo(s)\n", posCount)
	fmt.Printf("%d valor(es) negativo(s)\n", negCount)
}
