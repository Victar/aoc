package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day14/sample.txt")
	grid := parseInput(lines)
	const cycles = 1000000000

	grid, _ = simulateCycles(grid, cycles)
	load := calculateLoad(grid)
	fmt.Println(load)
}

// readInput reads the input file and returns a slice of lines.
func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

// parseInput converts input text lines into a grid representation.
func parseInput(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// simulateCycles performs the specified number of tilt cycles on the grid.
func simulateCycles(grid [][]rune, cycles int) ([][]rune, int) {
	maxRepeatedCycles := findCyclePattern(grid)
	reducedCycles := cycles % maxRepeatedCycles

	for cycle := 0; cycle < reducedCycles; cycle++ {
		grid = tilt(grid, 'N')
		grid = tilt(grid, 'W')
		grid = tilt(grid, 'S')
		grid = tilt(grid, 'E')
	}

	return grid, reducedCycles
}

// findCyclePattern finds the cycle after which the grid state repeats.
func findCyclePattern(grid [][]rune) int {
	seenStates := make(map[string]int)
	cycle := 0
	for {
		gridString := gridToString(grid)
		if prevCycle, exists := seenStates[gridString]; exists {
			return cycle - prevCycle
		}
		seenStates[gridString] = cycle
		grid = tilt(grid, 'N')
		grid = tilt(grid, 'W')
		grid = tilt(grid, 'S')
		grid = tilt(grid, 'E')
		cycle++
	}
}

// gridToString converts the grid to a string representation for comparison.
func gridToString(grid [][]rune) string {
	var result string
	for _, row := range grid {
		result += string(row)
	}
	return result
}

// tilt handles the grid tilting logic in a specified direction.
func tilt(grid [][]rune, direction rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	switch direction {
	case 'N':
		for x := 0; x < len(grid[0]); x++ {
			for y := 1; y < len(grid); y++ {
				if newGrid[y][x] == 'O' && newGrid[y-1][x] == '.' {
					newGrid[y][x], newGrid[y-1][x] = newGrid[y-1][x], newGrid[y][x]
					y = 0 // Start checking from top again after movement
				}
			}
		}
	case 'W':
		for y := 0; y < len(grid); y++ {
			for x := 1; x < len(grid[y]); x++ {
				if newGrid[y][x] == 'O' && newGrid[y][x-1] == '.' {
					newGrid[y][x], newGrid[y][x-1] = newGrid[y][x-1], newGrid[y][x]
					x = 0 // Start checking from left again after movement
				}
			}
		}
	case 'S':
		for x := 0; x < len(grid[0]); x++ {
			for y := len(grid) - 2; y >= 0; y-- {
				if newGrid[y][x] == 'O' && newGrid[y+1][x] == '.' {
					newGrid[y][x], newGrid[y+1][x] = newGrid[y+1][x], newGrid[y][x]
					y = len(grid) // Start from bottom after movement
				}
			}
		}
	case 'E':
		for y := 0; y < len(grid); y++ {
			for x := len(grid[y]) - 2; x >= 0; x-- {
				if newGrid[y][x] == 'O' && newGrid[y][x+1] == '.' {
					newGrid[y][x], newGrid[y][x+1] = newGrid[y][x+1], newGrid[y][x]
					x = len(grid[y]) // Start from right after movement
				}
			}
		}
	}

	return newGrid
}

// calculateLoad calculates the total load on the north support beams.
func calculateLoad(grid [][]rune) int {
	load := 0
	height := len(grid)
	for y, row := range grid {
		for _, cell := range row {
			if cell == 'O' {
				load += height - y
			}
		}
	}
	return load
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
