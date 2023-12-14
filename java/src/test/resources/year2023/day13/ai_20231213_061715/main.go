
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var patterns [][]string
	currentPattern := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" && len(currentPattern) > 0 {
			patterns = append(patterns, currentPattern)
			currentPattern = []string{}
		} else {
			currentPattern = append(currentPattern, line)
		}
	}
	if len(currentPattern) > 0 {
		patterns = append(patterns, currentPattern)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}

	summary := 0
	for _, pattern := range patterns {
		hReflect, hIndex := findHorizontalReflection(pattern)
		vReflect, vIndex := findVerticalReflection(pattern)

		if vReflect {
			summary += vIndex
		}

		if hReflect {
			summary += 100 * hIndex
		}
	}

	fmt.Println(summary)
}

func findVerticalReflection(pattern []string) (bool, int) {
	cols := len(pattern[0])
	for c := 0; c < cols/2; c++ {
		columnMatches := true
		for r := range pattern {
			if pattern[r][c] != pattern[r][cols-c-1] {
				columnMatches = false
				break
			}
		}
		if columnMatches {
			return true, c
		}
	}
	return false, 0
}

func findHorizontalReflection(pattern []string) (bool, int) {
	rows := len(pattern)
	for r := 0; r < rows/2; r++ {
		rowMatches := true
		for col := range pattern[r] {
			if pattern[r][col] != pattern[rows-r-1][col] {
				rowMatches = false
				break
			}
		}
		if rowMatches {
			return true, r
		}
	}
	return false, 0
}
