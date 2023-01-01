package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	dayStart, dayEnd                                                   int
	hourStart, hourEnd, minuteStart, minuteEnd, secondStart, secondEnd int
	parsedDuration time.Duration
	days, hours, minutes, seconds int
)

const (
	DAY_IN_HOURS = 24
	HOURS_IN_MINUTES = 60
	MINUTES_IN_SECONDS = 60
)

func main() {
	scanInputs()
	calculateDuration()
	parseDuration()
	printResult()

}

func scanInputs() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for i := 0; i < 4; i++ {
		scanner.Scan()
		line := scanner.Text()

		lines = append(lines, line)
	}

	dayStart, _ = strconv.Atoi(lines[0][4:])
	dayEnd, _ = strconv.Atoi(lines[2][4:])
	hourStart, _ = strconv.Atoi(lines[1][:2])
	minuteStart, _ = strconv.Atoi(lines[1][5:7])
	secondStart, _ = strconv.Atoi(lines[1][10:])
	hourEnd, _ = strconv.Atoi(lines[3][:2])
	minuteEnd, _ = strconv.Atoi(lines[3][5:7])
	secondEnd, _ = strconv.Atoi(lines[3][10:])
}

func calculateDuration() {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), int(dayStart), int(hourStart), int(minuteStart), int(secondStart), 0, time.Local)
	endTime := time.Date(now.Year(), now.Month(), int(dayEnd), int(hourEnd), int(minuteEnd), int(secondEnd), 0, time.Local)
	duration := endTime.Sub(startTime)

	parsedDuration, _ = time.ParseDuration(duration.String())
}

func parseDuration() {
	days = int(parsedDuration.Hours()) / DAY_IN_HOURS
	hours = int(parsedDuration.Hours()) % DAY_IN_HOURS
	minutes = int(parsedDuration.Minutes()) % HOURS_IN_MINUTES
	seconds = int(parsedDuration.Seconds()) % MINUTES_IN_SECONDS
}

func printResult() {
	fmt.Printf("%d dia(s)\n", days)
	fmt.Printf("%d hora(s)\n", hours)
	fmt.Printf("%d minuto(s)\n", minutes)
	fmt.Printf("%d segundo(s)\n", seconds)
}
