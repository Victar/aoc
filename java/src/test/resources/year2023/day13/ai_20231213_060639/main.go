package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day13/sample.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := [][]string{}

	// Read input patterns
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		patterns = append(patterns, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	fmt.Println(findReflectionNumber(patterns))
}

func findReflectionNumber(patterns [][]string) int {
	summary := 0

	for _, pattern := range patterns {
		rows := len(pattern)
		cols := len(pattern[0])

		// Check for vertical reflection line
		for col := 0; col < cols/2; col++ {
			mirror := true
			for row := 0; row < rows; row++ {
				if pattern[row][col] != pattern[row][cols-1-col] {
					mirror = false
					break
				}
			}
			if mirror {
				summary += col
				break
			}
		}

		// Check for horizontal reflection line
		for row := 0; row < rows/2; row++ {
			mirror := true
			for col := 0; col < cols; col++ {
				if pattern[row][col] != pattern[rows-1-row][col] {
					mirror = false
					break
				}
			}
			if mirror {
				summary += 100 * row
				break
			}
		}
	}
	return summary
}
