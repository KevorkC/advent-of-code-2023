// Can't account for the case of "twone", returns (2,2)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	first  int
	second int
}

func firstAndLastInt(line string) Pair {
	var pair Pair = Pair{-1, -1}
	// fmt.Println("Line:", line)
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

	// List of all numbers between 0 and 9 in both string and int form
	var numberList []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// *** Finding The First Number From The Front ***
	// Loop from the front to end, each time adding one more character to the search string
	for i := 0; i <= len(line); i++ {
		var assignedFirstNumbner bool = false
		for _, number := range numberList {
			// fmt.Println("Searching", line[:i], "for", number)
			if len(line[:i]) >= len(number) && strings.Contains(line[:i], number) {
				// Found the number
				if len(number) > 1 {
					// Look up in the map first
					var firstNumber, err = strconv.Atoi(wordToNumber[number])
					if err != nil {
						fmt.Println("Error converting word to number:", err)
					}
					// fmt.Println("Assigning first number to", firstNumber, "from", number, "in", line)
					pair.first = firstNumber
					assignedFirstNumbner = true
					break
				} else {
					var firstNumber, err = strconv.Atoi(number)
					if err != nil {
						fmt.Println("Error converting word to number:", err)
					}
					// fmt.Println("Assigning first number to", firstNumber, "from", number, "in", line)
					pair.first = firstNumber
					assignedFirstNumbner = true
					break
				}
			}
		}
		if assignedFirstNumbner {
			break
		}
	}

	// *** Finding The First Number From The Back ***
	for i := len(line) - 1; i >= 0; i-- {
		var assignedSecondNumber bool = false
		for _, number := range numberList {
			// fmt.Println("Going to check", line[i:], "for", number)
			if len(line[i:]) >= len(number) && strings.Contains(line[i:], number) {
				// fmt.Println("Found the number", number, "in", line[i:])
				// Found the number
				if len(number) > 1 {
					// Look up in the map first
					var secondNumber, err = strconv.Atoi(wordToNumber[number])
					if err != nil {
						fmt.Println("Error converting word to number:", err)
					}
					// fmt.Println("Assigning second number to", secondNumber, "from", number, "in", line)
					pair.second = secondNumber
					assignedSecondNumber = true
					break
				} else {
					var secondNumber, err = strconv.Atoi(number)
					if err != nil {
						fmt.Println("Error converting word to number:", err)
					}
					// fmt.Println("Assigning second number to", secondNumber, "from", number, "in", line)
					pair.second = secondNumber
					assignedSecondNumber = true
					break
				}
			}
		}
		if assignedSecondNumber {
			break
		}
	}

	if pair.first == -1 || pair.second == -1 {
		fmt.Println("ERROR Type -1: First =", pair.first, "Second =", pair.second, "Line =", line)
	}
	return pair
}

func main() {

	// Tests
	testAll()

	// Task
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

		pairList = append(pairList, firstAndLastInt(line))
	}

	var total_sum int = 0
	for _, pair := range pairList {
		total_sum += pair.first*10 + pair.second
	}

	fmt.Println(total_sum)
}

func testAll() {
	testAdd()
	testAdd2()
	testAdd3()
	testAdd4()
}

func testAdd() {
	got := firstAndLastInt("twone")
	want := Pair{2, 1}

	if got != want {
		fmt.Println("TEST1 - Got:", got, "Want:", want)
	}
}

func testAdd2() {
	got := firstAndLastInt("xtwonexonextwonex")
	want := Pair{2, 1}

	if got != want {
		fmt.Println("TEST2 - Got:", got, "Want:", want)
	}
}

func testAdd3() {
	got := firstAndLastInt("4twone5")
	want := Pair{4, 5}

	if got != want {
		fmt.Println("TEST3 - Got:", got, "Want:", want)
	}
}

func testAdd4() {
	got := firstAndLastInt("xsixxseven9eighh")
	want := Pair{6, 9}

	if got != want {
		fmt.Println("TEST4 - Got:", got, "Want:", want)
	}
}
