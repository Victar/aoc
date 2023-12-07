
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CardCounter map[rune]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := []string{}
	bids := []int{}

	// Read data from input.txt and separate hands from bids
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		for i := 0; i < len(fields); i += 2 {
			hands = append(hands, fields[i])
			bid, _ := strconv.Atoi(fields[i+1])
			bids = append(bids, bid)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Rank the hands and calculate total winnings
	totalWinnings := calculateWinningsWithJokers(hands, bids)

	// Output the answer
	fmt.Println(totalWinnings)
}

// Sort the hands by their strength, considering the new joker rule
func calculateWinningsWithJokers(hands []string, bids []int) int {
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j]) > 0
	})

	winnings := 0
	for i, hand := range hands {
		winnings += (i + 1) * bids[i]
	}

	return winnings
}

// Compare two hands with the new joker rule, returning positive if hand1 is stronger, 
// negative if hand2 is stronger, and zero if they are equally strong
func compareHands(hand1, hand2 string) int {
	strength1, ranks1 := handStrength(hand1)
	strength2, ranks2 := handStrength(hand2)
	if strength1 != strength2 {
		return strength1 - strength2
	}
	for i := 0; i < len(ranks1) && i < len(ranks2); i++ {
		if ranks1[i] != ranks2[i] {
			return ranks1[i] - ranks2[i]
		}
	}
	return 0
}

// Determine the strength and ranking of a hand with the joker rule applied
func handStrength(hand string) (int, []int) {
	counts := countCards(hand)

	// Find the best possible hand given the counts and presence of Jokers
	isJokerPresent := counts['J'] > 0
	delete(counts, 'J')
	cards := make([]int, 0, len(counts))
	for _, cnt := range counts {
		cards = append(cards, cnt)
	}
	sort.Slice(cards, func(i, j int) bool {
		return cards[i] > cards[j]
	})

	strength := 0
	ranks := cardRanks(counts, cards)

	// Determine the strength of the hand based on counts
	switch {
	case cards[0] == 5 || (cards[0] == 4 && isJokerPresent):
		strength = 7 // Five of a kind
	case cards[0] == 4:
		strength = 6 // Four of a kind
	case (cards[0] == 3 && cards[1] == 2) || (cards[0] == 3 && cards[1] == 1 && isJokerPresent):
		strength = 5 // Full house
	case cards[0] == 3:
		strength = 4 // Three of a kind
	case cards[0] == 2 && cards[1] == 2:
		strength = 3 // Two pair
	case cards[0] == 2:
		strength = 2 // One pair
	default:
		strength = 1 // High card
	}

	return strength, ranks
}

// Count the occurrences of each card in a hand
func countCards(hand string) CardCounter {
	counter := make(CardCounter)
	for _, card := range hand {
		counter[card]++
	}
	return counter
}

// Generate a slice of integers representing the ranks of the cards, from highest to lowest
func cardRanks(counts CardCounter, cards []int) []int {
	ranks := make([]int, 0, len(counts))
	cardValue := map[rune]int{
		'A': 14, 'K': 13, 'Q': 12, 'T': 10, '9': 9,
		'8': 8, '7': 7, '6': 6, '5': 5, '4': 4,
		'3': 3,'2': 2,  'J': 1,
	}
	for card, count := range counts {
		for i := 0; i < count; i++ {
			ranks = append(ranks, cardValue[card])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))
	return ranks
}
