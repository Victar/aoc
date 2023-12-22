
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row []int
		for _, r := range scanner.Text() {
			val, _ := strconv.Atoi(string(r))
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	rows := len(grid)
	cols := len(grid[0])
	minLoss := math.MaxInt32
	
	var dfs func(row, col, direction, moves, loss int)
	dfs = func(row, col, direction, moves, loss int) {
		if row == rows-1 && col == cols-1 {
			if loss < minLoss {
				minLoss = loss
			}
			return
		}

		if row < 0 || col < 0 || row >= rows || col >= cols {
			return
		}
		
		if moves == 3 {
			switch direction {
			case 0: // moving right
				dfs(row-1, col, 1, 1, loss+grid[row][col])
				dfs(row+1, col, 3, 1, loss+grid[row][col])
			case 1: // moving up
				dfs(row, col+1, 0, 1, loss+grid[row][col])
				dfs(row, col-1, 2, 1, loss+grid[row][col])
			case 2: // moving left
				dfs(row-1, col, 1, 1, loss+grid[row][col])
				dfs(row+1, col, 3, 1, loss+grid[row][col])
			case 3: // moving down
				dfs(row, col+1, 0, 1, loss+grid[row][col])
				dfs(row, col-1, 2, 1, loss+grid[row][col])
			}
		} else {
			switch direction {
			case 0: // moving right
				dfs(row, col+1, 0, moves+1, loss)
				dfs(row-1, col, 1, 1, loss+grid[row][col])
				dfs(row+1, col, 3, 1, loss+grid[row][col])
			case 1: // moving up
				dfs(row, col+1, 0, 1, loss+grid[row][col])
				dfs(row-1, col, 1, moves+1, loss)
				dfs(row, col-1, 2, 1, loss+grid[row][col])
			case 2: // moving left
				dfs(row, col-1, 2, moves+1, loss)
				dfs(row-1, col, 1, 1, loss+grid[row][col])
				dfs(row+1, col, 3, 1, loss+grid[row][col])
			case 3: // moving down
				dfs(row, col+1, 0, 1, loss+grid[row][col])
				dfs(row, col-1, 2, 1, loss+grid[row][col])
				dfs(row+1, col, 3, moves+1, loss)
			}
		}
	}

	dfs(0, 0, 0, 0, 0)
	fmt.Println(minLoss)
}
