package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type hand struct {
	cards   string
	bid     int
	rank    int
	winning int
}

// func cardsToRank(cards string) int {

// }

func stringtoHand(line string) hand {
	var newHand hand
	newHand.cards = line[:5]
	bid, e := strconv.Atoi(line[6:])
	if e != nil {
		fmt.Println("Error converting bid to int:", e)

	}
	newHand.bid = bid
	fmt.Println(newHand.cards, bid)
	// newHand.rank = cardsToRank(newHand.cards)
	newHand.winning = newHand.rank * newHand.bid

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
}
