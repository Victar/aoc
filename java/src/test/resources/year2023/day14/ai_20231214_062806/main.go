
package main

import (
	"bufio"
	"fmt"
	"os"
)

const cycles = 1000000000

// Delta positions (north, west, south, east)
var deltas = []struct{ dx, dy int }{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	load := calculateLoad(grid)
	fmt.Println("Load after tilting north:", load)

	for cycle := 0; cycle < cycles%4; cycle++ {
		grid = tiltGrid(grid)
	}

	load = calculateLoad(grid)
	fmt.Println("Load after", cycles, "cycles:", load)
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
		row := []rune(scanner.Text())
		grid = append(grid, row)
	}

	return grid, scanner.Err()
}

func calculateLoad(grid [][]rune) (load int) {
	for y := len(grid) - 1; y >= 0; y-- {
		for x, cell := range grid[y] {
			if cell == 'O' {
				load += y + 1
			}
		}
	}
	return
}

func tiltGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	for _, d := range deltas {
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] == 'O' {
					nextX, nextY := x+d.dx, y+d.dy
					for inBounds(nextX, nextY, grid) && grid[nextY][nextX] == '.' {
						grid[nextY][nextX] = 'O'
						grid[y][x] = '.'
						y, x = nextY, nextX
						nextX, nextY = x+d.dx, y+d.dy
					}
				}
			}
		}
	}

	return newGrid
}

func inBounds(x, y int, grid [][]rune) bool {
	return x >= 0 && y >= 0 && x < len(grid[0]) && y < len(grid)
}

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Part 1 - Tilt north and compute the load.
	finalGridNorth := tiltGrid(grid, 0)
	loadPart1 := calculateLoad(finalGridNorth)
	fmt.Println(loadPart1)
}

// readInput reads the contents of the given input filename and returns a grid representation.
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

	return grid, scanner.Err()
}

// tiltGrid tilts the given grid in the direction specified by delta index (0 = north, 1 = west, 2 = south, 3 = east).
func tiltGrid(grid [][]rune, deltaIdx int) [][]rune {
	newGrid := make([][]rune, len(grid))
	for i := range grid {
		newGrid[i] = make([]rune, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	delta := deltas[deltaIdx]
	check := func(x, y int) bool {
		if deltaIdx == 0 {
			return y >= 0
		} else if deltaIdx == 1 {
			return x >= 0
		} else if deltaIdx == 2 {
			return y < len(grid)
		} else {
			return x < len(grid[y])
		}
	}

	var move func(int, int)
	move = func(x, y int) {
		nextX, nextY := x+delta.dx, y+delta.dy
		if check(nextX, nextY) && newGrid[nextY][nextX] == '.' {
			newGrid[nextY][nextX] = 'O'
			newGrid[y][x] = '.'
			move(nextX, nextY)
		}
	}

	if deltaIdx == 0 || deltaIdx == 1 {
		for y := range newGrid {
			for x := range newGrid[y] {
				if newGrid[y][x] == 'O' {
					move(x, y)
				}
			}
		}
	} else {
		for y := len(newGrid) - 1; y >= 0; y-- {
			for x := len(newGrid[y]) - 1; x >= 0; x-- {
				if newGrid[y][x] == 'O' {
					move(x, y)
				}
			}
		}
	}

	return newGrid
}

// calculateLoad calculates the total load on the north support beams of the given grid.
func calculateLoad(grid [][]rune) int {
	load := 0
	for x := range grid[0] {
		for y := range grid {
			if grid[y][x] == 'O' {
				load += len(grid) - y
				break
			}
		}
	}
	return load
}

