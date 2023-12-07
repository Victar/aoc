package main

import (
	"adventofcode/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var DAY = "7"

// Card represents a Poker card with a rank and a face value.
type Card struct {
	Value rune
	Count int
}

// Hand represents a Poker hand with its cards, type and bid amount.
type Hand struct {
	Cards []Card
	Type  int
	Bid   int
	Rank  int
	Input string
}

// HandRanking constants to define hand types.
const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var cardValuesGold = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5,
	'4': 4, '3': 3, '2': 2,
}
var cardValuesSilver = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5,
	'4': 4, '3': 3, '2': 2,
}

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")

	if err != nil {
		panic(err)
	}
	var hands []Hand
	for _, line := range lines {
		parts := strings.Fields(line)
		handStr, bidStr := parts[0], parts[1]
		cards := countCards(handStr, false)
		handType := evaluateHand(cards)
		bid, _ := strconv.Atoi(bidStr)
		hands = append(hands, Hand{cards, handType, bid, 0, line})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], cardValuesSilver)
	})

	for i := range hands {
		hands[i].Rank = len(hands) - i
	}

	totalWinnings := 0
	for _, hand := range hands {
		totalWinnings += hand.Bid * hand.Rank
	}

	fmt.Println(totalWinnings)
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")

	if err != nil {
		panic(err)
	}
	var hands []Hand
	for _, line := range lines {
		parts := strings.Fields(line)
		handStr, bidStr := parts[0], parts[1]
		cards := countCards(handStr, true)
		handType := evaluateHand(cards)
		bid, _ := strconv.Atoi(bidStr)
		hands = append(hands, Hand{cards, handType, bid, 0, line})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], cardValuesGold)
	})

	for i := range hands {
		hands[i].Rank = len(hands) - i
	}

	totalWinnings := 0
	for _, hand := range hands {
		totalWinnings += hand.Bid * hand.Rank
	}

	fmt.Println(totalWinnings)
}

func countCards(handStr string, isGold bool) []Card {
	counter := make(map[rune]int)
	for _, ch := range handStr {
		counter[ch]++
	}
	jCount := 0

	var cards []Card
	for value, count := range counter {
		if isGold && value == 'J' {
			jCount = count
		} else {
			cards = append(cards, Card{value, count})
		}
	}
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Count == cards[j].Count {
			if isGold {
				return cardValuesGold[cards[i].Value] > cardValuesGold[cards[j].Value]
			} else {
				return cardValuesSilver[cards[i].Value] > cardValuesSilver[cards[j].Value]
			}
		}
		return cards[i].Count > cards[j].Count
	})
	if len(cards) > 0 {
		cards[0].Count += jCount
	} else {
		cards = append(cards, Card{'A', 5})
	}
	return cards
}

func evaluateHand(cards []Card) int {
	switch cards[0].Count {
	case 5:
		return FiveOfAKind
	case 4:
		return FourOfAKind
	case 3:
		if cards[1].Count == 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if cards[1].Count == 2 {
			return TwoPair
		}
		return OnePair
	}
	return HighCard
}

func compareHands(hand1, hand2 Hand, cardValues map[rune]int) bool {
	if hand1.Type != hand2.Type {
		return hand1.Type > hand2.Type
	}
	for i := 0; i < 5; i++ {
		card1 := cardValues[rune(hand1.Input[i])]
		card2 := cardValues[rune(hand2.Input[i])]
		if card1 != card2 {
			return card1 > card2
		}

	}
	return false
}
