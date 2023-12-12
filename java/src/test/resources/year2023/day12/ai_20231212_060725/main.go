
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Utility to calculate the binomial coefficient (n choose k)
func binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	k = min(k, n-k) // take advantage of symmetry
	c := 1
	for i := 0; i < k; i++ {
		c = c * (n - i) / (i + 1)
	}
	return c
}

// Utility to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Count the arrangements based on the given pattern and group sizes
func countArrangements(pattern string, groups []int) int {
	// Check for base cases where no calculations are needed
	if len(groups) == 0 {
		if strings.Contains(pattern, "?") {
			return 1 << strings.Count(pattern, "?")
		}
		return 1
	}

	// Count the number of ways to distribute '?' among different groups
	var ways, sumGroups int
	for _, group := range groups {
		sumGroups += group
	}

	questionMarks := strings.Count(pattern, "?")
	freeSpaces := questionMarks - (len(groups)-1) - sumGroups

	if freeSpaces < 0 {
		return 0
	}

	ways = binomial(freeSpaces+len(groups), len(groups))

	// Return number of ways to distribute '?' among groups
	return ways
}

func processLine(line string) int {
	parts := strings.Split(line, " ")
	pattern := parts[0]
	groupsString := strings.Split(parts[1], ",")
	groups := make([]int, len(groupsString))

	for i, gs := range groupsString {
		g, err := strconv.Atoi(gs)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return 0
		}
		groups[i] = g
	}

	return countArrangements(pattern, groups)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalArrangements := 0

	for scanner.Scan() {
		line := scanner.Text()
		totalArrangements += processLine(line)
	}

	fmt.Println(totalArrangements)
}
