package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid, startX, startY := readInput("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day10/input.txt")
	visited := make(map[int]map[int]int)
	visited[startY] = map[int]int{startX: 0}
	maxDist := walk(grid, startX, startY, startX, startY, 0, visited)

	// Printing the answer with newline to match the requirements.
	fmt.Println(maxDist)
}

func readInput(filename string) ([][]rune, int, int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	var startX, startY int
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := strings.TrimSpace(scanner.Text())
		row := make([]rune, len(line))
		for x, r := range line {
			row[x] = r
			if r == 'S' {
				startX, startY = x, y
			}
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid, startX, startY
}

func walk(grid [][]rune, prevX, prevY, x, y, dist int, visited map[int]map[int]int) int {
	// If already visited with a shorter distance, skip this path
	if v, found := visited[y][x]; found && v <= dist {
		return 0
	}

	// Mark as visited with current distance
	if _, found := visited[y]; !found {
		visited[y] = make(map[int]int)
	}
	visited[y][x] = dist

	// Define the possible directions to move
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	maxDist := dist

	// Try moving in each direction
	char := grid[y][x]
	for _, dir := range dirs {
		newX, newY := x+dir[0], y+dir[1]
		if validMove(grid, x, y, prevX, prevY, newX, newY, char) {
			newDist := walk(grid, x, y, newX, newY, dist+1, visited)
			if newDist > maxDist {
				maxDist = newDist
			}
		}
	}

	return maxDist
}

func validMove(grid [][]rune, x, y, prevX, prevY, newX, newY int, char rune) bool {
	if newX < 0 || newY < 0 || newX >= len(grid[0]) || newY >= len(grid) {
		return false
	}
	if newX == prevX && newY == prevY {
		return false
	}
	newChar := grid[newY][newX]
	switch char {
	case '|', 'S':
		if newChar != '|' && newChar != '-' && !isCorner(newChar) {
			return false
		}
	case '-':
		if newChar != '-' && newChar != '|' && !isCorner(newChar) {
			return false
		}
	case 'L', 'J', '7', 'F':
		if !isCorrectCorner(char, newX-x, newY-y) {
			return false
		}
	}
	return true
}

func isCorner(char rune) bool {
	return char == 'L' || char == 'J' || char == '7' || char == 'F'
}

func isCorrectCorner(corner rune, dx, dy int) bool {
	switch corner {
	case 'L':
		return (dx == 1 && dy == 0) || (dx == 0 && dy == -1)
	case 'J':
		return (dx == -1 && dy == 0) || (dx == 0 && dy == -1)
	case '7':
		return (dx == -1 && dy == 0) || (dx == 0 && dy == 1)
	case 'F':
		return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
	}
	return false
}

//func main() {
//	grid, startX, startY := readInput("input.txt")
//	visited := make(map[int]map[int]int)
//	visited[startY] = map[int]int{startX: 0}
//	maxDist := walk(grid, startX, startY, startX, startY, 0, visited)
//
//	// Print the answer only
//	fmt.Print(maxDist)
//}
