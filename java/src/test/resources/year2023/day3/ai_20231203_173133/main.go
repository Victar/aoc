
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var schematic [][]string
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, strings.Split(line, ""))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	partNumbersSum, gearRatiosSum := processSchematic(schematic)
	fmt.Println(partNumbersSum + gearRatiosSum)
}

func processSchematic(schematic [][]string) (int, int) {
	symbolRegex := regexp.MustCompile(`[*#$+]`)
	partNumbersSum := 0
	gearRatiosSum := 0
	for i := range schematic {
		for j, char := range schematic[i] {
			if char == "." || symbolRegex.MatchString(char) {
				continue
			}
			// Found a number, check if it's adjacent to a symbol
			if adjacentToSymbol(schematic, i, j, symbolRegex) {
				partNumber, _ := strconv.Atoi(char)
				partNumbersSum += partNumber
			}
			// Special handling for gears ('*')
			if char == "*" {
				ratio, isGear := getGearRatio(schematic, i, j)
				if isGear {
					gearRatiosSum += ratio
				}
			}
		}
	}
	return partNumbersSum, gearRatiosSum
}

func adjacentToSymbol(schematic [][]string, row int, col int, regex *regexp.Regexp) bool {
	directions := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range directions {
		nx, ny := row+dir.x, col+dir.y
		if nx >= 0 && ny >= 0 && nx < len(schematic) && ny < len(schematic[nx]) && regex.MatchString(schematic[nx][ny]) {
			return true
		}
	}
	return false
}

func getGearRatio(schematic [][]string, row int, col int) (int, bool) {
	numbers := []int{}
	directions := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range directions {
		nx, ny := row+dir.x, col+dir.y
		if nx >= 0 && ny >= 0 && nx < len(schematic) && ny < len(schematic[nx]) {
			char := schematic[nx][ny]
			if char != "." && char != "*" && char != "#" && char != "$" && char != "+" {
				number, err := strconv.Atoi(char)
				if err != nil {
					continue // Ignore non-numeric adjacent items
				}
				numbers = append(numbers, number)
			}
		}
	}
	if len(numbers) == 2 {
		return numbers[0] * numbers[1], true
	}
	return 0, false
}
