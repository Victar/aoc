
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	symbols := make(map[rune]bool)
	for _, line := range grid {
		for _, r := range line {
			if r != '.' {
				symbols[r] = true
			}
		}
	}

	sum := 0
	for y, line := range grid {
		for x, r := range line {
			if r >= '0' && r <= '9' {
				if hasAdjacentSymbol(x, y, grid, symbols) {
					num, err := strconv.Atoi(string(r))
					if err != nil {
						log.Fatalf("Invalid number detected: %v", err)
					}
					sum += num
				}
			}
		}
	}

	fmt.Println(sum)
}

func hasAdjacentSymbol(x, y int, grid [][]rune, symbols map[rune]bool) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			adjX := x + dx
			adjY := y + dy

			if adjX >= 0 && adjX < len(grid[0]) && adjY >= 0 && adjY < len(grid) {
				if _, ok := symbols[grid[adjY][adjX]]; ok {
					return true
				}
			}
		}
	}
	return false
}
