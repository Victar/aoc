
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countArrangements(line string) int {
	parts := strings.Split(line, " ")
	unknownSeq := parts[0]
	groupStrs := strings.Split(parts[1], ",")
	groups := make([]int, len(groupStrs))

	for i, gs := range groupStrs {
		num, _ := strconv.Atoi(gs)
		groups[i] = num
	}

	return countHelper(unknownSeq, groups, 0)
}

func countHelper(seq string, groups []int, currentGroup int) int {
	if len(seq) == 0 {
		return 1
	}

	if seq[0] == '#' || seq[0] == '.' {
		if seq[0] == '#' {
			if groups[currentGroup] > 0 {
				groups[currentGroup]--
				ret := countHelper(seq[1:], groups, currentGroup)
				groups[currentGroup]++
				return ret
			}
		} else {
			if currentGroup == 0 || groups[currentGroup-1] == 0 {
				return countHelper(seq[1:], groups, currentGroup)
			}
		}
		return 0
	}

	// seq[0] is '?'
	if currentGroup == 0 || groups[currentGroup-1] == 0 {
		count := 0
		// Try broken spring
		if currentGroup < len(groups) {
			groups[currentGroup]--
			count += countHelper(seq[1:], groups, currentGroup)
			groups[currentGroup]++
		}
		// Try operational spring
		count += countHelper(seq[1:], groups, currentGroup)
		return count
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	totalArrangements := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		totalArrangements += countArrangements(line)
	}

	fmt.Println(totalArrangements)
}
