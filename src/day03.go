package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func doesRuneContain(currentRune rune, allowedRunes []rune) bool {
	for _, allowedRune := range allowedRunes {
		if currentRune == allowedRune {
			return true
		}
	}
	return false
}

func checkForSymbolAroundIndex(grid [][]rune, row int, col int) bool {
	var symbol []rune = []rune{'#', '$', '%', '&', '*', '+', '-', '/', '?', '@'}

	// Checking location of the index relative to the grid
	var topRow bool = row == 0
	var bottomRow bool = row == len(grid)-1
	var leftCol bool = col == 0
	var rightCol bool = col == len(grid[row])-1

    // Check surrounding cells
    for i := -1; i <= 1; i++ { // i = Row
		for j := -1; j <= 1; j++ { // j = Col
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
			if doesRuneContain(grid[row+i][col+j], symbol) {
				return true
			}
		}
	}
	return false
}

/* Example input
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
*/

func getPartNumber(grid [][]rune) int {
	var ignoredRunes []rune = []rune{'.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} // 46, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57
	var numeralRunes []rune = ignoredRunes[1:]
	var sumAdjecentSymbolNumbers int = 0
	// Check each row rune by rune, until a numeral is found
	fmt.Println("Numerals:", numeralRunes)
	for y, row := range grid { // row = {467..114..}
		var foundNumeral bool = false
		var foundSymbolAroundNumeral = false
		var tempNumerals string = "" // 467
		for x, currentRune := range row {      // r = {4}, {$}, {.}, {-}
			isCurrentRuneANumeral := doesRuneContain(currentRune, numeralRunes)
			if isCurrentRuneANumeral {
				foundNumeral = checkForSymbolAroundIndex(grid, y, x)
				tempNumerals = tempNumerals + string(currentRune)
			} else if x == len(grid[y])-1 && foundNumeral && !foundSymbolAroundNumeral {
				foundNumeral = false
				currentNumberToSum, e = strconv.Atoi(tempNumerals)
				if e != nil {
					fmt.Println("Error converting string to int:", e)
				}
				sumAdjecentSymbolNumbers += currentNumberToSum
			}
		fmt.Println(tempNumerals)
	}

	return sumAdjecentSymbolNumbers
}

func main() {
	var filePath string = "files/day03"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var grid [][]rune
	
	// fmt.Printf("%c\n", ignoredRunes[0])

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var runeArray []rune = []rune(line)
		grid = append(grid, runeArray)
	}

	var partNumber int = getPartNumber(grid)
	fmt.Println(partNumber)

	// Printing the grid
	// for _, runeArray := range grid {
	// 	for _, c := range runeArray {
	// 		fmt.Printf("%c", c)
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(grid)
}
