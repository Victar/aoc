
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read Input File
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]rune, 0)

	for scanner.Scan() {
		row := []rune(scanner.Text())
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	load := calculateLoadAfterRolls(grid, 1000000000)
	fmt.Println(load)
}

func calculateLoad(grid [][]rune) int {
	load := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'O' {
				load += len(grid) - y
			}
		}
	}
	return load
}

func calculateLoadAfterRolls(grid [][]rune, cycles int) int {
	height := len(grid)
	width := len(grid[0])
	
	for cycle := 0; cycle < cycles; cycle++ {
		// Tilt North
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == 'O' {
					for ty := y; ty >= 1; ty-- {
						if grid[ty-1][x] == '.' {
							grid[ty-1][x], grid[ty][x] = grid[ty][x], grid[ty-1][x]
						} else {
							break
						}
					}
				}
			}
		}
		// Tilt West
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] == 'O' {
					for tx := x; tx >= 1; tx-- {
						if grid[y][tx-1] == '.' {
							grid[y][tx-1], grid[y][tx] = grid[y][tx], grid[y][tx-1]
						} else {
							break
						}
					}
				}
			}
		}
		// Tilt South
		for y := height - 1; y >= 0; y-- {
			for x := width - 1; x >= 0; x-- {
				if grid[y][x] == 'O' {
					for ty := y; ty < height-1; ty++ {
						if grid[ty+1][x] == '.' {
							grid[ty+1][x], grid[ty][x] = grid[ty][x], grid[ty+1][x]
						} else {
							break
						}
					}
				}
			}
		}
		// Tilt East
		for y := height - 1; y >= 0; y-- {
			for x := width - 1; x >= 0; x-- {
				if grid[y][x] == 'O' {
					for tx := x; tx < width-1; tx++ {
						if grid[y][tx+1] == '.' {
							grid[y][tx+1], grid[y][tx] = grid[y][tx], grid[y][tx+1]
						} else {
							break
						}
					}
				}
			}
		}
	}

	// Calculate load after 1000000000 cycles
	endLoad := calculateLoad(grid)
	return endLoad
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
