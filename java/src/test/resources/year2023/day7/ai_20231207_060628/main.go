
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Card struct {
	Value string
	Count int
}

type Hand struct {
	Cards       []Card
	Bid         int
	StrongValue string
}

func main() {
	hands, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading the input file:", err)
		return
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j])
	})

	totalWinnings := 0
	for rank, hand := range hands {
		totalWinnings += hand.Bid * (len(hands) - rank)
	}

	fmt.Println(totalWinnings)
}

func readInput(filename string) ([]Hand, error) {
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
		for i := 0; i < len(parts); i += 2 {
			handCards := parts[i]
			bid, _ := strconv.Atoi(parts[i+1])
			sortedHand := sortedHand(handCards)
			hands = append(hands, Hand{
				Cards:       sortedHand,
				Bid:         bid,
				StrongValue: string(sortedHand[0].Value),
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hands, nil
}

func compareHands(h1, h2 Hand) bool {
	h1Type := getHandType(h1.Cards)
	h2Type := getHandType(h2.Cards)

	if h1Type != h2Type {
		return h1Type > h2Type
	}

	for i := range h1.Cards {
		if h1.Cards[i].Value != h2.Cards[i].Value {
			return cardRank(h1.Cards[i].Value) > cardRank(h2.Cards[i].Value)
		}
	}
	return false
}

func sortedHand(cards string) []Card {
	cardCounter := make(map[rune]int)
	for _, c := range cards {
		cardCounter[c]++
	}

	var hand []Card
	for value, count := range cardCounter {
		hand = append(hand, Card{Value: string(value), Count: count})
	}

	sort.Slice(hand, func(i, j int) bool {
		if hand[i].Count == hand[j].Count {
			return cardRank(hand[i].Value) > cardRank(hand[j].Value)
		}
		return hand[i].Count > hand[j].Count
	})

	return hand
}

func cardRank(card string) int {
	cardOrder := "A23456789TJQK"
	return strings.Index(cardOrder, card)
}

func getHandType(cards []Card) int {
	switch {
	case cards[0].Count == 5:
		return 7
	case cards[0].Count == 4:
		return 6
	case cards[0].Count == 3 && cards[1].Count == 2:
		return 5
	case cards[0].Count == 3:
		return 4
	case cards[0].Count == 2 && cards[1].Count == 2:
		return 3
	case cards[0].Count == 2:
		return 2
	default:
		return 1
	}
}
