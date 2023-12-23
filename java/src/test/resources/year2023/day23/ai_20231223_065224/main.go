
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	grid            [][]rune
	visited         [][]bool
	longestHike     int
	directions      = map[rune]func(int, int) (int, int){}
	startR, startC  int
	dx              = []int{0, 0, -1, 1}
	dy              = []int{1, -1, 0, 0}
)

func init() {
    // Initialize the directions map for the slopes
	directions['^'] = func(r, c int) (int, int) { return r - 1, c }
	directions['v'] = func(r, c int) (int, int) { return r + 1, c }
	directions['>'] = func(r, c int) (int, int) { return r, c + 1 }
	directions['<'] = func(r, c int) (int, int) { return r, c - 1 }
}

func dfs(row, col, steps int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) ||
		visited[row][col] || grid[row][col] == '#' {
		return
	}
	if row == len(grid)-1 && grid[row][col] == '.' {
		if steps > longestHike {
			longestHike = steps
		}
		return
	}
	visited[row][col] = true
	defer func() { visited[row][col] = false }()

	if moveFunc, ok := directions[grid[row][col]]; ok {
		newR, newC := moveFunc(row, col)
		dfs(newR, newC, steps+1)
	} else {
		for i := 0; i < 4; i++ {
			newR, newC := row+dx[i], col+dy[i]
			dfs(newR, newC, steps+1)
		}
	}
}

func readInput(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, char := range line {
			row = append(row, char)
			if char == '.' {
				startR, startC = len(grid), len(row)-1
			}
		}
		grid = append(grid, row)
	}
	visited = make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	return scanner.Err()
}

func main() {
	err := readInput("input.txt")
	if err != nil {
		fmt.Println("Failed to read input:", err)
		return
	}

	dfs(startR, startC, 0)

	fmt.Println(longestHike)
}
