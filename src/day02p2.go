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

func cubeSetParser(line string) CubeSet {
	cubes := CubeSet{red: 0, green: 0, blue: 0}
	handfuls := strings.Split(strings.Split(line, ":")[1], ";") // Removing the Game Id prefix and splitting across the ";" seperator
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

			switch color {
			case "red":
				cubes.red += amount
			case "green":
				cubes.green += amount
			case "blue":
				cubes.blue += amount
			}
		}
	}
	fmt.Println("Line:", line)
	fmt.Println(cubes, "\n")
	return cubes
}

func possibilityOracle(cubeSet CubeSet) bool {
	var cubeMax CubeSet = CubeSet{red: 12, green: 13, blue: 14}
	var maxTotalCubes int = cubeMax.red + cubeMax.green + cubeMax.blue
	var cubeSetTotal int = cubeSet.red + cubeSet.green + cubeSet.blue
	if cubeSet.red <= cubeMax.red && cubeSet.green <= cubeMax.green && cubeSet.blue <= cubeMax.blue && maxTotalCubes >= cubeSetTotal {
		return true
	} else {
		return false
	}
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

	// cubeMax := CubeSet{red: 12, green: 13, blue: 14}
	var possibleGamesSum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newCubeSet := cubeSetParser(line)
		if possibilityOracle(newCubeSet) {
			possibleGamesSum++
		}
	}
	fmt.Println(possibleGamesSum)

}
