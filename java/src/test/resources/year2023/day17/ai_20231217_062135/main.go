
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	maxTurns = 3 // Maximum number of turns in a single direction
)

var (
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

type Point struct {
	x, y, dir, turns int
}

func minHeatLoss(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	memo := make(map[Point]int)

	var search func(x, y, dir, turns int) int
	search = func(x, y, dir, turns int) int {
		if x == rows-1 && y == cols-1 {
			return 0
		}

		point := Point{x, y, dir, turns}
		if val, ok := memo[point]; ok {
			return val
		}

		heatLoss := math.MaxInt32
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				if i == dir {
					if turns < maxTurns {
						heat := grid[nx][ny] + search(nx, ny, i, turns+1)
						if heat < heatLoss {
							heatLoss = heat
						}
					}
				} else {
					heat := grid[nx][ny] + search(nx, ny, i, 1)
					if heat < heatLoss {
						heatLoss = heat
					}
				}
			}
		}

		memo[point] = heatLoss
		return heatLoss
	}

	initialDir := 0 // Initial direction, it's actually irrelevant at the start point.
	initialTurns := 0
	minHeat := search(0, 0, initialDir, initialTurns)
	return minHeat
}

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
			val, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			row = append(row, val)
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(minHeatLoss(grid))
}
