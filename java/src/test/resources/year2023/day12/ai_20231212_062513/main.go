
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

	scanner := bufio.NewScanner(file)
	var arrangementCounts int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		pattern := parts[0]
		groups := parseGroups(parts[1])
		arrangementCounts += countArrangements(pattern, groups)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(arrangementCounts)
}

func parseGroups(groupStr string) []int {
	groupStrs := strings.Split(groupStr, ",")
	groups := make([]int, len(groupStrs))
	for i, gs := range groupStrs {
		num, _ := strconv.Atoi(gs)
		groups[i] = num
	}
	return groups
}

func countArrangements(pattern string, groups []int) int {
	return countRecursive(pattern, groups, 0, 0)
}

func countRecursive(pattern string, groups []int, groupIdx int, startIdx int) int {
	if groupIdx >= len(groups) {
		for i := startIdx; i < len(pattern); i++ {
			if pattern[i] == '?' {
				return 0
			}
		}
		return 1
	}

	groupSize := groups[groupIdx]
	count := 0
	for i := startIdx; i <= len(pattern)-groupSize; i++ {
		if canPlaceGroup(pattern, i, groupSize) {
			newPattern := replacePattern(pattern, i, groupSize)
			nextStartIdx := i + groupSize + 1
			count += countRecursive(newPattern, groups, groupIdx+1, nextStartIdx)
		}
	}

	return count
}

func canPlaceGroup(pattern string, startIdx int, groupSize int) bool {
	if startIdx != 0 && pattern[startIdx-1] == '#' {
		return false
	}
	for i := startIdx; i < startIdx+groupSize; i++ {
		if pattern[i] == '.' {
			return false
		}
	}
	if startIdx+groupSize < len(pattern) && pattern[startIdx+groupSize] == '#' {
		return false
	}
	return true
}

func replacePattern(pattern string, startIdx int, groupSize int) string {
	patternBytes := []byte(pattern)
	for i := startIdx; i < startIdx+groupSize; i++ {
		patternBytes[i] = '#'
	}
	return string(patternBytes)
}

