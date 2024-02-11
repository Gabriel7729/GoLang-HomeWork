package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	return file
}

func printSmallestDifferenceTeamWin() {

	file := readFile("data/football.dat")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var teamWithSmallestDiff string
	var smallestDiff int = 9999

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		if len(fields) < 8 || fields[0] == "Team" {
			continue
		}

		team := fields[1]
		goalsFor, _ := strconv.Atoi(fields[6])
		goalsAgainst, _ := strconv.Atoi(fields[8])

		goalDiff := abs(goalsFor - goalsAgainst)
		if goalDiff < smallestDiff {
			smallestDiff = goalDiff
			teamWithSmallestDiff = team
		}

	}

	fmt.Printf("Team with smallest difference is %s with a difference of %d\n", teamWithSmallestDiff, smallestDiff)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printSmallestTempSpread() {
	file := readFile("data/weather.dat")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var minSpread int = 9999
	var minSpreadDay int

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

	fmt.Printf("Day %d has the minimum temperature spread of %d\n", minSpreadDay, minSpread)
}

func main() {
	printSmallestDifferenceTeamWin()
	printSmallestTempSpread()
}
