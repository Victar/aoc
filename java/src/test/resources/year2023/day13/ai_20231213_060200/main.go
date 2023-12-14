
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func findReflection(pattern []string) int {
	nRows := len(pattern)
	nCols := len(pattern[0])
	for col := 0; col < nCols/2; col++ {
		isReflection := true
		for row := 0; row < nRows; row++ {
			if pattern[row][col] != pattern[row][nCols-col-1] {
				isReflection = false
				break
			}
		}
		if isReflection {
			return col // Vertical line reflection
		}
	}
	for row := 0; row < nRows/2; row++ {
		isReflection := true
		for col := 0; col < nCols; col++ {
			if pattern[row][col] != pattern[nRows-row-1][col] {
				isReflection = false
				break
			}
		}
		if isReflection {
			return 100 * row // Horizontal line reflection
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pattern := []string{}
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(pattern) > 0 {
				sum += findReflection(pattern)
				pattern = []string{}
			}
		} else {
			pattern = append(pattern, line)
		}
	}
	if len(pattern) > 0 {
		sum += findReflection(pattern)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Println(sum)
}
