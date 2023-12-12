
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
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		constraintString := parts[1]
		constraints := parseConstraints(constraintString)
		count := countArrangements(parts[0], constraints)
		total += count
	}
	fmt.Println(total)
}

func parseConstraints(s string) []int {
	constraintStrings := strings.Split(s, ",")
	constraints := make([]int, len(constraintStrings))
	for i, v := range constraintStrings {
		constraints[i], _ = strconv.Atoi(v)
	}
	return constraints
}

func countArrangements(row string, groups []int) int {
	// define a recursive function that tries to match the groups
	var tryMatch func(r []rune, g []int, index int) int
	tryMatch = func(r []rune, g []int, index int) int {
		if index >= len(r) {
			if len(g) == 0 {
				return 1
			}
			return 0
		}
		if len(g) == 0 {
			if strings.ContainsRune(string(r[index:]), '#') {
				return 0
			}
			return 1
		}

		possibilities := 0
		// case for operational (.)
		if r[index] != '#' {
			possibilities += tryMatch(r, g, index+1)
		}

		// case for damaged (#)
		if r[index] == '?' || r[index] == '#' {
			length := len(r)
			if length >= index+g[0] && (length <= index+g[0] || r[index+g[0]] != '#') {
				match := true
				for i := index; i < index+g[0]; i++ {
					if r[i] == '.' {
						match = false
						break
					}
				}
				if match {
					newIndex := index + g[0]
					for newIndex < len(r) && (r[newIndex] != '?' && r[newIndex] != '.') {
						newIndex++
					}
					possibilities += tryMatch(r, g[1:], newIndex+1)
				}
			}
		}
		return possibilities
	}

	runes := []rune(row)
	result := tryMatch(runes, groups, 0)
	return result
}

func main() {
	// Open file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file: ", err)
		return
	}
	defer file.Close()

	totalArrangements := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		springRow, groupsStr := parts[0], parts[1]
		groups := parseConstraintGroups(groupsStr)
		arrangements := countArrangements(springRow, groups)
		totalArrangements += arrangements
	}
	if scanner.Err() != nil {
		fmt.Println("Error reading from input file: ", scanner.Err())
		return
	}

	fmt.Println(totalArrangements)
}

// countArrangements counts the number of possible arrangements given a spring row and constraints
func countArrangements(springRow string, groups []int) int {
	return countArrangementsRecursive(springRow, groups, 0, 0)
}

// countArrangementsRecursive is a helper recursive function for countArrangements
func countArrangementsRecursive(springRow string, groups []int, groupIndex int, position int) int {
	if groupIndex >= len(groups) {
		// If we processed all groups, the remaining '?' must be '.' (operational)
		return 1
	}

	if position >= len(springRow) {
		return 0
	}

	char := springRow[position]
	groupSize := groups[groupIndex]
	count := 0

	if char == '.' || char == '?' {
		// Assume operational and continue
		count += countArrangementsRecursive(springRow, groups, groupIndex, position+1)
	}

	if char == '#' || char == '?' {
		// Assume the group of broken springs starts here if there's enough room
		if len(springRow)-position >= groupSize &&
			(position+groupSize == len(springRow) || springRow[position+groupSize] == '.' || springRow[position+groupSize] == '?') {
			// Continue after the group of broken springs
			count += countArrangementsRecursive(springRow, groups, groupIndex+1, position+groupSize+1)
		}
	}

	return count
}

// parseConstraintGroups parses the string of constraints into a slice of integers
func parseConstraintGroups(groupsStr string) []int {
	groupStr := strings.Split(groupsStr, ",")
	groups := make([]int, len(groupStr))

	for i, str := range groupStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Error parsing group constraints: ", err)
			return nil
		}
		groups[i] = num
	}

	return groups
}
