
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// LoadFromFile reads input from "input.txt" and returns a slice of strings representing each line.
func LoadFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// CalculateLoad calculates the total load on the north support beams.
func CalculateLoad(input []string) int {
	totalLoad := 0
	for rowIndex, line := range input {
		for _, char := range line {
			if char == 'O' {
				totalLoad += len(input) - rowIndex
			}
		}
	}
	return totalLoad
}

func main() {
	input, err := LoadFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	load := CalculateLoad(input)
	fmt.Print(load)
}
