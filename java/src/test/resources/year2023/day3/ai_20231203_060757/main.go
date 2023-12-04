
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var engineSchematic [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		engineSchematic = append(engineSchematic, strings.Split(line, ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	partSum := sumOfParts(engineSchematic)
	fmt.Println(partSum)
}

func sumOfParts(engineSchematic [][]string) int {
	totalSum := 0
	symbol := func(char string) bool {
		return char != "." && (char < "0" || char > "9")
	}

	for i, row := range engineSchematic {
		for j, char := range row {
			if symbol(char) {
				// Check all adjacent positions.
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue // Skip self.
						}
						newI, newJ := i+x, j+y
						if newI >= 0 && newI < len(engineSchematic) &&
							newJ >= 0 && newJ < len(engineSchematic[newI]) {
							adjacentChar := engineSchematic[newI][newJ]
							// If adjacent character is a digit, add it to the sum.
							if digit, err := strconv.Atoi(adjacentChar); err == nil {
								totalSum += digit
								// Clear the number to avoid double counting it.
								engineSchematic[newI][newJ] = "."
							}
						}
					}
				}
			}
		}
	}
	return totalSum
}
