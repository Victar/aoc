
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to read the input file and return a slice of strings where each string is a line in the file.
func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Function to solve the arrangement counting problem for a single line of the puzzle input.
func countArrangements(springRow string) int {
	parts := strings.Split(springRow, " ")
	rowData, groupData := parts[0], parts[1]
	groups := strings.Split(groupData, ",")
	intGroups := make([]int, 0, len(groups))

	// Converting the group sizes from string slice to int slice
	for _, grp := range groups {
		size, _ := strconv.Atoi(grp)
		intGroups = append(intGroups, size)
	}

	// Dynamic programming approach to count the arrangements
	memo := make(map[string]int)
	return count(rowData, intGroups, 0, memo)
}

// Recursive function with memoization to determine the number of valid arrangements.
func count(rowData string, groups []int, i int, memo map[string]int) int {
	if i == len(groups) {
		return 1
	}

	// Generate a memoization key
	key := fmt.Sprintf("%s|%d", rowData, i)
	if val, exists := memo[key]; exists {
		return val
	}

	count := 0
	groupSize := groups[i]
	pattern := "#" + strings.Repeat(".", groupSize-1)
	for start := strings.Index(rowData, "?"); start != -1 && start+groupSize <= len(rowData); start = strings.Index(rowData, "?") {
		newRow := rowData[:start] + pattern + rowData[start+groupSize:]
		if start+groupSize < len(rowData) && newRow[start+groupSize] == '#' {
			continue
		}
		if start > 0 && newRow[start-1] == '#' {
			continue
		}
		count += count(newRow, groups, i+1, memo)
	}
	memo[key] = count
	return count
}

func main() {
	inputLines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	totalArrangements := 0
	for _, line := range inputLines {
		totalArrangements += countArrangements(line)
	}

	fmt.Println(totalArrangements)
}
