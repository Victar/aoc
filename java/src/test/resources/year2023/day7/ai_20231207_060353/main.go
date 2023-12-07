
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Card struct {
	Rank string
}

type Hand struct {
	Cards    []Card
	Bid      int
	Strength int
	Type     HandType
}

func parseInput(filename string) ([]Hand, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		cards, bidStr := parts[:len(parts)-1], parts[len(parts)-1]
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			return nil, err
		}
		hand := Hand{
			Cards: make([]Card, len(cards)),
			Bid:   bid,
		}
		for i, r := range cards {
			hand.Cards[i] = Card{Rank: r}
		}
		hands = append(hands, hand)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return hands, nil
}

// The following functions are implementations of HandType determination and hand comparison.

func main() {
	hands, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Analyze and sort the hands by strength.
	for i := range hands {
		hands[i].Type, hands[i].Strength = analyzeHand(hands[i].Cards)
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j]) > 0
	})

	// Calculate total winnings.
	totalWinnings := 0
	for i, hand := range hands {
		rank := len(hands) - i // Weaker ranks have a lower value (1 for weakest)
		totalWinnings += hand.Bid * rank
	}

	// Print the total winnings as the result
	fmt.Println(totalWinnings)
}

// Below are stubs for the analyzeHand and compareHands functions.
// You need to implement logic for these functions to determine hand type, strength, and to compare them accordingly.

func analyzeHand(cards []Card) (HandType, int) {
	// Implement hand analysis logic to determine type and strength
	// Return the type and strength of the hand 
	return HighCard, 0 // placeholder
}

func compareHands(a, b Hand) int {
	// Implement hand comparison logic
	// Return an integer: positive if a > b, negative if a < b, 0 if equal
	return 0 // placeholder
}
