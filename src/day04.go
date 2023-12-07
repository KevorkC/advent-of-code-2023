package main

import (
	"bufio"
	"fmt"
	"os"
)

/* Example grid:
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
*/

func main() {
	var filePath string = "files/day04"
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

	// Printing the grid
	for _, runeArray := range grid {
		for _, c := range runeArray {
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}

}
