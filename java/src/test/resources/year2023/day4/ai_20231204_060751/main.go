
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cards, err := readCards("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	accumulated := processCards(cards)
	fmt.Println(cumulativeCount(accumulated))
}

func readCards(filename string) ([][2][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cards [][2][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		leftPart := strings.Fields(parts[0])
		rightPart := strings.Fields(parts[1])
		winningNumbers := convertToIntArray(leftPart)
		yourNumbers := convertToIntArray(rightPart)
		cards = append(cards, [2][]int{winningNumbers, yourNumbers})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cards, nil
}

func convertToIntArray(strs []string) []int {
	result := make([]int, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		result[i] = num
	}
	return result
}

func processCards(cards [][2][]int) []int {
	accumulated := make([]int, len(cards))
	copies := make([]int, len(cards))

	for i := range cards {
		copyWinners := checkAndCountWinningNumbers(cards[i][0], cards[i][1])
		accumulated[i]++
		for j := 1; j <= copyWinners; j++ {
			if i+j < len(cards) {
				copies[i+j]++
			}
		}
	}

	for i, count := range copies {
		for count > 0 {
			copyWinners := checkAndCountWinningNumbers(cards[i][0], cards[i][1])
			for j := 1; j <= copyWinners; j++ {
				if i+j < len(cards) {
					copies[i+j]++
				}
			}
			accumulated[i]++
			count--
		}
	}

	return accumulated
}

func checkAndCountWinningNumbers(winningNumbers []int, yourNumbers []int) int {
	winningMap := make(map[int]bool)
	for _, num := range winningNumbers {
		winningMap[num] = true
	}

	matchCount := 0
	for _, num := range yourNumbers {
		if winningMap[num] {
			matchCount++
		}
	}

	return matchCount
}

func cumulativeCount(accumulated []int) int {
	count := 0
	for _, val := range accumulated {
		count += val
	}
	return count
}
