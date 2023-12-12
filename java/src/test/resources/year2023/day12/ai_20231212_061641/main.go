
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
	totalArrangements := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		pattern := split[0]
		groups := parseGroups(split[1])
		count := countArrangements(pattern, groups)
		totalArrangements += count
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalArrangements)
}

func parseGroups(groupsStr string) []int {
	groups := strings.Split(groupsStr, ",")
	var nums []int
	for _, g := range groups {
		n, _ := strconv.Atoi(g)
		nums = append(nums, n)
	}
	return nums
}

func countArrangements(pattern string, groups []int) int {
	return countArrangementsRecursive(pattern, 0, groups, 0)
}

func countArrangementsRecursive(pattern string, patternIndex int, groups []int, groupIndex int) int {
	if groupIndex == len(groups) {
		for i := patternIndex; i < len(pattern); i++ {
			if pattern[i] == '?' {
				return 0
			}
		}
		return 1
	}

	if groups[groupIndex] == 0 {
		return countArrangementsRecursive(pattern, patternIndex, groups, groupIndex+1)
	}

	count := 0
	consecutiveHashes := 0
	for i := patternIndex; i < len(pattern); i++ {
		if pattern[i] == '#' || pattern[i] == '?' {
			consecutiveHashes++
			if consecutiveHashes == groups[groupIndex] {
				if i == len(pattern)-1 || pattern[i+1] == '.' {
					count += countArrangementsRecursive(pattern, i+1, groups, groupIndex+1)
				}
			}
		} else if pattern[i] == '.' {
			consecutiveHashes = 0
		}

		if consecutiveHashes > groups[groupIndex] {
			break
		}
	}

	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		pattern := split[0]
		groups := parseGroups(split[1])
		sum += countArrangements(pattern, groups)
	}

	fmt.Println(sum)
}
