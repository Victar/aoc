package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadRocks(plate [][]rune) int {
	rows := len(plate)
	loadSum := 0

	for i, row := range plate {
		for _, rock := range row {
			if rock == 'O' {
				loadSum += rows - i // Load is based on distance from south edge
			}
		}
	}

	return loadSum
}

func tiltNorth(plate [][]rune) {
	for i := range plate {
		for j := range plate[i] {
			if i > 0 && plate[i][j] == 'O' && plate[i-1][j] == '.' {
				plate[i][j], plate[i-1][j] = '.', 'O'
				// Once a rock moves, start over from the opposite end to account for the new empty space
				i = len(plate)
				break
			}
		}
	}
}

func readInput(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var surface [][]rune
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		surface = append(surface, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return surface
}

func main() {
	// Read the input from file
	platform := readInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day14/sample.txt")

	// Tilt the platform north until all rounded rocks are as far north as possible
	tiltNorth(platform)

	// Calculate load on the north support beams
	load := loadRocks(platform)

	// Print the solution to standard output
	fmt.Print(load)
}
