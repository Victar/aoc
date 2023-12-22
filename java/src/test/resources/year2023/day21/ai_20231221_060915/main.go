
package main

import (
	"bufio"
	"fmt"
	"os"
)

var steps = 64
var directions = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
var visited = make(map[[2]int]int)
var start [2]int

func readInput() ([][]byte, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var mapData [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mapData = append(mapData, []byte(scanner.Text()))
	}
	return mapData, scanner.Err()
}

func findStartingPosition(mapData [][]byte) (int, int) {
	for y, row := range mapData {
		for x, cell := range row {
			if cell == 'S' {
				return x, y
			}
		}
	}
	return -1, -1 // Ideally this should never happen
}

func countReachablePlots(mapData [][]byte, x, y, curStep int) int {
	if curStep == steps {
		if _, found := visited[[2]int{x, y}]; !found {
			visited[[2]int{x, y}] = curStep
			return 1
		}
		return 0
	}
	visited[[2]int{x, y}] = curStep
	count := 0
	for _, dir := range directions {
		newX, newY := x+dir[0], y+dir[1]
		if newY >= 0 && newY < len(mapData) && newX >= 0 && newX < len(mapData[newY]) && mapData[newY][newX] == '.' && visited[[2]int{newX, newY}] != curStep {
			count += countReachablePlots(mapData, newX, newY, curStep+1)
		}
	}
	return count
}

func main() {
	mapData, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	startX, startY := findStartingPosition(mapData)
	answer := countReachablePlots(mapData, startX, startY, 0)
	fmt.Println(answer)
}
