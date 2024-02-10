package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data/football.dat")
	if err != nil {
		fmt.Println("There was an error opening the file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var teamWithSmallestDifference string
	var smallestDifference int = 9999

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		if len(fields) < 10 || fields[0] == "Team" {
			continue
		}

		team := fields[1]
		goalsFor, _ := strconv.Atoi(fields[6])
		goalsAgainst, _ := strconv.Atoi(fields[8])

		goalDifference := abs(goalsFor - goalsAgainst)

		if goalDifference < smallestDifference {
			smallestDifference = goalDifference
			teamWithSmallestDifference = team
		}
	}

	fmt.Println("Team with smallest difference: ", teamWithSmallestDifference)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main2() {
	file, err := os.Open("data/weather.dat")
	if err != nil {
		fmt.Println("There was an error opening the file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var minSpread int = 9999
	var minSpreadDay int = 0

	for scanner.Scan() {

		line := scanner.Text()

		fields := strings.Fields(line)

		if len(fields) < 3 || fields[0] == "Dy" {
			continue
		}

		day, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}

		maxTemp, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}

		minTemp, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}

		spread := maxTemp - minTemp
		if spread < minSpread {
			minSpread = spread
			minSpreadDay = day
		}
	}

	fmt.Println("Day with minimum temperature spread: ", minSpreadDay)
}
