package main

import "fmt"

var grade1, grade2, grade3, grade4, gradeExam float64

const (
	WEIGHT_1    float64 = 2
	WEIGHT_2    float64 = 3
	WEIGHT_3    float64 = 4
	WEIGHT_4    float64 = 1
	SUM_WEIGHTS float64 = WEIGHT_1 + WEIGHT_2 + WEIGHT_3 + WEIGHT_4
)

func main() {
	scanEntries()
	average := calculateMedia()
	resultAverage(average)

}

func scanEntries() {
	fmt.Scanf("%f", &grade1)
	fmt.Scanf("%f", &grade2)
	fmt.Scanf("%f", &grade3)
	fmt.Scanf("%f", &grade4)
}

func calculateMedia() float64 {
	return ((grade1 * WEIGHT_1) +
		(grade2 * WEIGHT_2) +
		(grade3 * WEIGHT_3) +
		(grade4 * WEIGHT_4)) / SUM_WEIGHTS
}

func resultAverage(average float64) {
	fmt.Printf("Media: %.1f\n", average)
	if average >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if average < 5.0 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")
		fmt.Scanf("%f", &gradeExam)
		fmt.Printf("Nota do exame: %.1f\n", gradeExam)

		recalculatedGrade := (average + gradeExam) / 2
		if recalculatedGrade > 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado")
		}
		fmt.Printf("Media final: %.1f\n", recalculatedGrade)
	}
}
