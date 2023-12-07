package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func doesRuneContain(currentRune rune, allowedRunes []rune) bool {
	for _, allowedRune := range allowedRunes {
		if currentRune == allowedRune {
			return true
		}
	}
	return false
}

func discoverWholeNumber(grid [][]rune, row int, col int) int {
	// Check for invalid indices
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
		return 0
	}

	// Start from the given position and move leftwards to find the start of the number
	start := col
	for start > 0 && unicode.IsDigit(grid[row][start-1]) {
		start--
	}

	// Accumulate digits into a string
	numStr := ""
	for i := start; i < len(grid[row]) && unicode.IsDigit(grid[row][i]); i++ {
		numStr += string(grid[row][i])
	}

	// Convert the string to an integer
	wholeNumber, err := strconv.Atoi(numStr)
	if err != nil {
		return 0 // Return 0 if conversion fails
	}
	return wholeNumber
}

func checkGearedNumbers(grid [][]rune, row int, col int) ([]int, bool) {
	var numeral []rune = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	// Checking location of the index relative to the grid
	var topRow bool = row == 0
	var bottomRow bool = row == len(grid)-1
	var leftCol bool = col == 0
	var rightCol bool = col == len(grid[row])-1
	var partNumbers []int = []int{}

	// Check surrounding cells
	if grid[row][col] == rune('*') {
		for i := -1; i <= 1; i++ { // i = Row
			for j := -1; j <= 1; j++ { // j = Col
				// Skip checking all the cells that don't exist
				if i == 0 && j == 0 { // Skip the current index
					continue
				}
				if leftCol && j == -1 { // Skip the left column
					continue
				}
				if rightCol && j == 1 { // Skip the right column
					continue
				}
				if topRow && i == -1 { // Skip the top row
					continue
				}
				if bottomRow && i == 1 { // Skip the bottom row
					continue
				}
				// Skipped all cells that don't exist
				// Now check for a numberal around the current index, and add the number to the partNumbers slice
				if doesRuneContain(grid[row+i][col+j], numeral) {
					var addNumber int = discoverWholeNumber(grid, row+i, col+j)
					// Check if the number is already in the slice
					var numberAlreadyInSlice bool = false
					for _, number := range partNumbers {
						if number == addNumber {
							numberAlreadyInSlice = true
						}
					}
					if !numberAlreadyInSlice && addNumber != 0 {
						partNumbers = append(partNumbers, addNumber)
					}
				}
			}
		}
		if len(partNumbers) != 2 {
			return partNumbers, false
		}

		return partNumbers, true
	}
	fmt.Println("Error: checkGearedNumers() called on a non-gear rune, rune:", grid[row][col], "at row:", row, "col:", col)
	return nil, false
}

func getGearedPartNumbers(grid [][]rune) [][]int {
	var gearRune rune = rune('*')
	var gearedPartNumbers [][]int
	// Check each row rune by rune, until a gear is found
	for y, row := range grid { // row = {467..114..}
		for x, currentRune := range row { // r = {4}, {$}, {.}, {-}
			// The current rune is a numeral
			if currentRune == gearRune {
				var addGearedNumbers, foundNumbers = checkGearedNumbers(grid, y, x)
				if foundNumbers && len(addGearedNumbers) == 2 && addGearedNumbers[0] != 0 && addGearedNumbers[1] != 0 { // Making sure that there are two of them and that they are not zero
					fmt.Println("Adding set of geared numbers to gearedPartNumbers:", addGearedNumbers)
					gearedPartNumbers = append(gearedPartNumbers, addGearedNumbers)
				}
			}
		}
	}
	return gearedPartNumbers
}

func main() {
	var filePath string = "files/day03"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var runeArray []rune = []rune(line)
		grid = append(grid, runeArray)
	}

	// var partNumbers []int = getPartNumbers(grid)
	var gearRatioPartNumers [][]int = getGearedPartNumbers(grid)

	// Summing the geared pair numbers with each other
	var gearedPairSum int = 0
	for _, gearedPair := range gearRatioPartNumers {
		gearedPairSum += gearedPair[0] * gearedPair[1]
	}
	fmt.Println("Geared pair sum:", gearedPairSum)
}
