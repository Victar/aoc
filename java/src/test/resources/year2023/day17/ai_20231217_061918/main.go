
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	dx = []int{0, 1, 0, -1} // delta x for right, down, left, up
	dy = []int{1, 0, -1, 0}  // delta y for right, down, left, up
)

type state struct {
	x, y, dir, count int // position, direction, number of consecutive moves in a direction
}

func main() {
	// Read the grid from input.txt
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, r := range line {
			digit, _ := strconv.Atoi(string(r))
			row = append(row, digit)
		}
		grid = append(grid, row)
	}

	// Get dimensions of the grid
	height := len(grid)
	width := len(grid[0])

	// Initialize the minimum heat loss map
	minHeatLoss := make([][][]int, height)
	for i := range minHeatLoss {
		minHeatLoss[i] = make([][]int, width)
		for j := range minHeatLoss[i] {
			minHeatLoss[i][j] = make([]int, 4) // 4 for each direction
			for k := range minHeatLoss[i][j] {
				minHeatLoss[i][j][k] = math.MaxInt32 // Max heat loss initially
			}
		}
	}

	// BFS starting state
	queue := []state{{0, 0, -1, 0}}
	minHeatLoss[0][0][0] = 0 // Start without any heat loss

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check all 3 possible moves from the current state
		for i := -1; i <= 1; i++ {
			newDir := (current.dir + i + 4) % 4
			newX := current.x + dx[newDir]
			newY := current.y + dy[newDir]
			moveCount := current.count
			if current.dir == newDir || current.dir == -1 { // Continue or first move
				moveCount++
			} else { // New direction
				moveCount = 1
			}

			// Validate new position and consecutive move count
			if newX >= 0 && newX < height && newY >= 0 && newY < width && moveCount <= 3 {
				newHeatLoss := grid[newX][newY] + minHeatLoss[current.x][current.y][current.dir%4]
				if newHeatLoss < minHeatLoss[newX][newY][newDir] {
					minHeatLoss[newX][newY][newDir] = newHeatLoss
					queue = append(queue, state{newX, newY, newDir, moveCount})
				}
			}
		}
	}

	// Find the minimum heat loss to reach the bottom-right corner
	result := math.MaxInt32
	for _, loss := range minHeatLoss[height-1][width-1] {
		if loss < result {
			result = loss
		}
	}

	fmt.Println(result)
}
