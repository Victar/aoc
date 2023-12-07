
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card string
type Hand struct {
	Cards     []Card
	BidAmount int
	Rank      int
}

// EvaluateStrength returns a unique score for the hand based on its type and card strength
func (h *Hand) EvaluateStrength(cardValues map[Card]int) int {
	counts := make(map[Card]int)
	for _, card := range h.Cards {
		counts[card]++
	}

	typeScore, card1, card2 := 0, Card('0'), Card('0')

	switch len(counts) {
	case 1:
		typeScore = 7000000 // Five of a kind
	case 2:
		for card, count := range counts {
			if count == 4 {
				card1 = card
				typeScore = 6000000 // Four of a kind
				break
			}
			if count == 3 {
				card1 = card
				typeScore = 5000000 // Full house
			} else {
				card2 = card
			}
		}
	case 3:
		for card, count := range counts {
			if count == 3 {
				card1 = card
				typeScore = 4000000 // Three of a kind
				break
			}
			if count == 2 {
				if card1 == '0' || cardValues[card] > cardValues[card1] {
					card2 = card1
					card1 = card
				} else if card2 == '0' || cardValues[card] > cardValues[card2] {
					card2 = card
				}
			}
		}
		if card2 != '0' {
			typeScore = 3000000 // Two pair
		}
	case 4:
		for card, count := range counts {
			if count == 2 {
				card1 = card
				break
			}
		}
		typeScore = 2000000 // One pair
	case 5:
		sortedCards := make([]Card, 0, 5)
		for card := range counts {
			sortedCards = append(sortedCards, card)
		}
		sort.Slice(sortedCards, func(i, j int) bool {
			return cardValues[sortedCards[i]] > cardValues[sortedCards[j]]
		})
		card1 = sortedCards[0]
		card2 = sortedCards[1]
		typeScore = 1000000 // High card
	}

	score := typeScore + cardValues[card1]*10000 + cardValues[card2]*100
	return score
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	cardValues := map[Card]int{
		"A": 14, "K": 13, "Q": 12, "J": 11, "T": 10,
		"9": 9, "8": 8, "7": 7, "6": 6, "5": 5,
		"4": 4, "3": 3, "2": 2,
	}
	var hands []Hand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		splitInput := strings.Split(inputLine, " ")
		if len(splitInput)%2 != 0 {
			fmt.Fprintf(os.Stderr, "invalid input format\n")
			os.Exit(1)
		}
		for i := 0; i < len(splitInput); i += 2 {
			bid, _ := strconv.Atoi(splitInput[i+1])
			hand := Hand{
				Cards:     strings.Split(splitInput[i], ""),
				BidAmount: bid,
			}
			hands = append(hands, hand)
		}
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].EvaluateStrength(cardValues) > hands[j].EvaluateStrength(cardValues)
	})

	totalWinnings := 0
	for i, hand := range hands {
		hand.Rank = i + 1
		totalWinnings += hand.BidAmount * hand.Rank
	}

	fmt.Println(totalWinnings)

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error scanning file: %v\n", err)
	}
}

