
package main

import (
	"bufio"
	"fmt"
	"os"
)

const spinCycleCount = 1000000000

type Point struct {
	x, y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	// Solve Part 1
	part1Load := calculateLoad(grid, false)
	fmt.Println(part1Load)

	// Solve Part 2
	part2Load := calculateLoad(grid, true)
	fmt.Println(part2Load)
}

func calculateLoad(grid [][]rune, performSpinCycles bool) int {
	if performSpinCycles {
		// Run 1000000000 cycles.
		// After each cycle, the position of the rocks will be the same as after the first cycle.
		// So we just need to calculate the north load after the first cycle and multiply it by 1000000000.
		// Perform one spin cycle (north, west, south, east)
		tiltPlatform(grid, 'N')
		tiltPlatform(grid, 'W')
		tiltPlatform(grid, 'S')
		tiltPlatform(grid, 'E')
	}

	// Calculate load
	load := 0
	for x, row := range grid {
		for _, cell := range row {
			if cell == 'O' {
				load += len(grid) - x
			}
		}
	}

	if performSpinCycles {
		// Multiply by spinCycleCount (1000000000) since the grid is in the same state after each full cycle
		load *= spinCycleCount
	}

	return load
}

func tiltPlatform(grid [][]rune, direction rune) {
	deltas := map[rune]Point{
		'N': {0, -1},
		'W': {-1, 0},
		'S': {0, 1},
		'E': {1, 0},
	}

	delta := deltas[direction]
	width, height := len(grid[0]), len(grid)

	// Transform grid according to tilt direction
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch direction {
			case 'N':
				x = width - 1 - x // Reverse direction for x-axis
			case 'W':
				y = height - 1 - y // Reverse direction for y-axis
			}

			if grid[y][x] == 'O' {
				curX, curY := x, y
				for {
					nextX, nextY := curX+delta.x, curY+delta.y
					if nextX < 0 || nextY < 0 || nextX >= width || nextY >= height || grid[nextY][nextX] != '.' {
						break
					}
					grid[nextY][nextX], grid[curY][curX] = grid[curY][curX], grid[nextY][nextX]
					curX, curY = nextX, nextY
				}
			}

			switch direction {
			case 'N':
				x = width - 1 - x // Reverse direction for x-axis
			case 'W':
				y = height - 1 - y // Reverse direction for y-axis
			}
		}
	}
}
