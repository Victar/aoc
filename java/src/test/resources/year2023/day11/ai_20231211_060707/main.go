
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readInput("input.txt")
	if err != nil {
		panic(err)
	}

	expandedGrid := expandUniverse(grid)

	galaxies := findGalaxies(expandedGrid)
	lengthSum := calculateSumOfShortestPaths(galaxies, expandedGrid)

	fmt.Println(lengthSum)
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}
	return lines, scanner.Err()
}

// Expands the universe by duplicating the rows and columns with no galaxies.
func expandUniverse(grid [][]rune) [][]rune {
	colExpansion := make([]bool, len(grid[0]))
	rowExpansion := make([]bool, len(grid))

	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				rowExpansion[y] = true
				colExpansion[x] = true
			}
		}
	}

	expanded := make([][]rune, 0, len(grid)*2)
	for y, row := range grid {
		newRow := make([]rune, 0, len(row)*2)
		for x, cell := range row {
			newRow = append(newRow, cell)
			if !colExpansion[x] {
				newRow = append(newRow, cell)
			}
		}
		expanded = append(expanded, newRow)
		if !rowExpansion[y] {
			expanded = append(expanded, newRow)
		}
	}

	return expanded
}

func findGalaxies(grid [][]rune) []coords {
	var galaxies []coords
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				galaxies = append(galaxies, coords{x: x, y: y})
			}
		}
	}
	return galaxies
}

type coords struct {
	x, y int
}

func calculateSumOfShortestPaths(galaxies []coords, grid [][]rune) int {
	lengthSum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pathLength := bfsShortestPath(grid, galaxies[i], galaxies[j])
			lengthSum += pathLength
		}
	}
	return lengthSum
}

func bfsShortestPath(grid [][]rune, start, end coords) int {
	type entry struct {
		position coords
		distance int
	}

	visited := make(map[coords]bool)
	queue := []entry{{position: start, distance: 0}}
	directions := []coords{
		{x: -1, y: 0},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.position == end {
			return current.distance
		}

		for _, dir := range directions {
			next := coords{x: current.position.x + dir.x, y: current.position.y + dir.y}
			if next.x >= 0 && next.x < len(grid[0]) && next.y >= 0 && next.y < len(grid) {
				if _, seen := visited[next]; !seen {
					visited[next] = true
					queue = append(queue, entry{position: next, distance: current.distance + 1})
				}
			}
		}
	}

	return -1
}
