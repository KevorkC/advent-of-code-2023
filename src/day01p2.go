package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type Pair struct {
	first  int
	second int
}

func firstAndLastInt(line string) Pair {
	pair := Pair{first: -1, second: -1}

	var pointerLocA int = 0
	var pointerLocB int = len(line) - 1

	for pointerLocA <= pointerLocB {
		if pair.first == -1 && unicode.IsDigit(rune(line[pointerLocA])) { // True if pointerLocA is currently pointing to an integer
			pair.first = int(line[pointerLocA] - '0') // The - '0' is to convert ASCII to number
		} else if pair.first == -1 {
			pointerLocA++
		}

		if pair.second == -1 && unicode.IsDigit(rune(line[pointerLocB])) { // True if pointerLocA is currently pointing to an integer
			pair.second = int(line[pointerLocB] - '0') // The - '0' is to convert ASCII to number
		} else if pair.second == -1 {
			pointerLocB--
		}

		if pair.first != -1 && pair.second != -1 {
			return pair
		}
	}

	fmt.Printf("Warning: Loop exited when either first or second variable in pair is still -1, the line was %s\n", line)
	return pair
}

func lineLiteralizer(originalLine string) string {
	wordToNumber := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	fmt.Println("Original line:", originalLine)
	fmt.Println("Literal line:", literalLine)
	return literalLine
}

func main() {
	var filePath string = "files/day01"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close() // Close file when main() ends or program crashes

	// const testStr string = "lkj8asd2j"
	// fmt.Println(firstAndLastInt(testStr))

	// List of pairs
	var pairList []Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert each word to its number equivalent
		literalLine := lineLiteralizer(line)

		pairList = append(pairList, firstAndLastInt(literalLine))
	}

	var total_sum int = 0
	for _, pair := range pairList {
		total_sum += pair.first*10 + pair.second
	}

	fmt.Println(total_sum)
}
