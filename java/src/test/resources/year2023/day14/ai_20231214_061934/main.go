package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RollRocksNorth updates the grid state by rolling all rounded rocks north.
func RollRocksNorth(grid []string) []string {
	// Convert the grid of strings to a grid of runes for easier manipulation
	runesGrid := make([][]rune, len(grid))
	for i, row := range grid {
		runesGrid[i] = []rune(row)
	}

	height := len(runesGrid)
	width := len(runesGrid[0])

	// Roll each rounded rock north
	for col := 0; col < width; col++ {
		for row := 1; row < height; row++ {
			currentCell := runesGrid[row][col]
			if currentCell == 'O' {
				targetRow := row
				// Find the highest position the rock can roll to
				for targetRow-1 >= 0 && runesGrid[targetRow-1][col] == '.' {
					targetRow--
				}
				if targetRow != row {
					// Move the rock to the new position
					runesGrid[row][col] = '.'
					runesGrid[targetRow][col] = 'O'
				}
			}
		}
	}

	// Convert the grid of runes back to a grid of strings
	newGrid := make([]string, height)
	for i, row := range runesGrid {
		newGrid[i] = string(row)
	}

	return newGrid
}

// CalculateLoad calculates the total load on the north support beams.
func CalculateLoad(grid []string) int {
	load := 0
	height := len(grid)

	for row, line := range grid {
		for _, cell := range line {
			if cell == 'O' {
				// Load is the distance from the bottom plus one (the row the rock is on)
				load += (height - row)
			}
		}
	}

	return load
}

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day14/sample.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid []string

	// Read grid from input.txt
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input.txt:", err)
		return
	}

	// Roll all rounded rocks north
	grid = RollRocksNorth(grid)
	// Calculate the total load on the north support beams
	totalLoad := CalculateLoad(grid)

	fmt.Println(totalLoad)
}
