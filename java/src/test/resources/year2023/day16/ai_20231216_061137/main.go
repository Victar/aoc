
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Read the input file.
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	// Read the grid into a slice of strings.
	scanner := bufio.NewScanner(input)
	var grid []string
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Simulate the beam of light.
	energizedTiles := simulateBeam(grid)

	// Output the number of energized tiles.
	fmt.Println(energizedTiles)
}

// simulateBeam runs the simulation and returns number of energized tiles.
func simulateBeam(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	energized := make([][]bool, rows)
	for i := range energized {
		energized[i] = make([]bool, cols)
	}

	var dfs func(r, c int, dirR, dirC int)
	dfs = func(r, c int, dirR, dirC int) {
		// Check if the position is out of the grid bounds.
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}

		// Energize the tile.
		energized[r][c] = true

		switch grid[r][c] {
		case '.':
			// Continue in the same direction.
			dfs(r+dirR, c+dirC, dirR, dirC)

		case '/':
			// Reflect the beam 90 degrees.
			newDirR, newDirC := dirC, dirR
			dfs(r+newDirR, c+newDirC, newDirR, newDirC)

		case '\\':
			// Reflect the beam 90 degrees.
			newDirR, newDirC := -dirC, -dirR
			dfs(r+newDirR, c+newDirC, newDirR, newDirC)

		case '|', '-':
			// Check if we're hitting the pointy end of a splitter.
			if (dirR == 0 && grid[r][c] == '-') || (dirC == 0 && grid[r][c] == '|') {
				// Continue in the same direction.
				dfs(r+dirR, c+dirC, dirR, dirC)
			} else {
				// Split the beam if hitting the flat side of a splitter.
				if dirR != 0 {
					dfs(r, c+1, 0, 1)
					dfs(r, c-1, 0, -1)
				} else {
					dfs(r+1, c, 1, 0)
					dfs(r-1, c, -1, 0)
				}
			}
		}
	}

	// Initial direction is right.
	dfs(0, 0, 0, 1)

	// Count the number of energized tiles.
	count := 0
	for _, row := range energized {
		for _, tile := range row {
			if tile {
				count++
			}
		}
	}

	return count
}
