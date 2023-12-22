
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct{ Row, Col int }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			value, _ := strconv.Atoi(string(c))
			row = append(row, value)
		}
		grid = append(grid, row)
	}

	minHeatLoss := findMinHeatLoss(grid)
	fmt.Println(minHeatLoss)
}

func findMinHeatLoss(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	minHeatLossGrid := initializeHeatLossGrid(rows, cols)

	for i := 1; i <= 3; i++ {
		for r := 0; r < rows; r++ {
			for c := i; c < cols; c++ {
				updateHeatLoss(r, c, i-1, 0, grid, minHeatLossGrid)
			}
		}
		for r := i; r < rows; r++ {
			for c := 0; c < cols; c++ {
				updateHeatLoss(r, c, 0, i-1, grid, minHeatLossGrid)
			}
		}
	}

	return minHeatLossGrid[rows-1][cols-1]
}

func initializeHeatLossGrid(rows, cols int) [][]int {
	heatLossGrid := make([][]int, rows)
	for i := range heatLossGrid {
		heatLossGrid[i] = make([]int, cols)
		for j := range heatLossGrid[i] {
			heatLossGrid[i][j] = 1<<31 - 1 // Set to max int to find minimum later
		}
	}
	heatLossGrid[0][0] = 0
	return heatLossGrid
}

func updateHeatLoss(row, col, dRow, dCol int, grid, minHeatLossGrid [][]int) {
	newRow, newCol := row-dRow, col-dCol
	if newRow >= 0 && newCol >= 0 {
		newHeat := minHeatLossGrid[newRow][newCol] + grid[row][col]
		if newHeat < minHeatLossGrid[row][col] {
			minHeatLossGrid[row][col] = newHeat
			for i := 1; i <= 3; i++ {
				if i != dRow { // Not straight up or down from the current direction
					if row+i < len(grid) {
						updateHeatLoss(row+i, col, 1, 0, grid, minHeatLossGrid)
					}
					if row-i >= 0 {
						updateHeatLoss(row-i, col, -1, 0, grid, minHeatLossGrid)
					}
				}
				if i != dCol { // Not straight left or right from the current direction
					if col+i < len(grid[0]) {
						updateHeatLoss(row, col+i, 0, 1, grid, minHeatLossGrid)
					}
					if col-i >= 0 {
						updateHeatLoss(row, col-i, 0, -1, grid, minHeatLossGrid)
					}
				}
			}
		}
	}
}
