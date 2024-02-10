package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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
