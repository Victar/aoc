
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MaxSteps = 3

func main() {
	grid, err := readGridFromFile("input.txt")
	if err != nil {
		panic(err)
	}
	
	minHeatLoss := findMinHeatLoss(grid)
	fmt.Println(minHeatLoss)
}

func findMinHeatLoss(grid [][]int) int {
	// We'll keep a memoization table where the key is a tuple
	// (row, col, stepsInDirection, previousDirection)
	memo := make(map[[4]int]int)
	rows := len(grid)
	cols := len(grid[0])
	return dfs(grid, 0, 0, 0, -1, rows, cols, memo)
}

// dfs performs a depth-first search and returns the minimum heat loss.
func dfs(grid [][]int, row, col, steps, prevDir, rows, cols int, memo map[[4]int]int) int {
	// If out of bounds or steps are exceeded, return infinity
	if row < 0 || col < 0 || row >= rows || col >= cols || steps > MaxSteps {
		return 1 << 30
	}

	if row == rows-1 && col == cols-1 {
		return grid[row][col] // Reached destination - no more heat loss
	}

	key := [4]int{row, col, steps, prevDir}
	if val, ok := memo[key]; ok {
		return val
	}

	// Apply the heat loss for this block unless it's the starting point
	heatLoss := 0
	if row != 0 || col != 0 {
		heatLoss = grid[row][col]
	}

	minHeatLoss := 1 << 30 // Start with a very large number

	// The next three directions we can move in
	directions := [][]int{
		{0, 1}, // Right
		{-1, 0}, // Up
		{0, -1}, // Left
		{1, 0}, // Down
	}

	// Try all possible moves
	for nextDir, dir := range directions {
		if nextDir != prevDir {
			newSteps := 1
			if (prevDir+2)%4 == nextDir { // if we continue in the same general direction
				newSteps += steps
			}
			nextRow := row + dir[0]
			nextCol := col + dir[1]
			heat := dfs(grid, nextRow, nextCol, newSteps, nextDir, rows, cols, memo)
			if heatLoss+heat < minHeatLoss {
				minHeatLoss = heatLoss + heat
			}
		}
	}

	memo[key] = minHeatLoss
	return minHeatLoss
}

func readGridFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, runeValue := range line {
			i, err := strconv.Atoi(string(runeValue))
			if err != nil {
				return nil, err
			}
			row = append(row, i)
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}
