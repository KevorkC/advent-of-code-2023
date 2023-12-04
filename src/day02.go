package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeSet struct {
	red   int
	green int
	blue  int
}

// Config: 12 red cubes, 13 green cubes, and 14 blue cubes
// Example:
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

func possibilityOracle(line string) bool {
	var cubeMax CubeSet = CubeSet{red: 12, green: 13, blue: 14}
	handfuls := strings.Split(line, ";")
	// fmt.Println("Handfuls:", handfuls)
	for _, handful := range handfuls {
		colors := strings.Split(handful, ",")
		for _, color := range colors {
			var amountAndColor = strings.Split(color, " ")
			// Converting to int
			var amount, err = strconv.Atoi(amountAndColor[1])
			if err != nil {
				fmt.Println("Error - Failed to convert string 'amount' to intreger:", err)
			}
			var color string = amountAndColor[2]

			if strings.Contains(color, "red") && amount > cubeMax.red {
				return false
			}
			if strings.Contains(color, "green") && amount > cubeMax.green {

				return false
			}
			if strings.Contains(color, "blue") && amount > cubeMax.blue {
				return false
			}
		}
	}
	return true
}

func main() {
	// Testing
	// string := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	// fmt.Println(possibilityOracle(string))

	// Task
	var filePath string = "files/day02"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// cubeMax := CubeSet{red: 12, green: 13, blue: 14}
	var possibleGamesIdSum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var gameIdAndRest = strings.Split(line, ":")
		var gameId, err = strconv.Atoi(strings.Split(gameIdAndRest[0], " ")[1])
		if err != nil {
			fmt.Println("Error - Failed to convert string 'gameId' to intreger:", err)
		}
		var restOfLine string = gameIdAndRest[1]
		if possibilityOracle(restOfLine) {
			possibleGamesIdSum += (gameId)
		}
	}
	fmt.Println(possibleGamesIdSum)
}
