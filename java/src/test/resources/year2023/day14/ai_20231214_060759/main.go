
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	fmt.Println(calculateTotalLoad(grid))
}

// calculateTotalLoad will simulate the rolling of rocks and calculate the total load on the north beams.
func calculateTotalLoad(grid []string) int {
	height := len(grid)
	load := 0

	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < height; row++ {
			if grid[row][col] == 'O' {
				load += (height - row)
				break
			} else if grid[row][col] == '#' {
				break
			}
		}
	}
	return load
}
