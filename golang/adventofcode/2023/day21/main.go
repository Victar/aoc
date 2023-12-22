package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "21"

var dr = []int{-1, 1, 0, 0}
var dc = []int{0, 0, 1, -1}

func main() {
	runSilver()
	//runGold()
}

type Point struct {
	x, y int
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var grid [][]bool
	visited := make(map[Point]bool)

	for r, line := range lines {
		row := make([]bool, len(line))
		for c, ch := range line {
			if ch == 'S' {
				visited[Point{r, c}] = true
			}
			row[c] = ch == '.' || ch == 'S'
		}
		grid = append(grid, row)
	}

	const steps = 64
	for i := 0; i < steps; i++ {
		visited = getPoint(visited, grid)
	}
	fmt.Println(len(visited))
	printVisited(grid, visited)
}

func getPoint(visited map[Point]bool, grid [][]bool) map[Point]bool {
	newVisited := make(map[Point]bool)

	rows := len(grid)
	cols := len(grid[0])
	for key, _ := range visited {
		r := key.x
		c := key.y
		for k := 0; k < 4; k++ {
			nr, nc := r+dr[k], c+dc[k]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] {
				if grid[nr][nc] {
					newVisited[Point{nr, nc}] = true
				}
			}
		}
	}
	return newVisited
}

func printVisited(grid [][]bool, visited map[Point]bool) {
	rows := len(grid)
	cols := len(grid[0])
	for row := 0; row < rows; row++ {
		line := []rune{}
		for col := 0; col < cols; col++ {
			cur := '#'
			if grid[row][col] {
				if visited[Point{row, col}] {
					cur = 'o'
				} else {
					cur = '.'
				}

			} else {
				if visited[Point{row, col}] {
					cur = 'x'
				}
			}
			line = append(line, cur)
		}
		println(string(line))
	}
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
