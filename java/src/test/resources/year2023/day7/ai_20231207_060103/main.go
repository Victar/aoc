
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	FiveOfAKind   = 7
	FourOfAKind   = 6
	FullHouse     = 5
	ThreeOfAKind  = 4
	TwoPair       = 3
	OnePair       = 2
	HighCard      = 1
)

var cardStrength = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type Hand struct {
	Cards  string
	Bid    int
	Score  int
	Values []int
}

func (h *Hand) calculateScore() {
	// Count the frequency of each card
	counts := make(map[rune]int)
	for _, c := range h.Cards {
		counts[c]++
	}

	// Create a slice of values for sorting
	for k, v := range counts {
		h.Values = append(h.Values, v*100+cardStrength[k])
	}

	// Sort the slice of values in decreasing order
	sort.Slice(h.Values, func(i, j int) bool {
		return h.Values[i] > h.Values[j]
	})

	// Determine the score of the hand based on frequency of cards (hand strength)
	if len(counts) == 2 {
		if h.Values[0]/100 == 4 {
			h.Score = FourOfAKind
		} else {
			h.Score = FullHouse
		}
	} else if len(counts) == 3 {
		if h.Values[0]/100 == 3 {
			h.Score = ThreeOfAKind
		} else {
			h.Score = TwoPair
		}
	} else if len(counts) == 4 {
		h.Score = OnePair
	} else {
		h.Score = HighCard
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var hands []Hand
	for scanner.Scan() {
		cards := scanner.Text()
		scanner.Scan()
		bidStr := scanner.Text()
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			panic(err)
		}
		hand := Hand{
			Cards: cards,
			Bid:   bid,
		}
		hand.calculateScore()
		hands = append(hands, hand)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Sort the hands based on their calculated scores and their values
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Score == hands[j].Score {
			// Compare values if the hand types are the same
			for k := range hands[i].Values {
				if hands[i].Values[k] != hands[j].Values[k] {
					return hands[i].Values[k] > hands[j].Values[k]
				}
			}
		}
		return hands[i].Score > hands[j].Score
	})

	// Calculate the total winnings
	totalWinnings := 0
	for rank, hand := range hands {
		totalWinnings += hand.Bid * (rank + 1)
	}

	fmt.Println(totalWinnings)
}
