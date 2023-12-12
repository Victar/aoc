
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countArrangements(pattern string, groups []int) int {
	dp := make([]int, len(pattern)+1)
	dp[len(pattern)] = 1

	// Calculate a cache of skips to make pattern traversal faster.
	skipCache := make([]int, len(pattern))
	j := 0
	for i := range skipCache {
		skipCache[i] = j
		if pattern[i] != '#' {
			j++
		}
	}

	// Recursive function to backtrack and count configurations.
	var backtrack func(pos, groupIndex int) int
	backtrack = func(pos, groupIndex int) int {
		if groupIndex == len(groups) {
			return 1
		}
		if skipCache[pos] < groups[groupIndex] {
			return 0
		}

		if dp[pos] != 0 {
			return dp[pos]
		}

		count := 0

		// Try to place a group at the current position.
		if isPlaceable(pattern, pos, groups[groupIndex]) {
			count += backtrack(pos+groups[groupIndex]+1, groupIndex+1)
		}

		// Skip a safe location and continue.
		if pattern[pos] != '#' {
			count += backtrack(pos+1, groupIndex)
		}

		dp[pos] = count
		return count
	}

	return backtrack(0, 0)
}

func isPlaceable(pattern string, pos, size int) bool {
	if pos+size > len(pattern) {
		return false
	}
	for i := pos; i < pos+size; i++ {
		if pattern[i] == '.' {
			return false
		}
	}
	if pos+size < len(pattern) && pattern[pos+size] == '#' {
		return false
	}
	if pos > 0 && pattern[pos-1] == '#' {
		return false
	}
	return true
}

func parseInput(filename string) ([]string, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var patterns []string
	var groups [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		patterns = append(patterns, strings.ReplaceAll(parts[0], "?", "#"))
		group := make([]int, 0)
		nums := strings.Split(parts[1], ",")
		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			group = append(group, n)
		}
		groups = append(groups, group)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return patterns, groups
}

func main() {
	patterns, groups := parseInput("input.txt")

	total := 0
	for i := range patterns {
		total += countArrangements(patterns[i], groups[i])
	}
	fmt.Println(total)
}
