
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Card represents a card with a label and a given strength value.
type Card struct {
	Label   string
	IsJoker bool
	Value   int
}

// Hand represents a set of 5 cards together with its type and leading cards.
type Hand struct {
	Type    int
	Leading []int
	Cards   []Card
	Bid     int
}

// Hand strength constants.
const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	FiveOfAKind
)

// List of all card labels from highest to lowest.
var cardValues = map[string]int{
	"A": 12, "K": 11, "Q": 10, "J": 0, "T": 9, "9": 8, "8": 7, "7": 6, "6": 5, "5": 4, "4": 3, "3": 2, "2": 1,
}

// ReadHands reads the input and returns a slice of Hands.
func ReadHands(filename string) ([]Hand, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		handStr := parts[0]
		bidStr := parts[1]
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			return nil, err
		}
		hand := ParseHand(handStr)
		hand.Bid = bid
		hands = append(hands, hand)
	}
	return hands, scanner.Err()
}

// ParseHand converts a string representation of a hand into a Hand struct.
func ParseHand(handStr string) Hand {
	cards := []Card{}
	for _, label := range handStr {
		isJoker := label == 'J'
		value, _ := cardValues[string(label)]
		cards = append(cards, Card{
			Label:   string(label),
			Value:   value,
			IsJoker: isJoker,
		})
	}
	return Hand{Cards: cards}
}

// GetHandStrength returns the type and leading cards of a hand.
func GetHandStrength(hand *Hand, jokerStrength int) {
	cards := hand.Cards
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].IsJoker {
			return false
		}
		if cards[j].IsJoker {
			return true
		}
		return cards[i].Value > cards[j].Value
	})
	counter := make(map[int]int)
	for _, card := range cards {
		if card.IsJoker {
			continue
		}
		counter[card.Value]++
	}

	// Calculate the actual hand type.
	handType := HighCard
	leading := []int{}
	switch len(counter) {
	case 2: // Full House or Four of a Kind or Five of a Kind
		firstValue := -1
		for value, count := range counter {
			if count == 4 {
				handType = FourOfAKind
				leading = []int{value}
				break
			}
			if count == 3 {
				firstValue = value
			}
			if count == 2 && firstValue != -1 {
				handType = FullHouse
				leading = []int{firstValue, value}
				break
			}
		}
		if handType == HighCard && jokerStrength > 0 { // Five of the same type.
			handType = FiveOfAKind
			for value, count := range counter {
				if count == 4 {
					leading = []int{value}
					break
				}
			}
		}
	case 3: // Three of a Kind or Two Pair
		pairs := []int{}
		for value, count := range counter {
			if count == 3 {
				handType = ThreeOfAKind
				leading = []int{value}
				break
			}
			if count == 2 {
				pairs = append(pairs, value)
			}
		}
		if handType != ThreeOfAKind {
			handType = TwoPair
			sort.Ints(pairs)
			leading = pairs
		}
		if handType == HighCard && jokerStrength > 1 {
			handType = ThreeOfAKind
			for value := range counter {
				leading = append(leading, value)
			}
		}
	case 4: // One Pair
		for value, count := range counter {
			if count == 2 {
				handType = OnePair
				leading = []int{value}
			} else {
				leading = append(leading, value)
			}
		}
		sort.Ints(leading)
		if handType == HighCard && jokerStrength > 2 {
			handType = OnePair
			leading = leading[1:]
		}
	case 5: // High Card
		for value := range counter {
			leading = append(leading, value)
		}
		sort.Ints(leading)
	}

	hand.Type = handType
	hand.Leading = leading
}

// RankHands sorts and ranks a list of hands.
func RankHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		// Check hand types.
		if hands[i].Type != hands[j].Type {
			return hands[i].Type > hands[j].Type
		}
		// Compare leading cards.
		for k := 0; k < len(hands[i].Leading); k++ {
			if hands[i].Leading[k] != hands[j].Leading[k] {
				return hands[i].Leading[k] > hands[j].Leading[k]
			}
		}
		// Compare cards in order.
		for k := 0; k < len(hands[i].Cards); k++ {
			if hands[i].Cards[k].Value != hands[j].Cards[k].Value {
				return hands[i].Cards[k].Value > hands[j].Cards[k].Value
			}
		}
		return false
	})
}

// CalculateWinnings calculates the total winnings of ranked hands.
func CalculateWinnings(hands []Hand) int {
	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += hand.Bid * (i + 1)
	}
	return totalWinnings
}

// AdjustJokerValues modifies the value of jokers for wildcard evaluation.
func AdjustJokerValues(hands []Hand, jokerValue int) {
	for i := range hands {
		for j := range hands[i].Cards {
			if hands[i].Cards[j].IsJoker {
				hands[i].Cards[j].Value = jokerValue
				break
			}
		}
	}
}

func main() {
	filename := "input.txt"
	hands, err := ReadHands(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Part 1
	for i := range hands {
		GetHandStrength(&hands[i], 0)
	}
	RankHands(hands)
	totalWinnings := CalculateWinnings(hands)
	fmt.Println(totalWinnings)

	// Part 2
	hands, _ = ReadHands(filename) // reset the hands.
	for i := range hands {
		GetHandStrength(&hands[i], 8) // set J to 8 (max) for evaluating hand type
	}
	AdjustJokerValues(hands, 0) // set J back to 0 for ranking
	RankHands(hands)
	totalWinnings = CalculateWinnings(hands)
	fmt.Println(totalWinnings)
}
