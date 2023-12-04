
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input.txt"
	schematic, err := readSchematic(inputFile)
	if err != nil {
		fmt.Println("Error reading schematic:", err)
		return
	}

	sum, err := sumOfPartNumbers(schematic)
	if err != nil {
		fmt.Println("Error calculating sum:", err)
		return
	}

	fmt.Println(sum)
}

func readSchematic(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var schematic [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			row := []rune(line)
			schematic = append(schematic, row)
		}
	}
	return schematic, scanner.Err()
}

func sumOfPartNumbers(schematic [][]rune) (int, error) {
	height := len(schematic)
	width := len(schematic[0])
	sum := 0

	for y, row := range schematic {
		for x, cell := range row {
			if cell >= '0' && cell <= '9' {
				// Check surrounding cells
				if hasSymbolAdjacent(schematic, x, y, width, height) {
					// Try to parse integer if cell is 'part number'
					val, err := strconv.Atoi(string(cell))
					if err != nil {
						return 0, err
					}
					sum += val
				}
			}
		}
	}

	return sum, nil
}

func hasSymbolAdjacent(schematic [][]rune, x, y, width, height int) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue // skip the current cell
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < width && ny >= 0 && ny < height {
				adjCell := schematic[ny][nx]
				if adjCell != '.' && adjCell < '0' || adjCell > '9' {
					return true
				}
			}
		}
	}
	return false
}
