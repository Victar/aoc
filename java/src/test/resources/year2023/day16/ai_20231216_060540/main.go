
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	grid := make([][]rune, 0) // This will store the grid layout
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	type point struct {
		x, y int
	}
	directions := map[point]point{
		{0, -1}: {0, 1},
		{0, 1}:  {0, -1},
		{-1, 0}: {1, 0},
		{1, 0}:  {-1, 0},
	}

	var turn = func(dir point, tile rune) point {
		if tile == '/' {
			return point{-dir.y, -dir.x}
		}
		if tile == '\\' {
			return point{dir.y, dir.x}
		}
		return dir
	}

	var energizedTiles int
	visit := make(map[point]bool)
	var dfs func(x, y int, dir point)

	dfs = func(x, y int, dir point) {
		if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
			return
		}

		pos := point{x, y}
		if !visit[pos] {
			visit[pos] = true
			energizedTiles++
		}

		tile := grid[y][x]

		switch tile {
		case '.', '|', '-':
			dfs(x+dir.x, y+dir.y, dir)
		case '/', '\\':
			newDir := turn(dir, tile)
			dfs(x+newDir.x, y+newDir.y, newDir)
		}

		if tile == '|' && dir.y != 0 || tile == '-' && dir.x != 0 {
			dfs(x+directions[dir].x, y+directions[dir].y, directions[dir])
		}
	}

	dfs(0, 0, point{1, 0}) // Start from the top-left corner going right

	fmt.Println(energizedTiles)
}
