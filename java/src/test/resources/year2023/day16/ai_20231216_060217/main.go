package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tile rune

const (
	Empty     Tile = '.'
	Mirror1        = '/'
	Mirror2        = '\\'
	Splitter1      = '|'
	Splitter2      = '-'
)

type Direction [2]int

var directions = map[Tile]map[Direction][]Direction{
	Mirror1: {
		{0, 1}:  {{-1, 0}},
		{1, 0}:  {{0, -1}},
		{0, -1}: {{1, 0}},
		{-1, 0}: {{0, 1}},
	},
	Mirror2: {
		{0, 1}:  {{1, 0}},
		{1, 0}:  {{0, 1}},
		{0, -1}: {{-1, 0}},
		{-1, 0}: {{0, -1}},
	},
	Splitter1: {
		{0, 1}:  {{-1, 0}, {1, 0}},
		{0, -1}: {{-1, 0}, {1, 0}},
	},
	Splitter2: {
		{1, 0}:  {{0, -1}, {0, 1}},
		{-1, 0}: {{0, -1}, {0, 1}},
	},
}

func main() {
	grid, err := readInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day16/sample.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	energized := simulateBeam(grid, 0, 0, Direction{0, 1})
	fmt.Println(energized)
}

func simulateBeam(grid [][]Tile, startX, startY int, initialDir Direction) int {
	var beam func(x, y int, dir Direction)
	visited := make(map[int]map[int]bool)

	beam = func(x, y int, dir Direction) {
		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			return // Out of bounds
		}
		if _, ok := visited[x]; !ok {
			visited[x] = make(map[int]bool)
		}
		if visited[x][y] { // Already visited this tile for this direction
			return
		}
		visited[x][y] = true // Mark this tile visited

		tile := grid[x][y]
		nextDirs := directions[tile][dir]
		if len(nextDirs) == 0 { // If no mapping for direction, continue in same direction
			nextDirs = append(nextDirs, dir)
		}

		for _, nextDir := range nextDirs {
			beam(x+nextDir[0], y+nextDir[1], nextDir)
		}
	}

	beam(startX, startY, initialDir)

	count := 0
	for _, row := range visited {
		count += len(row)
	}
	return count
}

func readInput(filename string) ([][]Tile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]Tile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Tile, len(line))
		for i, ch := range line {
			row[i] = Tile(ch)
		}
		grid = append(grid, row)
	}
	return grid, scanner.Err()
}
