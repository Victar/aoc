
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
	totalCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		springConditions := parts[0]
		numbers := strings.Split(parts[1], ",")
		groupSizes := make([]int, len(numbers))
		for i, n := range numbers {
			groupSizes[i], _ = strconv.Atoi(n)
		}

		count := countArrangements(springConditions, groupSizes, 0, 0)
		totalCount += count
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalCount)
}

func countArrangements(spring string, groups []int, pos int, groupIndex int) int {
	if groupIndex >= len(groups) {
		for i := pos; i < len(spring); i++ {
			if spring[i] == '?' {
				return 0 // Any remaining '?' must be invalid now
			}
		}
		return 1
	}

	if pos >= len(spring) {
		return 0
	}

	total := 0

	if spring[pos] == '?' {
		// Try with broken spring
		total += countArrangements(setCharAt(spring, pos, '#'), groups, pos+1, groupIndex)

		// Move to next group or end if no more groups
		nextGroupIndex := groupIndex
		if groups[groupIndex] == 1 {
			nextGroupIndex++
		} else {
			groupsCopy := make([]int, len(groups))
			copy(groupsCopy, groups)
			groupsCopy[groupIndex]--
			return total + countArrangements(spring, groupsCopy, pos+1, groupIndex)
		}

		// Try with operational spring, skip to the next group if current group is single broken spring
		return total + countArrangements(setCharAt(spring, pos, '.'), groups, pos+1, nextGroupIndex)
	}

	if spring[pos] == '#' {
		if groups[groupIndex] == 1 {
			return countArrangements(spring, groups, pos+1, groupIndex+1)
		}
		groupsCopy := make([]int, len(groups))
		copy(groupsCopy, groups)
		groupsCopy[groupIndex]--
		return countArrangements(spring, groupsCopy, pos+1, groupIndex)
	}

	// If the spring at pos is operational and there's a group of broken springs pending
	if groups[groupIndex] > 0 {
		return countArrangements(spring, groups, pos+1, groupIndex)
	}

	return total
}

func setCharAt(str string, index int, char rune) string {
	out := []rune(str)
	out[index] = char
	return string(out)
}
