
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Position struct defines x and y coordinates
type Position struct {
	x, y int
}

// Deltas are the coordinate offsets for surrounding cells
var deltas = []Position{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

// IsSymbol checks if a character is a symbol
func IsSymbol(r rune) bool {
	return r == '*' || r == '+' || r == '#' || r == '$'
}

// ReadEngineSchematic reads the schematic from a file and returns it as a slice of strings
func ReadEngineSchematic(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// CalculateSumOfParts extracts part numbers and calculates the sum
func CalculateSumOfParts(schematic []string) (int, int) {
	height := len(schematic)
	width := len(schematic[0])
	totalPartsSum := 0
	gearRatiosSum := 0

	for y, line := range schematic {
		for x, char := range line {
			if IsSymbol(char) {
				for _, d := range deltas {
					nx, ny := x+d.x, y+d.y
					if nx >= 0 && nx < width && ny >= 0 && ny < height && schematic[ny][nx] >= '0' && schematic[ny][nx] <= '9' {
						numStr := string(schematic[ny][nx])
						num, _ := strconv.Atoi(numStr)
						totalPartsSum += num
					}
				}
			} else if char == '*' {
				// Count part numbers and store them for gear ratio calculation
				partNumbers := []int{}
				for _, d := range deltas {
					nx, ny := x+d.x, y+d.y
					if nx >= 0 && nx < width && ny >= 0 && ny < height && schematic[ny][nx] >= '0' && schematic[ny][nx] <= '9' {
						numStr := string(schematic[ny][nx])
						num, _ := strconv.Atoi(numStr)
						partNumbers = append(partNumbers, num)
					}
				}

				// Check if the gear has exactly two numbers
				if len(partNumbers) == 2 {
					gearRatiosSum += partNumbers[0] * partNumbers[1]
				}
			}
		}
	}

	return totalPartsSum, gearRatiosSum
}

func main() {
	schematic, err := ReadEngineSchematic("input.txt")
	if err != nil {
		fmt.Println("Error reading schematic:", err)
		return
	}

	part1, part2 := CalculateSumOfParts(schematic)

	// Print result for Part One
	fmt.Println(part1)

	// Print result for Part Two
	fmt.Println(part2)
}
