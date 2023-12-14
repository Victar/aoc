package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findReflections(input []string) int {
	count := 0

	for _, line := range input {
		pattern := strings.TrimSpace(line)
		// Check vertical reflection
		for i := 0; i < len(pattern)/2; i++ {
			if pattern[i] == pattern[len(pattern)-1-i] {
				count += len(pattern) - len(pattern)/2
				break
			}
		}
		// Check horizontal reflection
		for i := 0; i < len(input)/2; i++ {
			if input[i] == input[len(input)-1-i] {
				count += 100 * (len(input) - len(input)/2)
				break
			}
		}
	}

	return count
}

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day13/sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	result := findReflections(lines)
	fmt.Println(result)
}
