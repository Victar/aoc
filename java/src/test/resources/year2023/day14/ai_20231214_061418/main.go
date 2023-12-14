package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	input, err := readInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day14/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := calculateTotalLoad(input)
	fmt.Println(result)
}

func readInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line))
	}
	return input, scanner.Err()
}

func calculateTotalLoad(dish [][]rune) int {
	rows := len(dish)
	cols := len(dish[0])
	totalLoad := 0

	for col := 0; col < cols; col++ {
		load := 0
		for row := rows - 1; row >= 0; row-- {
			if dish[row][col] == 'O' {
				load += rows - row
				// Move round rock to the topmost position
				dish[row][col] = '.'
				dish[load][col] = 'O'
			} else if dish[row][col] == '#' {
				// Cube rocks don't roll, so move to the first free position found above it
				for r := row; r >= 1; r-- {
					if dish[r-1][col] == '.' {
						dish[r-1][col] = '#'
						dish[r][col] = '.'
					} else {
						break
					}
				}
			}
		}
		totalLoad += load
	}
	return totalLoad
}
