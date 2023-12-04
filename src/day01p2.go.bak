// Can't account for the case of "twone", returns (2,2)

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
	// fmt.Println("\nOriginal line:  ", originalLine)

	var replacedFirst string = replaceFirstWord(originalLine)
	// fmt.Println("Replaced first: ", replacedFirst)

	var replacedBoth string = replaceLastWord(replacedFirst)
	// fmt.Println("Literal line:   ", replacedBoth)

	return replacedBoth
}

func replaceFirstWord(originalLine string) string {
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

	for i := 0; i < len(originalLine); i++ {
		for word, number := range wordToNumber {
			if len(word) <= len(originalLine[i:]) && originalLine[i:i+len(word)] == word {
				return originalLine[:i] + number + originalLine[i+len(word):]
			}
		}
	}
	return originalLine // Return the original string if no replacement is done
}

func replaceLastWord(originalLine string) string {
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

	for i := len(originalLine); i >= 0; i-- {
		for word, number := range wordToNumber {
			if i-len(word) >= 0 && originalLine[i-len(word):i] == word {
				return originalLine[:i-len(word)] + number + originalLine[i:]
			}
		}
	}
	return originalLine // Return the original string if no replacement is done
}

func main() {
	var filePath string = "files/day01"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close() // Close file when main() ends or program crashes

	// List of pairs
	var pairList []Pair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println("\nOriginal Line:    ", line)
		// Convert each word to its number equivalent
		var literalLine string = lineLiteralizer(line)

		fmt.Println("Literalized line: ", literalLine)
		pairList = append(pairList, firstAndLastInt(literalLine))
	}

	var total_sum int = 0
	for _, pair := range pairList {
		total_sum += pair.first*10 + pair.second
	}

	fmt.Println(total_sum)

	// const testStr string = "xtwonexonextwonex"
	const testStr string = "twone"
	fmt.Println(testStr)
	var replaced string = replaceFirstWord(testStr)
	fmt.Println(replaced)
	replaced = replaceLastWord(replaced)
	fmt.Println(replaced)
	fmt.Println(firstAndLastInt(replaced))
}
