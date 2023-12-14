
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findReflection(line []rune, isVertical bool) int {
	length := len(line)
	for i := 0; i < length/2; i++ {
		if line[i] != line[length-1-i] {
			return i
		}
	}
	// If no mismatches are found, it's a middle reflection.
	if isVertical {
		return length / 2
	}
	return length/2 - 1
}

func isVerticalReflection(pattern []string) bool {
	for i := range pattern[0] {
		for j := 1; j < len(pattern); j++ {
			if pattern[j][i] != pattern[len(pattern)-1-j][i] {
				return false
			}
		}
	}
	return true
}

func calculateScore(pattern []string) int {
	if isVerticalReflection(pattern) {
		// Find column to reflect
		verticalLine := findReflection([]rune(pattern[0]), true)
		return verticalLine
	} else {
		// Find row to reflect
		horizontalLine := findReflection([]rune(pattern), false)
		return 100 * horizontalLine
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		} else {
			score += calculateScore(lines)
			lines = []string{}
		}
	}
	// Calculate score for last pattern if the file doesn't end with a blank line
	if len(lines) > 0 {
		score += calculateScore(lines)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(score)
}
