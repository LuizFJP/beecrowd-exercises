package main

import "fmt"

var A, B, C, D int64
var totalMinutesDuration, hours, minutes int64
const (
	HOURS_IN_MINUTES = 60
)

func main() {

	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)
	fmt.Scanf("%d", &D)

	if A >= C {
		if A == C && B < D {
			totalMinutesDuration = D - B
		} else if B != 0 {
			totalMinutesDuration = ((23-A)*HOURS_IN_MINUTES + (HOURS_IN_MINUTES - B)) + ((C * HOURS_IN_MINUTES) + D)
		} else {
			totalMinutesDuration = (24-A)*HOURS_IN_MINUTES + ((C * HOURS_IN_MINUTES) + D)
		}
	} else {
		totalMinutesDuration = ((C * 60) + D) - ((A * 60) + B) 
	}
	hours = totalMinutesDuration / 60
	minutes = totalMinutesDuration % 60
	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", hours, minutes)

}
