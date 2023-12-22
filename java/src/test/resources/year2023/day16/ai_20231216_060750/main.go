
// main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	grid := readInput("input.txt")
	count := countEnergizedTiles(grid)
	fmt.Println(count)
}

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func countEnergizedTiles(grid [][]rune) int {
	visited := make(map[Point][]Point)
	var beam = Point{0, 1} // starting rightward
	processBeam(grid, &visited, Point{0, 0}, beam)
	count := 0
	for _, directions := range visited {
		count += 1
	}
	return count
}

func processBeam(grid [][]rune, visited *map[Point][]Point, position, direction Point) {
	for {
		position.X += direction.X
		position.Y += direction.Y

		if position.X < 0 || position.X >= len(grid) || position.Y < 0 || position.Y >= len(grid[position.X]) {
			return
		}

		cell := grid[position.X][position.Y]
		switch cell {
		case '.', '|', '-':
			if cell != '.' {
				mirrorReflected := reflect(direction, cell)
				if mirrorReflected != (Point{0, 0}) {
					direction = mirrorReflected
				}
			}
		case '/':
			direction = Point{-direction.Y, -direction.X}
		case '\\':
			direction = Point{direction.Y, direction.X}
		default:
			return
		}

		_, found := (*visited)[position]
		if !found {
			(*visited)[position] = []Point{direction}
		} else if !contains((*visited)[position], direction) {
			(*visited)[position] = append((*visited)[position], direction)
		}
	}
}

func reflect(direction Point, cell rune) Point {
	if (direction.X == 0 && cell == '|') || (direction.Y == 0 && cell == '-') {
		return Point{-direction.X, -direction.Y} // Reflect
	} else if (direction.X != 0 && cell == '|') || (direction.Y != 0 && cell == '-') {
		// Split
		horizontal := Point{0, direction.Y}
		vertical := Point{direction.X, 0}
		return Point{horizontal.X + vertical.X, horizontal.Y + vertical.Y} // Combine
	}
	return Point{0, 0} // No reflection
}

func contains(slice []Point, item Point) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
