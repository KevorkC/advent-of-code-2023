package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// countMatches counts how many numbers match between the scratchcard's own numbers and the drawn numbers
func countMatches(ownNumbers, drawnNumbers []string) int {
	matches := 0
	for _, own := range ownNumbers {
		for _, drawn := range drawnNumbers {
			if own == drawn {
				matches++
				break // Stop searching if a match is found for the current number
			}
		}
	}
	return matches // Return the total count of matches
}

// processCards processes the list of scratchcards and calculates the total number of cards, including copies
func processCards(cardList []string) int {
	cardCopies := make([]int, len(cardList)) // Initialize a slice to track the number of copies for each card
	for i := range cardCopies {
		cardCopies[i] = 1 // Initially, there is one copy of each card
	}

	totalCards := 0
	for i, card := range cardList {
		parts := strings.Split(card, "|") // Split each card's data into own numbers and drawn numbers
		if len(parts) != 2 {
			fmt.Println("Invalid card format:", card)
			continue // Skip this card if the format is incorrect
		}
		ownNumbers := strings.Fields(parts[0])   // Extract the card's own numbers
		drawnNumbers := strings.Fields(parts[1]) // Extract the drawn numbers
		matches := countMatches(ownNumbers, drawnNumbers)

		// Update the count of copies for subsequent cards based on the number of matches
		for j := 1; j <= matches; j++ {
			if (i + j) < len(cardList) { // Check that the index is within the bounds of the list
				cardCopies[i+j] += cardCopies[i] // Add the number of copies of this and previous cards to the next card
			}
		}
		totalCards += cardCopies[i] // Add the number of copies of this card to the total count
	}

	fmt.Println("Card copies:", cardCopies)
	return totalCards // Return the total number of scratchcards, including copies
}

func main() {
	filePath := "files/day04"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cardList []string
	for scanner.Scan() {
		cardList = append(cardList, scanner.Text())
	}

	totalCards := processCards(cardList)
	fmt.Println("Total number of scratchcards:", totalCards)
}
