
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInput(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" { // skip empty lines
			continue
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}

func findReflection(lines []string) int {
	height := len(lines)
	width := len(lines[0])

	// Check for vertical reflection line
	for col := 0; col < width/2; col++ {
		mirror := true
		for row := 0; row < height; row++ {
			if lines[row][col] != lines[row][width-col-1] {
				mirror = false
				break
			}
		}
		if mirror {
			return col
		}
	}

	// Check for horizontal reflection line
	for row := 0; row < height/2; row++ {
		mirror := true
		for col := 0; col < width; col++ {
			if lines[row][col] != lines[height-row-1][col] {
				mirror = false
				break
			}
		}
		if mirror {
			return 100 * row
		}
	}

	return 0
}

func main() {
	lines, err := readInput("input.txt")
	check(err)

	total := 0
	currentPattern := []string{}

	for _, line := range lines {
		if strings.Contains(line, ".") || strings.Contains(line, "#") {
			currentPattern = append(currentPattern, line)
		} else {
			// Finished reading a pattern or reached an empty line.
			if len(currentPattern) != 0 {
				total += findReflection(currentPattern)
				currentPattern = []string{} // reset the current pattern
			}
		}
	}
	if len(currentPattern) != 0 {
		// If there's a pattern at the end of the input without a trailing empty line.
		total += findReflection(currentPattern)
	}

	fmt.Println(total)
}
