
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cnt := simulateLight(grid)
	fmt.Println(cnt)
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func simulateLight(grid [][]rune) int {
	energized := make(map[[2]int]bool)
	x, y := 0, 0
	dx, dy := 1, 0 // Beam starts going to the right

	for {
		// If out of bounds, stop simulating
		if y >= len(grid) || y < 0 || x >= len(grid[y]) || x < 0 {
			break
		}
		energized[[2]int{x, y}] = true

		switch grid[y][x] {
		case '/':
			dx, dy = -dy, -dx
		case '\\':
			dx, dy = dy, dx
		case '|':
			if dy != 0 {
				// Split beam
				energized[[2]int{x + 1, y}] = true
				energized[[2]int{x - 1, y}] = true
			}
		case '-':
			if dx != 0 {
				// Split beam
				energized[[2]int{x, y + 1}] = true
				energized[[2]int{x, y - 1}] = true
			}
		}

		// Move the beam
		x += dx
		y += dy
	}

	// Count energized tiles
	count := 0
	for _, value := range energized {
		if value {
			count++
		}
	}

	return count
}
