package main

import (
	"bufio"
	"fmt"
	"os"
)



// func checkNumberSourroundingsForSymbol(grid [][]rune, row int, col int, numberDigits int, ignoredRunes []rune) bool {
// 	// Size of the grid
// 	var rowLength int = len(grid)
// 	// var colLength int = len(grid[0])

// 	// Is the number located in the beginning or the end of the row?
// 	beginning, end := false, false
// 	if row == 0 {
// 		beginning = true
// 		fmt.Println("Beginning =", beginning)
// 	} else if row+numberDigits == rowLength {
// 		end = true
// 		fmt.Println("End =", end)
// 	}

// 	// Checking the above row, then the current row, then the row below

// 	return beginning
// }

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

func doesRuneContain(currentRune rune, allowedRunes []rune) bool {
	for _, allowedRune := range allowedRunes {
		if currentRune == allowedRune {
			return true
		}
	}
	return false
}

func checkForSymbolAroundIndex(grid [][]rune, row int, col int) bool {
	var symbolFound bool = false
	var symbol []rune = []rune{'#', '$', '%', '&', '*', '+', '-', '/', '?', '@'}

	// Checking location of the index relative to the grid
	var topRow bool = row == 0
	var bottomRow bool = row == len(grid)-1
	var leftCol bool = col == 0
	var rightCol bool = col == len(grid[row])-1

    // Check surrounding cells
    for i := -1; i <= 1; i++ {
		// If the index is not on the top or bottom row, check the row above and below
        if !topRow && doesRuneContain(grid[row-1][col+i], symbol) && col+i >= 0 && col+i < len(grid[row-1]) {
            return true
        }
        if !bottomRow && doesRuneContain(grid[row+1][col+i], symbol) && col+i >= 0 && col+i < len(grid[row+1]) {
            return true
        }
    }

    if !leftCol && doesRuneContain(grid[row][col-1], symbol) {
        return true
    }
    if !rightCol && doesRuneContain(grid[row][col+1], symbol) {
        return true
    }

	return symbolFound
}

func getPartNumber(grid [][]rune) int {
	var ignoredRunes []rune = []rune{'.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'} // 46, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57
	var numeralRunes []rune = ignoredRunes[1:]
	var sumAdjecentSymbolNumbers int = 0
	// Check each row rune by rune, until a numeral is found
	fmt.Println("Numerals:", numeralRunes)
	for y, row := range grid { // row = {467..114..}
		var foundNumeral bool = false
		var tempNumerals string = "" // 467
		for x, currentRune := range row {      // r = {4}, {$}, {.}, {-}
			if doesRuneContain(currentRune, numeralRunes) {
				foundNumeral = checkForSymbolAroundIndex(grid, y, x)
		}
		fmt.Println(tempNumerals)
	}

	return 0
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
