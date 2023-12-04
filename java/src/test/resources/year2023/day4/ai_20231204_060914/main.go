
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cards := readCards("input.txt")

	// Part Two logic
	cardCounts := make(map[int]int, len(cards))
	for i := range cards {
		cardCounts[i+1] = 1 // Initialize counts to 1 for each card
	}

	for _, card := range cards {
		matches := getMatchingNumbers(card)
		// For every match, make copies of subsequent cards
		for j := 1; j <= matches; j++ {
			nextCardIdx := card.index + j
			if nextCardIdx <= len(cards) {
				cardCounts[nextCardIdx] += cardCounts[card.index]
			}
		}
	}

	totalCards := 0
	for _, count := range cardCounts {
		totalCards += count
	}
	fmt.Println(totalCards)
}

type scratchCard struct {
	index int
	wins  map[int]struct{}
	have  []int
}

func readCards(filename string) []scratchCard {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var cards []scratchCard
	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		index++
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		winStrings := strings.Fields(parts[0])
		haveStrings := strings.Fields(parts[1])
		wins := make(map[int]struct{}, len(winStrings))
		have := make([]int, len(haveStrings))

		for _, w := range winStrings {
			win, _ := strconv.Atoi(w)
			wins[win] = struct{}{}
		}

		for i, h := range haveStrings {
			have[i], _ = strconv.Atoi(h)
		}

		cards = append(cards, scratchCard{
			index: index,
			wins:  wins,
			have:  have,
		})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return cards
}

func getMatchingNumbers(card scratchCard) int {
	matches := 0
	for _, n := range card.have {
		if _, exists := card.wins[n]; exists {
			matches++
		}
	}
	return matches
}
