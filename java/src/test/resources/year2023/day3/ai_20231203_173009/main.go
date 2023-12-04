
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open and read the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var schematic [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	symbols := map[rune]bool{
		'*': true,
		'+': true,
		'#': true,
		'$': true,
	}

	partSum := 0
	gearRatiosSum := 0

	// Find the sums of part numbers and gear ratios
	for y, row := range schematic {
		for x, char := range row {
			if char >= '0' && char <= '9' {
				// Check for symbols adjacent to the number
				if isAdjacentToSymbol(schematic, x, y, symbols) {
					partNumber := int(char - '0')

					// Getting full part number (which might be multiple digits)
					partNumber, x = getFullPartNumber(schematic, x, y)

					partSum += partNumber
				}
			} else if char == '*' {
				// Check if it's a gear
				ratio, isGear := calculateGearRatio(schematic, x, y)
				if isGear {
					gearRatiosSum += ratio
				}
			}
		}
	}

	fmt.Println(partSum + gearRatiosSum)
}

func isAdjacentToSymbol(sch [][]rune, x, y int, symbols map[rune]bool) bool {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newX, newY := x+i, y+j
			if newX >= 0 && newX < len(sch[0]) && newY >= 0 && newY < len(sch) {
				if _, ok := symbols[sch[newY][newX]]; ok {
					return true
				}
			}
		}
	}
	return false
}

func calculateGearRatio(sch [][]rune, x, y int) (int, bool) {
	partNumbers := []int{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newX, newY := x+i, y+j
			if newX >= 0 && newX < len(sch[0]) && newY >= 0 && newY < len(sch) {
				if sch[newY][newX] >= '0' && sch[newY][newX] <= '9' {
					partNumber, _ := getFullPartNumber(sch, newX, newY)
					partNumbers = append(partNumbers, partNumber)
				}
			}
		}
	}

	if len(partNumbers) == 2 {
		return partNumbers[0] * partNumbers[1], true
	}

	return 0, false
}

func getFullPartNumber(sch [][]rune, x, y int) (int, int) {
	// Get a full number that might span multiple digits to the left
	numberStr := string(sch[y][x])
	for i := x - 1; i >= 0; i-- {
		if sch[y][i] >= '0' && sch[y][i] <= '9' {
			numberStr = string(sch[y][i]) + numberStr
		} else {
			break
		}
	}
	// Convert string to integer part number
	partNumber, _ := strconv.Atoi(numberStr)
	return partNumber, x + len(numberStr) - 1
}
