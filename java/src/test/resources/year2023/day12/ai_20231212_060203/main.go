
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]string, error) {
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

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func parseLine(line string) ([]rune, []int) {
	parts := strings.Split(line, " ")
	springState := []rune(parts[0])

	groupSizesStr := strings.Split(parts[1], ",")
	groupSizes := make([]int, len(groupSizesStr))
	for i, str := range groupSizesStr {
		groupSize, _ := strconv.Atoi(str)
		groupSizes[i] = groupSize
	}

	return springState, groupSizes
}

func countArrangements(springState []rune, groupSizes []int) int {
	// Recursive function to count the arrangements
	var count func(int, int) int
	count = func(idx, groupIdx int) int {
		if groupIdx == len(groupSizes) {
			for ; idx < len(springState); idx++ { 
				if springState[idx] == '#' { 
					return 0
				}
			}
			return 1
		}

		total := 0
		counter := groupSizes[groupIdx]
		for i := idx; i < len(springState); i++ {
			if springState[i] == '.' {
				continue
			}
			if i+counter > len(springState) {
				break
			}
			if springState[i+counter] == '#' {
				continue
			}
			valid := true
			for j := 0; j < counter; j++ {
				if springState[i+j] == '.' {
					valid = false
					break
				}
			}
			if valid {
				old := make([]rune, counter)
				copy(old, springState[i:i+counter])
				for j := 0; j < counter; j++ { 
					springState[i+j] = '#' 
				}
				if i+counter < len(springState) {
					springState[i+counter] = '.' 
				}
				total += count(i+counter+1, groupIdx+1)
				for j := 0; j < counter; j++ { 
					springState[i+j] = old[j]
				}
			}
		}
		return total
	}

	return count(0, 0)
}

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	sum := 0
	for _, line := range lines {
		springState, groupSizes := parseLine(line)
		sum += countArrangements(springState, groupSizes)
	}

	fmt.Println(sum)
}
