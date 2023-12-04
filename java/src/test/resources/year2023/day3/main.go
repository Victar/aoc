
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " \r\n")
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning input file:", err)
		return
	}

	sumPartNumbers, sumGearRatios := processSchematic(lines)
	fmt.Println("Sum of all part numbers:", sumPartNumbers)
	fmt.Println("Sum of all gear ratios:", sumGearRatios)
}

// processSchematic processes the engine schematic and returns the sum of all part numbers and gear ratios.
func processSchematic(lines []string) (int, int64) {
	var sumPartNumbers int
	var sumGearRatios int64
	for y, line := range lines {
		for x, char := range line {
			if char >= '0' && char <= '9' {
				if isAdjacentToSymbol(x, y, lines) {
					value := int(char - '0')
					sumPartNumbers += value
					gearVal := findAndMultiplyAdjacentNumbers(x, y, lines)
					if gearVal != -1 {
						sumGearRatios += gearVal
					}
				}
			}
		}
	}
	return sumPartNumbers, sumGearRatios
}

// isAdjacentToSymbol checks if a numeric character is adjacent to a symbol.
func isAdjacentToSymbol(x, y int, lines []string) bool {
	adjacentOffsets := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, offset := range adjacentOffsets {
		adjX, adjY := x+offset.dx, y+offset.dy
		if adjY >= 0 && adjY < len(lines) && adjX >= 0 && adjX < len(lines[adjY]) {
			adjChar := lines[adjY][adjX]
			if adjChar != '.' && adjChar != ' ' && (adjChar < '0' || adjChar > '9') {
				return true
			}
		}
	}
	return false
}

// findAndMultiplyAdjacentNumbers calculates the gear ratio if the character at the given coordinate is a gear.
func findAndMultiplyAdjacentNumbers(x, y int, lines []string) int64 {
	if lines[y][x] != '*' {
		return -1
	}

	numbers := make([]int, 0)
	adjacentOffsets := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, offset := range adjacentOffsets {
		adjX, adjY := x+offset.dx, y+offset.dy
		if adjY >= 0 && adjY < len(lines) && adjX >= 0 && adjX < len(lines[adjY]) {
			adjChar := lines[adjY][adjX]
			if adjChar >= '0' && adjChar <= '9' {
				number, _ := strconv.Atoi(string(adjChar))
				numbers = append(numbers, number)
			}
		}
	}

	if len(numbers) == 2 {
		return int64(numbers[0]) * int64(numbers[1])
	}
	return -1
}
