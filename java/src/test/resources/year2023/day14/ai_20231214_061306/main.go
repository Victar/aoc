
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	totalLoad := calculateLoad(grid)
	fmt.Println(totalLoad)
}

func calculateLoad(grid []string) int {
	tiltedGrid := tiltNorth(grid)
	rows := len(tiltedGrid)
	totalLoad := 0

	for _, row := range tiltedGrid {
		for idx, char := range row {
			if char == 'O' {
				totalLoad += rows - idx
			}
		}
	}

	return totalLoad
}

func tiltNorth(grid []string) []string {
	columns := len(grid[0])
	tilted := make([]string, columns)

	for i := 0; i < columns; i++ {
		var col []rune
		for _, row := range grid {
			col = append(col, rune(row[i]))
		}

		col = moveRocks(col)
		for j := 0; j < len(col); j++ {
			if len(tilted[j]) < i+1 {
				tilted[j] = strings.Repeat(".", i)
			}
			tilted[j] += string(col[j])
		}
	}

	return tilted
}

func moveRocks(col []rune) []rune {
	result := make([]rune, len(col))
	copy(result, col)

	for i := 0; i < len(result); i++ {
		if result[i] == 'O' {
			for j := i; j >= 0; j-- {
				if result[j] == '.' {
					result[j], result[i] = result[i], result[j]
					i = j
				} else {
					break
				}
			}
		}
	}

	return result
}
