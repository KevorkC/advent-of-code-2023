package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type hand struct {
	cards string
	bid   int
	_type int
}

func cardsToType(cards string) int {
	// Creating a hashMap to store the count of each card
	var cardMap = make(map[rune]int)
	for _, card := range cards {
		cardMap[card]++
	}

	// Finding how many unique cards there are
	var uniqueCards int = len(cardMap)

	// Five of a kind = 7, Four of a kind = 6, Full house = 5, Three of a kind = 4, Two pair = 3, One pair = 2, High card = 1

	// Looping through the hashMap to find the type of the hand
	for _, value := range cardMap {
		if value == 5 {
			//fmt.Println(cards, "= Five of a kind")
			return 7
		} else if value == 4 {
			//fmt.Println(cards, "= Four of a kind")
			return 6
		} else if value == 3 && uniqueCards == 2 {
			//fmt.Println(cards, "= Full house")
			return 5
		} else if value == 3 && uniqueCards == 3 {
			//fmt.Println(cards, "= Three of a kind")
			return 4
		} else if value == 2 && uniqueCards == 3 {
			//fmt.Println(cards, "= Two pair")
			return 3
		} else if value == 2 && uniqueCards == 4 {
			//fmt.Println(cards, "= One pair")
			return 2
		} else if value == 1 && uniqueCards == 5 {
			//fmt.Println(cards, "= High card")
			return 1
		}
	}
	return -1
}

type CardStrength rune

const (
	Two CardStrength = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Joker
	Queen
	King
	Ace
)

var order = []CardStrength{Ace, King, Queen, Joker, Ten, Nine, Eight, Seven, Six, Five, Four, Three, Two}

func runeToCardStrength(r rune) CardStrength {
	switch r {
	case 'A':
		return Ace
	case 'K':
		return King
	case 'Q':
		return Queen
	case 'J':
		return Joker
	case 'T':
		return Ten
	case '9':
		return Nine
	case '8':
		return Eight
	case '7':
		return Seven
	case '6':
		return Six
	case '5':
		return Five
	case '4':
		return Four
	case '3':
		return Three
	case '2':
		return Two
	default:
		panic("Invalid rune")
	}
}

func sortHands(hands []hand) []hand {

	// Sort the hands from lowest to highest type
	for i := 0; i < len(hands); i++ {
		for j := 0; j < len(hands)-1; j++ {
			if hands[j]._type > hands[j+1]._type {
				hands[j], hands[j+1] = hands[j+1], hands[j]
			}
		}
	}

	// Splitting the hands in to lists of the same type
	typeMap := make(map[int][]hand)
	for _, hand := range hands {
		typeMap[hand._type] = append(typeMap[hand._type], hand)
	}

	for _, handList := range typeMap {
		sort.Slice(handList, func(iHandA, iHandB int) bool {
			handA := handList[iHandA]
			handB := handList[iHandB]
			for i := 0; i < len(handA.cards); i++ {
				if runeToCardStrength(rune(handA.cards[i])) == runeToCardStrength(rune(handB.cards[i])) {
					continue
				}
				return runeToCardStrength(rune(handA.cards[i])) < runeToCardStrength(rune(handB.cards[i]))
			}
			return true
		})
	}

	var sortedHandsList []hand
	for i := 1; i <= 7; i++ {
		sortedHandsList = append(sortedHandsList, typeMap[i]...)
	}

	// Printing the list
	for _, hand := range sortedHandsList {
		fmt.Println(hand.cards, hand._type)
	}
	println("")

	return sortedHandsList
}

func stringtoHand(line string) hand {
	var newHand hand
	newHand.cards = line[:5]
	bid, e := strconv.Atoi(line[6:])
	if e != nil {
		fmt.Println("Error converting bid to int:", e)

	}
	newHand.bid = bid
	newHand._type = cardsToType(newHand.cards)
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

	var orderedHands []hand = sortHands(hands)

	// Calculating the total winnings for all the hands,
	// by multiplying the rank with the bid
	var totalWinnings int = 0
	for i, hand := range orderedHands {
		totalWinnings += hand.bid * (i + 1)
	}

	fmt.Println("Total Winnings:", totalWinnings)

}
