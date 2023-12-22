
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Empty    = '.'
	MirrorF  = '/'
	MirrorB  = '\\'
	Splitter = '|'
	SplitterH = '-'
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	count := simulateBeam(grid)
	fmt.Println(count)
}

func simulateBeam(grid [][]rune) int {
	dx, dy := 1, 0 // initial direction of the beam: rightward

	// Add an additional row and column to the grid to handle beam splits at the edges.
	extendedGrid := make([][]bool, len(grid)+1)
	for i := range extendedGrid {
		extendedGrid[i] = make([]bool, len(grid[0])+1)
	}

	var energize func(x, y, dx, dy int)
	energize = func(x, y, dx, dy int) {
		for x >= 0 && y >= 0 && y < len(grid) && x < len(grid[y]) {
			tile := grid[y][x]
			extendedGrid[y][x] = true

			switch tile {
			case MirrorF:
				dx, dy = dy, dx // reflect beam
			case MirrorB:
				dx, dy = -dy, -dx // reflect beam
			case Splitter:
				// Beam hits the flat end of a splitter and splits
				if dy == 0 {
					// Horizontal splitter
					energize(x, y-1, 0, -1)
					energize(x, y+1, 0, 1)
					return
				} else {
					// Vertical splitter
					energize(x-1, y, -1, 0)
					energize(x+1, y, 1, 0)
					return
				}
			case SplitterH:
				// Beam hits the flat end of a splitter and splits
				if dx == 0 {
					// Vertical splitter
					energize(x-1, y, -1, 0)
					energize(x+1, y, 1, 0)
					return
				} else {
					// Horizontal splitter
					energize(x, y-1, 0, -1)
					energize(x, y+1, 0, 1)
					return
				}
			}

			x += dx
			y += dy
		}
	}

	// Start the simulation from the top left corner, going right.
	energize(0, 0, dx, dy)

	// Count energized tiles
	count := 0
	for y := range grid {
		for x := range grid[y] {
			if extendedGrid[y][x] {
				count++
			}
		}
	}

	return count
}
