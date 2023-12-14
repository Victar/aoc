
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var platform []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		platform = append(platform, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// We need to tilt north, this means we check where `O`s can move to the first row and stop moving if `#` is encountered.
	// When `O` moves north, we calculate its load (distance from the bottom), and finally, sum up the total load.

	rows := len(platform)
	cols := len(platform[0])
	totalLoad := 0

	for c := 0; c < cols; c++ {
		canMove := true
		for r := 0; r < rows && canMove; r++ {
			if strings.HasSuffix(platform[r], " ") {
				// Edge-case handling for lines shorter than others due to trailing space removal in input
				platform[r] += " "
			}
			switch platform[r][c] {
			case '#':
				canMove = false
			case 'O':
				if canMove { // Only the top-most `O` in a column contributes to the load
					totalLoad += rows - r
					canMove = false
				}
			}
		}
	}

	fmt.Println(totalLoad)
}
