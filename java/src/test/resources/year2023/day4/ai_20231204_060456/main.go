
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to calculate the number of matching numbers on a scratchcard and the resulting score.
func calculateScore(winningNumbers, cardNumbers []int) int {
    matchCount := 0
    for _, cn := range cardNumbers {
        for _, wn := range winningNumbers {
            if cn == wn {
                matchCount++
                break
            }
        }
    }
    if matchCount == 0 {
        return 0
    }
    // Calculate the score based on the number of matches.
    score := 1
    for i := 1; i < matchCount; i++ {
        score *= 2
    }
    return score
}

func main() {
    inputFile, err := os.Open("input.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
        os.Exit(1)
    }
    defer inputFile.Close()

    scanner := bufio.NewScanner(inputFile)
    totalPoints := 0
    for scanner.Scan() {
        // Extract winning and card numbers as slices of integers.
        line := scanner.Text()
        parts := strings.Split(line, " | ")
        winningStr := strings.Fields(parts[0])
        cardStr := strings.Fields(parts[1])

        winningNumbers := make([]int, len(winningStr))
        cardNumbers := make([]int, len(cardStr))

        for i, s := range winningStr {
            number, _ := strconv.Atoi(s)
            winningNumbers[i] = number
        }
        for i, s := range cardStr {
            number, _ := strconv.Atoi(s)
            cardNumbers[i] = number
        }

        // Calculate the score for the card and add it to the total points.
        score := calculateScore(winningNumbers, cardNumbers)
        totalPoints += score
    }

    // Check for scanning errors.
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
        os.Exit(1)
    }

    // Print the total points as required.
    fmt.Println(totalPoints)
}
