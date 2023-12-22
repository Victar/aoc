
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dx         = [4]int{0, 1, 0, -1}
	dy         = [4]int{1, 0, -1, 0}
	steps      = 64
	maxResults = 0
)

func canMove(grid [][]byte, x, y, steps int) int {
	if steps == 0 {
		if grid[x][y] == 'O' {
			return 0
		}
		grid[x][y] = 'O'
		return 1
	}

	grid[x][y] = '#'

	result := 0
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			if grid[nx][ny] == '.' {
				result += canMove(grid, nx, ny, steps-1)
			}
		}
	}

	grid[x][y] = '.'

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var (
		grid [][]byte
		sx   int
		sy   int
	)
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		row := []byte(scanner.Text())
		grid = append(grid, row)
		for x, char := range row {
			if char == 'S' {
				sx, sy = x, y
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	grid[sx][sy] = '.' // Mark the starting point as a regular garden plot.
	result := canMove(grid, sx, sy, steps)
	fmt.Println(result)
}
