
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	matrix := readInput("input.txt")
	galaxies := locateGalaxies(matrix)

	expandFactor := []int{2, 1000000} // Factors for Parts 1 and 2
	for _, factor := range expandFactor {
		expandedMatrix := expandMatrix(matrix, factor)
		total := calculateTotalShortestPaths(galaxies, expandedMatrix, factor)
		fmt.Println(total)
	}
}

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var matrix [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		matrix = append(matrix, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return matrix
}

func locateGalaxies(matrix [][]rune) [][2]int {
	galaxies := [][2]int{}
	for y, row := range matrix {
		for x, char := range row {
			if char == '#' {
				galaxies = append(galaxies, [2]int{x, y})
			}
		}
	}
	return galaxies
}

func expandMatrix(matrix [][]rune, factor int) [][]rune {
	rows, cols := len(matrix), len(matrix[0])
	emptyRows, emptyCols := make([]bool, rows), make([]bool, cols)

	// Determine which rows and columns are empty
	for y, row := range matrix {
		for x, char := range row {
			if char == '#' {
				emptyCols[x] = true
				emptyRows[y] = true
			}
		}
	}
	for i := range emptyRows {
		emptyRows[i] = !emptyRows[i]
	}
	for i := range emptyCols {
		emptyCols[i] = !emptyCols[i]
	}

	// Count the total number of rows and columns after expansion
	newRows, newCols := 0, 0
	for _, empty := range emptyRows {
		if empty {
			newRows += factor
		} else {
			newRows++
		}
	}

	for _, empty := range emptyCols {
		if empty {
			newCols += factor
		} else {
			newCols++
		}
	}

	// Create a new expanded matrix with the appropriate size
	expandedMatrix := make([][]rune, newRows)
	for i := range expandedMatrix {
		expandedMatrix[i] = make([]rune, newCols)
	}

	// Map old positions to the new expanded matrix
	for oldY := 0; oldY < rows; oldY++ {
		for oldX := 0; oldX < cols; oldX++ {
			newY := 0
			for i := 0; i < oldY; i++ {
				if emptyRows[i] {
					newY += factor
				} else {
					newY++
				}
			}
			newX := 0
			for i := 0; i < oldX; i++ {
				if emptyCols[i] {
					newX += factor
				} else {
					newX++
				}
			}

			expandedMatrix[newY][newX] = matrix[oldY][oldX]
		}
	}

	return expandedMatrix
}

func calculateTotalShortestPaths(galaxies [][2]int, matrix [][]rune, factor int) int {
	sum := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			// This uses the Manhattan distance because galaxies are only connected in up, down, left, and right directions.
			// Since the empty rows and columns have expanded uniformly, the Manhattan distance formula still holds for the shortest path.
			dist := abs(g2[0]-g1[0]) + abs(g2[1]-g1[1])
			sum += adjustDistanceForExpansion(matrix, g1, g2, dist, factor)
		}
	}
	return sum
}

func adjustDistanceForExpansion(matrix [][]rune, g1, g2 [2]int, dist, factor int) int {
	if factor == 2 {
		return dist // For part 1 expansion factor of 2 does not affect the galaxy distance
	}
	for i := min(g1[0], g2[0]); i <= max(g1[0], g2[0]); i++ {
		if matrix[g1[1]][i] == '.' {
			dist += factor - 1
		}
	}
	for i := min(g1[1], g2[1]); i <= max(g1[1], g2[1]); i++ {
		if matrix[i][g1[0]] == '.' {
			dist += factor - 1
		}
	}
	return dist
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
