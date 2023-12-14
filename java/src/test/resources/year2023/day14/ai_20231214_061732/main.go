
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Failed to open input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the grid and count the rocks that will move north
	var totalLoad int
	rowNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if char == 'O' {
				totalLoad += rowNumber
			}
		}
		rowNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	fmt.Println(totalLoad)
}
