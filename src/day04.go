package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filePath string = "files/day04"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Task: Find how many of the winning numbers are on the cards.
	// Each winning number doubles the points given from that card.
	var totalPoints int = 0

	for scanner.Scan() {
		var card string = scanner.Text()
		var winningAndMyNumbers []string = strings.Split(card[9:], " | ")

		var winningNumbers []int
		var myNumbers []int

		for _, winningNumber := range strings.Split(winningAndMyNumbers[0], " ") {
			winningNumber = strings.TrimSpace(winningNumber)
			if winningNumber == "" {
				continue
			}
			var winningNumberInt, err = strconv.Atoi(winningNumber)
			if err != nil {
				fmt.Println("Error - Failed to convert string to intreger:", err, "String:", winningNumber)
			}
			winningNumbers = append(winningNumbers, winningNumberInt)
		}
		for _, myNumber := range strings.Split(winningAndMyNumbers[1], " ") {
			myNumber = strings.TrimSpace(myNumber)
			if myNumber == "" {
				continue
			}
			var myNumberInt, err = strconv.Atoi(myNumber)
			if err != nil {
				fmt.Println("Error - Failed to convert string to intreger:", err, "String:", myNumber)
			}
			myNumbers = append(myNumbers, myNumberInt)
		}

		// Now we have the winning numbers and my numbers in two lists
		// Compare how many from the winning numbers are in my numbers
		var currentCardPoints int = 0
		for _, winningNumber := range winningNumbers {
			// Check if winningNumber exists in myNumber
			for _, mynummyNumber := range myNumbers {
				if winningNumber == mynummyNumber {
					if currentCardPoints == 0 {
						currentCardPoints = 1
					} else {
						currentCardPoints *= 2
					}
				}
			}
		}

		totalPoints += currentCardPoints
	}

	fmt.Println(totalPoints)
}
