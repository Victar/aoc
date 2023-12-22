
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

type Point struct {
	x, y int
}

func reflect(direction int, mirror rune) int {
	switch mirror {
	case '/':
		switch direction {
		case UP:
			return LEFT
		case DOWN:
			return RIGHT
		case LEFT:
			return UP
		case RIGHT:
			return DOWN
		}
	case '\\':
		switch direction {
		case UP:
			return RIGHT
		case DOWN:
			return LEFT
		case LEFT:
			return DOWN
		case RIGHT:
			return UP
		}
	}
	return direction
}

func loadGrid() ([][]rune, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid, scanner.Err()
}

func inGrid(p Point, grid [][]rune) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[p.y])
}

func nextPoint(p Point, dir int) Point {
	switch dir {
	case UP:
		p.y--
	case DOWN:
		p.y++
	case LEFT:
		p.x--
	case RIGHT:
		p.x++
	}
	return p
}

func solve(grid [][]rune) int {
	direction := RIGHT
	pos := Point{x: 0, y: 0}
	energized := make(map[Point]bool)

	for inGrid(pos, grid) {
		switch grid[pos.y][pos.x] {
		case '.', '|', '-':
			// Continue in the same direction
		case '/', '\\':
			direction = reflect(direction, grid[pos.y][pos.x])
		default:
			// If we encounter an unknown character, ignore it
		}

		energized[pos] = true

		if grid[pos.y][pos.x] == '|' || grid[pos.y][pos.x] == '-' {
			// If it's a splitter, create a new beam
			if direction == UP || direction == DOWN {
				newPos := nextPoint(pos, LEFT)
				if inGrid(newPos, grid) && grid[newPos.y][newPos.x] == '.' {
					energized[newPos] = true
				}
				newPos = nextPoint(pos, RIGHT)
				if inGrid(newPos, grid) && grid[newPos.y][newPos.x] == '.' {
					energized[newPos] = true
				}
			} else {
				newPos := nextPoint(pos, UP)
				if inGrid(newPos, grid) && grid[newPos.y][newPos.x] == '.' {
					energized[newPos] = true
				}
				newPos = nextPoint(pos, DOWN)
				if inGrid(newPos, grid) && grid[newPos.y][newPos.x] == '.' {
					energized[newPos] = true
				}
			}
		}

		pos = nextPoint(pos, direction)
	}

	return len(energized)
}

func main() {
	grid, err := loadGrid()
	if err != nil {
		fmt.Println("Error loading grid:", err)
		os.Exit(1)
	}

	result := solve(grid)
	fmt.Println(result)
}
