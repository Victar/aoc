
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
	HighCard      = 1
	OnePair       = 2
	TwoPair       = 3
	ThreeOfAKind  = 4
	FullHouse     = 5
	FourOfAKind   = 6
	FiveOfAKind   = 7
	WildCard = 'J'
)

var cardValues = map[string]int{
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
	"A": 14,
}
// Include J as wild card for comparison
for i := 2; i <= 14; i++ {
    if i == 11 {
        continue // Skip the J
    }
    cardOrder = append(cardOrder, string(rune('A'+i)))
}
cardOrder = append(cardOrder, "J") // J is weakest now

type Hand struct {
	Cards     string
	Bid       int
	Strength  int
	HighCards []string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var hands []Hand

	// Read inputs from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for i, part := range parts {
			if i%2 == 0 {
				hands = append(hands, Hand{Cards: part}) // Set Cards
			} else {
				bid, _ := strconv.Atoi(part)
				hands[len(hands)-1].Bid = bid // Set bid
			}
		}
	}

	evaluateHands(hands)
	sortHands(hands)

	totalWinnings := 0
	for rank, hand := range hands {
		totalWinnings += (rank + 1) * hand.Bid
	}

	fmt.Println(totalWinnings)
}

// include functions evaluateHands and sortHands accordingly...

// Add a simplified evaluateHands function
func evaluateHands(hands []Hand) {
	// This function will evaluate each hand and assign Strength and HighCards fields.
}

// Add a simplified sortHands function
func sortHands(hands []Hand) {
	// This function will sort hands first by their Strength and then by their HighCards for tie-breaking.
	// You will need to implement comparison considering Joker-as-wildcard scenario for part two.
}
