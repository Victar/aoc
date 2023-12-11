
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	dx = map[rune]int{'N': 0, 'S': 0, 'E': 1, 'W': -1}
	dy = map[rune]int{'N': -1, 'S': 1, 'E': 0, 'W': 0}
)

// A Tile represents a single character in the pipe maze.
type Tile struct {
	Char  rune
	North bool
	South bool
	East  bool
	West  bool
}

// PipeMaze holds the data structure for the maze of pipes.
type PipeMaze struct {
	Grid map[int]map[int]*Tile
}

// NewPipeMaze creates a new PipeMaze from input data.
func NewPipeMaze(input []string) *PipeMaze {
	pm := &PipeMaze{
		Grid: make(map[int]map[int]*Tile),
	}

	for y, line := range input {
		for x, char := range line {
			if pm.Grid[y] == nil {
				pm.Grid[y] = make(map[int]*Tile)
			}
			tile := &Tile{Char: char}
			switch char {
			case '|', 'S':
				tile.North = true
				tile.South = true
			case '-':
				tile.East = true
				tile.West = true
			case 'L':
				tile.North = true
				tile.East = true
			case 'J':
				tile.North = true
				tile.West = true
			case '7':
				tile.South = true
				tile.West = true
			case 'F':
				tile.South = true
				tile.East = true
			}
			pm.Grid[y][x] = tile
		}
	}

	return pm
}

// FindStart finds the starting position of the animal in the maze.
func (pm *PipeMaze) FindStart() (int, int) {
	for y, row := range pm.Grid {
		for x, tile := range row {
			if tile.Char == 'S' {
				return x, y
			}
		}
	}
	return -1, -1
}

// FindFarthestPoint returns the maximum steps required to travel from start
func (pm *PipeMaze) FindFarthestPoint(startX, startY int) int {
	directions := []rune{'N', 'E', 'S', 'W'}
	visited := map[int]map[int]bool{}

	var dfs func(x, y int, dir rune, dist int) int
	dfs = func(x, y int, dir rune, dist int) int {
		if visited[y][x] {
			return dist - 1 // We've completed the loop; return the distance
		}
		if visited[y] == nil {
			visited[y] = make(map[int]bool)
		}
		visited[y][x] = true

		maxDist := dist
		for _, nextDir := range directions {
			nextX, nextY := x+dx[nextDir], y+dy[nextDir]
			nextTile, exists := pm.Grid[nextY][nextX]
			if !exists || !nextTile.CanEnterFrom(oppositeDir(nextDir)) {
				continue
			}
			maxDist = max(maxDist, dfs(nextX, nextY, nextDir, dist+1))
		}

		return maxDist
	}

	tile := pm.Grid[startY][startX]
	for _, dir := range directions {
		if tile.CanEnterFrom(oppositeDir(dir)) {
			x, y := startX+dx[dir], startY+dy[dir]
			maxDist := dfs(x, y, dir, 1)
			if maxDist > 0 {
				return maxDist
			}
		}
	}

	return -1
}

func (t *Tile) CanEnterFrom(dir rune) bool {
	switch dir {
	case 'N':
		return t.South
	case 'S':
		return t.North
	case 'E':
		return t.West
	case 'W':
		return t.East
	}
	return false
}

func oppositeDir(dir rune) rune {
	switch dir {
	case 'N':
		return 'S'
	case 'S':
		return 'N'
	case 'E':
		return 'W'
	case 'W':
		return 'E'
	}
	panic("invalid direction")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inputLines []string
	for scanner.Scan() {
		inputLines = append(inputLines, strings.TrimSpace(scanner.Text()))
	}

	pipeMaze := NewPipeMaze(inputLines)
	startX, startY := pipeMaze.FindStart()
	farthest := pipeMaze.FindFarthestPoint(startX, startY)
	fmt.Println(farthest)
}
