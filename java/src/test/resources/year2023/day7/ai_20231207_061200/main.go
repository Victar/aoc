
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// CardOrderMaps holds the strength of each card label
var CardOrderMaps = map[rune]int{
	'A': 13, 'K': 12, 'Q': 11, 'J': 10, 'T': 9,
	'9': 8, '8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1,
}

// HandType represents the type of hand in Camel Cards game
type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

// Hand represents a Camel Cards hand and its bid
type Hand struct {
	Cards []rune
	Bid   int
	Type  HandType
	Rank  int
}

// classifyHandType determines the type of a hand
func classifyHandType(hand *Hand) {
	labels := make(map[rune]int)
	for _, card := range hand.Cards {
		labels[card]++
	}
	switch len(labels) {
	case 5:
		hand.Type = HighCard
	case 4:
		hand.Type = OnePair
	case 3:
		for _, count := range labels {
			if count == 3 {
				hand.Type = ThreeOfAKind
				return
			}
		}
		hand.Type = TwoPair
	case 2:
		for _, count := range labels {
			if count == 4 {
				hand.Type = FourOfAKind
				return
			}
		}
		hand.Type = FullHouse
	case 1:
		hand.Type = FiveOfAKind
	}
}

// compareHands orders the hands higher to lower based on their types and card strength
func compareHands(a, b *Hand) bool {
	if a.Type != b.Type {
		return a.Type > b.Type
	}
	for i := 0; i < 5; i++ {
		if a.Cards[i] != b.Cards[i] {
			return CardOrderMaps[a.Cards[i]] > CardOrderMaps[b.Cards[i]]
		}
	}
	return false
}

// sortHands sorts the hands and sets the ranks on them
func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(&hands[i], &hands[j])
	})
	for i, _ := range hands {
		hands[i].Rank = len(hands) - i
	}
}

// parseInput parses the input file and returns the hands with their bids
func parseInput(filename string) ([]Hand, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		bid, _ := strconv.Atoi(parts[1])
		hand := Hand{
			Cards: strings.Split(parts[0], ""),
			Bid:   bid,
		}
		hands = append(hands, hand)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return hands, nil
}

func main() {
	// Assume the input file is in the same directory as the main.go
	hands, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input file: %s\n", err)
		os.Exit(1)
	}

	// Process and classify each hand type
	for i := range hands {
		classifyHandType(&hands[i])
	}

	sortHands(hands)

	// Calculate and print the total winnings
	totalWinnings := 0
	for _, hand := range hands {
		totalWinnings += hand.Bid * hand.Rank
	}
	fmt.Println(totalWinnings)
}
