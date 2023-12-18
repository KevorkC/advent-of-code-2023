package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type hand struct {
	cards   string
	bid     int
	rank    int
	winning int
}

func cardsToRank(cards string) int {
	// Creating a hashMap to store the count of each card
	var cardMap = make(map[rune]int)
	for i, r := range cards {
		cardMap[r] += i
	}

	// Finding how many unique cards there are
	var uniqueCards int = len(cardMap)

	if uniqueCards == 5 {
		return 1
	}

	for _, value := range cardMap {
		if value == 5 {
			return 5
		} else if value == 4 {
			return 4
		} else if value == 3 && uniqueCards == 2 {
			return 3
		} else if value == 3 && uniqueCards == 3 {
			return 2
		} else if value == 2 && uniqueCards == 3 {
			return 2
		}
	}

	return math.MinInt
}

func stringtoHand(line string) hand {
	var newHand hand
	newHand.cards = line[:5]
	bid, e := strconv.Atoi(line[6:])
	if e != nil {
		fmt.Println("Error converting bid to int:", e)

	}
	newHand.bid = bid
	newHand.rank = cardsToRank(newHand.cards)
	newHand.winning = newHand.rank * newHand.bid

	fmt.Println(newHand.cards, newHand.bid, newHand.rank)
	return newHand
}

func main() {
	file, err := os.Open("files/day07")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hands []hand
	for scanner.Scan() {
		var line string = scanner.Text()
		hands = append(hands, stringtoHand(line))
	}

	var totalWinnings = 0
	for _, hand := range hands {
		totalWinnings += hand.winning
	}

	fmt.Println(totalWinnings)
}
