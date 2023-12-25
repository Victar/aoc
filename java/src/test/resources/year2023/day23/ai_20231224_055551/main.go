
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Map [][]
type Visited [][]
type Hike struct {
	Trail Map
	Visited Visited
	MaxLength int
}

func (h *Hike) FindLongestHike(row, col, length int, slopes bool) {
	if h.Visited[row][col] {
		return
	}

	h.Visited[row][col] = true
	defer func() { h.Visited[row][col] = false }() // backtrack

	if length > h.MaxLength {
		h.MaxLength = length
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // UP, DOWN, LEFT, RIGHT
	for _, dir := range directions {
		newRow, newCol := row+dir[0], col+dir[1]
		if newRow >= 0 && newRow < len(h.Trail) && newCol >= 0 && newCol < len(h.Trail[0]) {
			tile := h.Trail[newRow][newCol]
			if tile == '.' || (slopes && (tile == '>' || tile == '<' || tile == 'v' || tile == '^')) {
				h.FindLongestHike(newRow, newCol, length+1, slopes)
			}
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read from input.txt: %v", err)
	}

	// Convert data to map and process the data here (omitted for brevity)

	// Example processing, should be replaced with actual parsing logic
	trailMap := Map{ /* ... fill in the map ... */ }
	visitedMap := Visited{ /* ... same size as trailMap, initialized to false ... */ }

	hike := Hike{
		Trail: trailMap,
		Visited: visitedMap,
		MaxLength: 0,
	}

	startRow, startCol := /* ... find starting positions ... */

	// Part 1
	hike.FindLongestHike(startRow, startCol, 0, true)
	fmt.Println(hike.MaxLength)

	// Part 2
	hike.MaxLength = 0 // Reset for part 2
	hike.FindLongestHike(startRow, startCol, 0, false)
	fmt.Println(hike.MaxLength)
}
