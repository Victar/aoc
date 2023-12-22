package main

import (
	"bufio"
	"fmt"
	"os"
)

// Directions: N, S, E, W
var dr = []int{-1, 1, 0, 0}
var dc = []int{0, 0, 1, -1}

func readInput(filename string) ([][]bool, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]bool
	var startR, startC int
	for r := 0; scanner.Scan(); r++ {
		line := scanner.Text()
		row := make([]bool, len(line))
		for c, ch := range line {
			if ch == 'S' {
				startR, startC = r, c // Save the starting position
			}
			row[c] = ch == '.' || ch == 'S'
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid, startR, startC
}

func stepCounter(grid [][]bool, startR, startC, steps int) int {
	rows := len(grid)
	cols := len(grid[0])
	dpCurr := make([][]int, rows)
	dpPrev := make([][]int, rows)
	for i := range dpCurr {
		dpCurr[i] = make([]int, cols)
		dpPrev[i] = make([]int, cols)
	}

	dpPrev[startR][startC] = 1

	for s := 0; s < steps; s++ {
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] {
					dpCurr[r][c] = 0
					for k := 0; k < 4; k++ {
						nr, nc := r+dr[k], c+dc[k]
						if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] {
							dpCurr[r][c] += dpPrev[nr][nc]
						}
					}
				}
			}
		}
		dpCurr, dpPrev = dpPrev, dpCurr // Swap the dp arrays
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dpPrev[r][c] > 0 {
				count++
			}
		}
	}

	return count
}

func main() {
	grid, startR, startC := readInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day21/input.txt")
	const steps = 64
	result := stepCounter(grid, startR, startC, steps)
	fmt.Print(result)
}
