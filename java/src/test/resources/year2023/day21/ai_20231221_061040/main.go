
package main

import (
	"bufio"
	"fmt"
	"os"
)

const steps = 64

var (
	dirs = [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} // N, E, S, W
)

// Count the number of reachable plots from the starting 'S' in exactly 'steps' moves
func countReachablePlots(grid [][]rune, startX, startY int) int {
	memo := map[[3]int]bool{}
	return dfs(startX, startY, steps, grid, memo)
}

// Depth-first search that counts unique reachable plots
func dfs(x, y, remainingSteps int, grid [][]rune, memo map[[3]int]bool) int {
	if remainingSteps == 0 {
		if grid[y][x] == '.' {
			return 1
		}
		return 0
	}

	state := [3]int{x, y, remainingSteps}
	if seen, ok := memo[state]; ok && seen {
		return 0
	}
	memo[state] = true

	height, width := len(grid), len(grid[0])
	count := 0
	for _, dir := range dirs {
		nextX, nextY := x+dir[0], y+dir[1]
		if nextX >= 0 && nextX < width && nextY >= 0 && nextY < height && grid[nextY][nextX] != '#' {
			count += dfs(nextX, nextY, remainingSteps-1, grid, memo)
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	var startX, startY int

	// Read grid from file and find the starting 'S' position
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		gridRow := []rune(line)
		grid = append(grid, gridRow)

		if sx := runeIndex('S', gridRow); sx != -1 {
			startX, startY = sx, y
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning input file:", err)
		return
	}

	// Call function to count reachable plots and print result
	result := countReachablePlots(grid, startX, startY)
	fmt.Println(result)
}

// Helper function to find the index of rune 'r' in a slice of runes 's'
func runeIndex(r rune, s []rune) int {
	for i, runeValue := range s {
		if runeValue == r {
			return i
		}
	}
	return -1
}
