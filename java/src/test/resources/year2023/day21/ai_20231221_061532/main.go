
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	filename        = "input.txt"
	steps           = 64
	directions      = 4
	startMarker     = 'S'
	gardenPlot      = '.'
	invalidPosition = '#'
)

var (
	dx = [directions]int{0, 1, 0, -1}
	dy = [directions]int{1, 0, -1, 0}
)

func canMove(grid [][]rune, x, y int) bool {
	if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != invalidPosition {
		return true
	}
	return false
}

func countSteps(grid [][]rune, sx, sy, remainingSteps int, memo map[int]map[int]map[int]int) int {
	if remainingSteps == 0 {
		return 1
	}

	if _, ok := memo[sx][sy][remainingSteps]; ok {
		return memo[sx][sy][remainingSteps]
	}

	count := 0
	for i := 0; i < directions; i++ {
		nx, ny := sx+dx[i], sy+dy[i]
		if canMove(grid, nx, ny) {
			count += countSteps(grid, nx, ny, remainingSteps-1, memo)
		}
	}

	if _, ok := memo[sx]; !ok {
		memo[sx] = make(map[int]map[int]int)
	}
	if _, ok := memo[sx][sy]; !ok {
		memo[sx][sy] = make(map[int]int)
	}
	memo[sx][sy][remainingSteps] = count
	return count
}

func main() {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	grid := make([][]rune, len(lines))
	startX, startY := 0, 0

	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		for j, char := range line {
			grid[i][j] = char
			if char == startMarker {
				startX, startY = i, j
			}
		}
	}

	memo := make(map[int]map[int]map[int]int)
	reachable := countSteps(grid, startX, startY, steps, memo)

	fmt.Println(reachable)
}
