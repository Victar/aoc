
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(solve(grid))
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid, scanner.Err()
}

func solve(grid [][]rune) int {
	// Assume initial beam position and direction (to the right)
	x, y, dx, dy := 0, 0, 1, 0
	energized := make(map[[2]int]bool)

	for y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		tile := grid[y][x]
		energized[[2]int{y, x}] = true

		switch tile {
		case '/':
			dx, dy = -dy, -dx
		case '\\':
			dx, dy = dy, dx
		case '-':
			if dy != 0 {
				dy = -dy
			}
		case '|':
			if dx != 0 {
				dx = -dx
			}
		}

		// Check for splitters, split the beam if necessary
		if (dx != 0 && tile == '|') || (dy != 0 && tile == '-') {
			// Splitter encountered; calculate the split's directions
			splitDx, splitDy := dy, -dx
			followBeam(grid, x+splitDx, y+splitDy, splitDx, splitDy, energized)
			splitDx, splitDy = -dy, dx
			followBeam(grid, x+splitDx, y+splitDy, splitDx, splitDy, energized)
			break
		}

		x += dx
		y += dy
	}

	return len(energized)
}

func followBeam(grid [][]rune, x, y, dx, dy int, energized map[[2]int]bool) {
	for y >= 0 && y < len(grid) && x >= 0 && x < len(grid[y]) {
		tile := grid[y][x]
		energized[[2]int{y, x}] = true

		switch tile {
		case '/':
			dx, dy = -dy, -dx
		case '\\':
			dx, dy = dy, dx
		case '-':
			if dy != 0 {
				dy = -dy
			}
		case '|':
			if dx != 0 {
				dx = -dx
			}
		}

		if (dx != 0 && tile == '|') || (dy != 0 && tile == '-') {
			// Reached another splitter, stop this beam to prevent infinite loops
			break
		}

		x += dx
		y += dy
	}
}
