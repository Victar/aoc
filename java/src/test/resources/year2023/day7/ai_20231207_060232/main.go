
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

// Card values ordered from highest to lowest.
var cardValues = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5,
	'4': 4, '3': 3, '2': 2,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		handStr, bidStr := parts[0], parts[1]
		cards := countCards(handStr)
		handType := evaluateHand(cards)
		bid, _ := strconv.Atoi(bidStr)
		hands = append(hands, Hand{cards, handType, bid, 0})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
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

// countCards takes a hand string and returns a slice of Card with their counts.
func countCards(handStr string) []Card {
	counter := make(map[rune]int)
	for _, ch := range handStr {
		counter[ch]++
	}
	var cards []Card
	for value, count := range counter {
		cards = append(cards, Card{value, count})
	}
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Count == cards[j].Count {
			return cardValues[cards[i].Value] > cardValues[cards[j].Value]
		}
		return cards[i].Count > cards[j].Count
	})
	return cards
}

// evaluateHand determines the type of the hand.
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

// compareHands determines if hand1 is stronger than hand2.
func compareHands(hand1, hand2 Hand) bool {
	if hand1.Type != hand2.Type {
		return hand1.Type > hand2.Type
	}
	for i := range hand1.Cards {
		if hand1.Cards[i].Count != hand2.Cards[i].Count ||
			cardValues[hand1.Cards[i].Value] != cardValues[hand2.Cards[i].Value] {
			return cardValues[hand1.Cards[i].Value] > cardValues[hand2.Cards[i].Value]
		}
	}
	return false
}
