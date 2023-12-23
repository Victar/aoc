
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Position struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(strings.TrimSpace(string(data)), "\n")

	// Initialize direction vectors for the slopes
	directions := map[rune]Position{
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
		'^': {0, -1},
	}

	maxLength := 0

	var hike func(x, y, length int, visited map[Position]bool)
	hike = func(x, y, length int, visited map[Position]bool) {
		if y == len(input)-1 && input[y][x] == '.' {
			// We've reached the end; update maxLength if we've found a longer path
			if length > maxLength {
				maxLength = length
			}
			return
		}
		for dir, p := range directions {
			newX, newY := x+p.x, y+p.y
			if newX >= 0 && newY >= 0 && newX < len(input[0]) && newY < len(input) {
				pos := Position{newX, newY}
				if (!visited[pos]) && (input[newY][newX] == '.' || input[newY][newX] == dir) {
					newVisited := make(map[Position]bool)
					for k, v := range visited {
						newVisited[k] = v
					}
					newVisited[pos] = true
					hike(newX, newY, length+1, newVisited)
				}
			}
		}
	}

	// Find the starting position
	for sx := 0; sx < len(input[0]); sx++ {
		if input[0][sx] == '.' {
			visited := make(map[Position]bool)
			visited[Position{sx, 0}] = true
			hike(sx, 0, 1, visited) // initialize length as 1 to include the starting position
			break
		}
	}

	fmt.Println(maxLength)
}
