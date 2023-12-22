
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

const (
	maxTurn = 3
)

func readGrid(filename string) ([][]int, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, c := range line {
			grid[i][j] = int(c - '0')
		}
	}

	return grid, nil
}

func minHeatLoss(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	type state struct {
		x, y, dx, dy, streak, loss int
	}

	var dfs func(s state) int
	cache := map[state]int{}

	dfs = func(s state) int {
		if val, ok := cache[s]; ok {
			return val
		}

		if s.x == n-1 && s.y == m-1 {
			cache[s] = s.loss
			return s.loss
		}

		minLoss := math.MaxInt
		directions := []struct{ dx, dy int }{
			{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		}

		for _, d := range directions {
			if s.dx == d.dx && s.dy == d.dy && s.streak == maxTurn {
				continue // Cannot continue in the same direction
			}

			if s.dx == d.dx && s.dy == d.dy ||
				(s.dx == 0 && s.dy == 0) {
				// continue in the same direction or for the first move
				nx, ny := s.x+d.dx, s.y+d.dy
				if nx >= 0 && nx < n && ny >= 0 && ny < m {
					ns := state{nx, ny, d.dx, d.dy, s.streak + 1, s.loss + grid[ny][nx]}
					loss := dfs(ns)
					if loss < minLoss {
						minLoss = loss
					}
				}
			} else if (s.dx != -d.dx || s.dy != -d.dy) && (s.dx != 0 || s.dy != 0) {
				// Turn to a different direction, not reverse or first move
				nx, ny := s.x+d.dx, s.y+d.dy
				if nx >= 0 && nx < n && ny >= 0 && ny < m {
					ns := state{nx, ny, d.dx, d.dy, 1, s.loss + grid[ny][nx]}
					loss := dfs(ns)
					if loss < minLoss {
						minLoss = loss
					}
				}
			}
		}

		cache[s] = minLoss
		return minLoss
	}

	initialState := state{0, 0, 0, 0, 1, 0}
	return dfs(initialState)
}

func main() {
	grid, err := readGrid("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(minHeatLoss(grid))
}
