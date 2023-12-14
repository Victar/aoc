
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	// Variables to hold the summary computation
	verticalReflectionSum := 0
	horizontalReflectionSum := 0
	pattern := []string{}

	// Process each pattern in the input file
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" { // An empty line signals the end of a pattern
			vert, horiz := findReflection(pattern)
			verticalReflectionSum += vert
			horizontalReflectionSum += horiz
			pattern = []string{} // Reset the pattern
		} else {
			pattern = append(pattern, line)
		}
	}

	if len(pattern) > 0 { // Process the last pattern if the file didn't end with an empty line
		vert, horiz := findReflection(pattern)
		verticalReflectionSum += vert
		horizontalReflectionSum += horiz
	}

	// Adding 100 multiplied by the number of rows above each horizontal line of reflection
	horizontalReflectionSum *= 100

	// Calculate the final summary number
	finalSummary := verticalReflectionSum + horizontalReflectionSum
	fmt.Println(finalSummary)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// findReflection finds the position of the line of reflection for a given pattern.
// It returns the number of columns to the left of the line of vertical reflection and
// the number of rows above the line of horizontal reflection.
func findReflection(pattern []string) (int, int) {
	length := len(pattern[0])
	rows := len(pattern)

	// Check for vertical reflection
	for col := 0; col < length/2; col++ {
		mirrorCol := length - 1 - col
		matched := true
		for _, line := range pattern {
			if line[col] != line[mirrorCol] {
				matched = false
				break
			}
		}
		if matched {
			return col, 0
		}
	}

	// Check for horizontal reflection
	for row := 0; row < rows/2; row++ {
		mirrorRow := rows - 1 - row
		if strings.Compare(pattern[row], pattern[mirrorRow]) == 0 {
			return 0, row
		}
	}

	return 0, 0
}
