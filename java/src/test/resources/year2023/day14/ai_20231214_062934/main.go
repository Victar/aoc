
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const cycles = 1000000000
const (
	empty = '.'
	round = 'O'
	cube  = '#'
)

type Point struct{ x, y int }

func main() {
	content := loadInput("input.txt")
	grid := parseToGrid(content)
	height := len(grid)

	load := calculateInitialLoad(grid)
	for cycle := 1; cycle <= cycles; cycle++ {
		grid = performCycle(grid)
	}
	loadAfterCycles := calculateLoadAfterCycles(grid, height)
	fmt.Println(loadAfterCycles)
}

func loadInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}

func parseToGrid(content []string) [][]rune {
	grid := make([][]rune, len(content))
	for i, line := range content {
		grid[i] = []rune(line)
	}
	return grid
}

func calculateInitialLoad(grid [][]rune) int {
	load := 0
	for rowIndex, row := range grid {
		for _, mirror := range row {
			if mirror == round {
				load += (len(grid) - rowIndex)
			}
		}
	}
	return load
}

func calculateLoadAfterCycles(grid [][]rune, height int) int {
	load := 0
	for _, row := range grid {
		for colIndex, mirror := range row {
			if mirror == round {
				load += (height - colIndex)
			}
		}
	}
	return load
}

func performCycle(grid [][]rune) [][]rune {
	grid = slide(grid, 0, -1) // north
	grid = slide(grid, -1, 0) // west
	grid = slide(grid, 0, 1)  // south
	grid = slide(grid, 1, 0)  // east
	return grid
}

func slide(grid [][]rune, dx, dy int) [][]rune {
	height, width := len(grid), len(grid[0])
	moved := true
	for moved {
		moved = false
		newGrid := make([][]rune, height)
		for i := range newGrid {
			newGrid[i] = make([]rune, width)
			copy(newGrid[i], grid[i])
		}
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == round {
					newX, newY := x+dx, y+dy
					if newX >= 0 && newY >= 0 && newX < width && newY < height && grid[newY][newX] == empty {
						newGrid[y][x], newGrid[newY][newX] = newGrid[newY][newX], newGrid[y][x]
						moved = true
					}
				}
			}
		}
		grid = newGrid
	}
	return grid
}
