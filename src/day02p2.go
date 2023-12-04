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

// Example:
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green

func gamePowerCalculator(line string) int {
	var cubeMin CubeSet = CubeSet{red: 0, green: 0, blue: 0}
	handfuls := strings.Split(line, ";")
	fmt.Println("Handfuls:", handfuls)
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
			switch color {
			case "red":
				if amount > cubeMin.red {
					cubeMin.red = amount
				}
			case "green":
				if amount > cubeMin.green {
					cubeMin.green = amount
				}
			case "blue":
				if amount > cubeMin.blue {
					cubeMin.blue = amount
				}
			}
		}
	}
	return cubeMin.red * cubeMin.green * cubeMin.blue
}

func main() {
	// Testing
	// string := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	// fmt.Println(cubeSetParser(string))

	// Task
	var filePath string = "files/day02"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var sumOfPowers = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var gameIdAndRest = strings.Split(line, ":")
		var restOfLine string = gameIdAndRest[1]
		sumOfPowers += gamePowerCalculator(restOfLine)
	}
	fmt.Println(sumOfPowers)
}
