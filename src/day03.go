package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func getPartNumber(grid [][]rune, ignoredRunes []rune) {
	var sumAdjecentSymbolNumbers int = 0

	return sumAdjecentSymbolNnumbers
}

func main() {
	var filePath string = "files/day03"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var grid [][]rune
	var ignoredRunes []rune = []rune{46, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57} // . 0 1 2 3 4 5 6 7 8 9
	// fmt.Printf("%c\n", ignoredRunes[0])

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var runeArray []rune = []rune(line)
		grid = append(grid, runeArray)
	}

	var partNumber int = getPartNumber(grid, ignoredRunes)
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
