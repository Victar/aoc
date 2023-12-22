
package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction int

const (
	right direction = iota
	up
	left
	down
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var contraption [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		contraption = append(contraption, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(energizeTiles(contraption))
}

func energizeTiles(contraption [][]rune) int {
	energizedTiles := make(map[int]map[int]bool)
	var energize func(x, y int, dir direction)
	energize = func(x, y int, dir direction) {
		for y >= 0 && x >= 0 && y < len(contraption) && x < len(contraption[y]) {
			if energizedTiles[y] == nil {
				energizedTiles[y] = make(map[int]bool)
			}
			energizedTiles[y][x] = true

			tile := contraption[y][x]
			switch tile {
			case '.':
				// Continue in the same direction
			case '/':
				switch dir {
				case right:
					dir = up
				case up:
					dir = right
				case left:
					dir = down
				case down:
					dir = left
				}
			case '\\':
				switch dir {
				case right:
					dir = down
				case up:
					dir = left
				case left:
					dir = up
				case down:
					dir = right
				}
			case '|', '-':
				if (dir == right && tile == '|') || (dir == up && tile == '-') {
					energize(x, y, down)
					energize(x, y, up)
					return
				} else if (dir == left && tile == '|') || (dir == down && tile == '-') {
					energize(x, y, up)
					energize(x, y, down)
					return
				}
				// Pass through as empty space
			default:
				return
			}

			switch dir {
			case right:
				x++
			case left:
				x--
			case up:
				y--
			case down:
				y++
			}
		}
	}

	energize(0, 0, right)

	count := 0
	for _, row := range energizedTiles {
		count += len(row)
	}
	return count
}
