
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid, startX, startY := loadGrid("input.txt")
	if grid == nil {
		fmt.Println("Failed to load grid.")
		os.Exit(1)
	}

	visited := make(map[int]map[int]struct{})
	for i := range grid {
		visited[i] = make(map[int]struct{})
	}

	traverseGrid(grid, visited, startX, startY)

	maxDistance := 0
	queue := []pair{{x: startX, y: startY, dist: 0}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.dist > maxDistance {
			maxDistance = p.dist
		}

		for _, dir := range directions {
			newX, newY := p.x+dir.dx, p.y+dir.dy
			if _, ok := visited[newX][newY]; ok {
				delete(visited[newX], newY)
				queue = append(queue, pair{x: newX, y: newY, dist: p.dist + 1})
			}
		}
	}

	fmt.Println(maxDistance)
}

type pair struct {
	x, y, dist int
}

var directions = []pair{
	{dx: -1, dy: 0}, // Up
	{dx: 1, dy: 0},  // Down
	{dx: 0, dy: -1}, // Left
	{dx: 0, dy: 1},  // Right
}

func canTravel(grid [][]rune, x, y, dx, dy int) bool {
	if x+dx < 0 || x+dx >= len(grid) || y+dy < 0 || y+dy >= len(grid[0]) {
		return false
	}
	r := grid[x+dx][y+dy]
	switch r {
	case '|', '-', 'L', 'J', '7', 'F', 'S':
		return true
	}
	return false
}

func traverseGrid(grid [][]rune, visited map[int]map[int]struct{}, x, y int) {
	if _, ok := visited[x][y]; ok {
		return
	}
	visited[x][y] = struct{}{}

	for _, dir := range directions {
		newX, newY := x+dir.dx, y+dir.dy
		if canTravel(grid, x, y, dir.dx, dir.dy) {
			traverseGrid(grid, visited, newX, newY)
		}
	}
}

func loadGrid(filename string) ([][]rune, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, 0, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	startX, startY := -1, -1
	var grid [][]rune

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := []rune(line)
		for x, char := range row {
			if char == 'S' {
				startX, startY = y, x
			}
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return nil, 0, 0
	}

	return grid, startX, startY
}
