package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var patterns [][]string
	currentPattern := []string{}

	// Read the input file and group patterns
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentPattern) > 0 {
				patterns = append(patterns, currentPattern)
				currentPattern = []string{}
			}
		} else {
			currentPattern = append(currentPattern, line)
		}
	}
	// Add the last pattern if exists
	if len(currentPattern) > 0 {
		patterns = append(patterns, currentPattern)
	}

	// Process each pattern
	sum := 0
	for _, pattern := range patterns {
		sum += processPattern(pattern)
	}

	// Print the final result
	fmt.Println(sum)
}

// processPattern finds a line of reflection and returns the sum according to the instructions
func processPattern(pattern []string) int {
	verticalSum := tryVerticalReflection(pattern)
	if verticalSum >= 0 {
		return verticalSum
	}
	horizontalSum := tryHorizontalReflection(pattern)
	if horizontalSum >= 0 {
		return horizontalSum * 100
	}

	return 0
}

// tryVerticalReflection tries to find a vertical line of reflection
func tryVerticalReflection(pattern []string) int {
	n := len(pattern[0])
	for i := 1; i < n; i++ {
		if isVerticalReflection(pattern, i) {
			return i
		}
	}
	return -1
}

// isVerticalReflection checks if there is a vertical reflection at the given column index
func isVerticalReflection(pattern []string, index int) bool {
	for _, row := range pattern {
		left := reverseString(row[:index])
		right := row[index:]

		// Mirror the longer side if uneven
		if len(left) > len(right) {
			left = left[:len(right)]
		} else {
			right = right[:len(left)]
		}

		if left != right {
			return false
		}
	}
	return true
}

// reverseString reverses a string
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// tryHorizontalReflection tries to find a horizontal line of reflection
func tryHorizontalReflection(pattern []string) int {
	n := len(pattern)
	for i := 1; i < n; i++ {
		if isHorizontalReflection(pattern, i) {
			return i
		}
	}
	return -1
}

// isHorizontalReflection checks if there is a horizontal reflection at the given row index
func isHorizontalReflection(pattern []string, index int) bool {
	for i := 0; i < index; i++ {
		if i >= len(pattern)-index {
			break
		}
		if pattern[i] != pattern[len(pattern)-i-1] {
			return false
		}
	}
	return true
}

func init() {
	fmt.Println(strings.Repeat("=", 80))
}

// summarizePattern analyzes a pattern to find its reflection
// and returns the summary value for that pattern according to the rules.
func summarizePattern(pattern []string) int {
	// Check for reflection column-wise
	for col := 1; col < len(pattern[0]); col++ {
		mirror := true
		for _, row := range pattern {
			left, right := row[:col], reverse(row[col:])
			if !strings.HasPrefix(right, left) {
				mirror = false
				break
			}
		}
		if mirror {
			return col
		}
	}

	// Check for reflection row-wise
	for row := 1; row < len(pattern); row++ {
		top, bottom := pattern[:row], reversePattern(pattern[row:])
		if reflectPatterns(top, bottom) {
			return 100 * row
		}
	}

	return 0
}

// reverse a string
func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// reversePattern reverses a slice of strings
func reversePattern(pattern []string) []string {
	reversed := make([]string, len(pattern))
	for i := range pattern {
		reversed[len(reversed)-1-i] = pattern[i]
	}
	return reversed
}

// reflectPatterns checks if two patterns reflect each other.
func reflectPatterns(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
