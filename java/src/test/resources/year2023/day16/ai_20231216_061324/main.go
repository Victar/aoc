
// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type Direction int

const (
	Right Direction = iota
	Left
	Up
	Down
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	energizedTiles := simulateLightBeam(grid)
	fmt.Println(energizedTiles)
}

func simulateLightBeam(grid []string) int {
	height := len(grid)
	width := len(grid[0])
	energized := make(map[Point]bool)
	directions := map[Direction][]int{
		Right: {0, 1},
		Left:  {0, -1},
		Up:    {-1, 0},
		Down:  {1, 0},
	}
	mirrorDirection := map[rune]map[Direction]Direction{
		'/': {
			Right: Up,
			Down:  Left,
			Left:  Down,
			Up:    Right,
		},
		'\\': {
			Right: Down,
			Up:    Left,
			Left:  Up,
			Down:  Right,
		},
	}

	// Starting point
	queue := []Point{{x: 0, y: 0}}
	currentDir := Right

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for {
			if p.x < 0 || p.x >= height || p.y < 0 || p.y >= width {
				break
			}

			// Mark the current tile as energized
			energized[p] = true

			tile := rune(grid[p.x][p.y])
			switch tile {
			case '.', '|', '-':
				dx, dy := directions[currentDir][0], directions[currentDir][1]
				p.x += dx
				p.y += dy
				if tile != '.' {
					// Additional beams for splitter
					if tile == '|' && (currentDir == Left || currentDir == Right) {
						queue = append(queue, Point{p.x + 1, p.y})
						queue = append(queue, Point{p.x - 1, p.y})
					} else if tile == '-' && (currentDir == Up || currentDir == Down) {
						queue = append(queue, Point{p.x, p.y + 1})
						queue = append(queue, Point{p.x, p.y - 1})
					}
					break
				}
			case '/', '\\':
				currentDir = mirrorDirection[tile][currentDir]
				dx, dy := directions[currentDir][0], directions[currentDir][1]
				p.x += dx
				p.y += dy
			default:
				break
			}
		}
	}

	return len(energized)
}
