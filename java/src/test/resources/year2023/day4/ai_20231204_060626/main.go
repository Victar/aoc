
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var scratchcards [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " | ")
		if len(parts) != 2 {
			continue
		}
		cardNumbers := strings.Fields(parts[0])
		yourNumbers := strings.Fields(parts[1])
		numbers := make([]int, len(cardNumbers)+len(yourNumbers))
		for i, n := range append(cardNumbers, yourNumbers...) {
			numbers[i], _ = strconv.Atoi(n)
		}
		scratchcards = append(scratchcards, numbers)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	totalCardCount := processScratchcards(scratchcards)
	fmt.Println(totalCardCount)
}

func processScratchcards(scratchcards [][]int) int {
	cardCount := len(scratchcards)
	totalCards := make([]int, cardCount)

	for i := 0; i < cardCount; i++ {
		totalCards[i] = 1
	}

	for i := 0; i < cardCount; i++ {
		winningSet := make(map[int]struct{})
		for _, num := range scratchcards[i][:len(scratchcards[i])/2] {
			winningSet[num] = struct{}{}
		}

		matchCount := 0
		for _, num := range scratchcards[i][len(scratchcards[i])/2:] {
			if _, found := winningSet[num]; found {
				matchCount++
			}
		}

		additionalCards := matchCount * totalCards[i]
		for j := i + 1; j < i+1+matchCount && j < cardCount; j++ {
			totalCards[j] += additionalCards
		}
	}

	total := 0
	for i := 0; i < cardCount; i++ {
		total += totalCards[i]
	}
	return total
}
