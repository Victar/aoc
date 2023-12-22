
package main

import (
	"bufio"
	"fmt"
	"os"
)

const totalSteps = 64

var (
	dx               = []int{0, 0, -1, 1}
	dy               = []int{-1, 1, 0, 0}
	uniqueGardenPlots = map[int]bool{}
)

func walk(garden [][]rune, x, y, steps int) {
	if steps == totalSteps {
		uniqueGardenPlots[x*len(garden[0])+y] = true
		return
	}

	garden[x][y] = '#'
	for i := 0; i < 4; i++ {
		nx, ny := x+dx[i], y+dy[i]
		if nx >= 0 && nx < len(garden) && ny >= 0 && ny < len(garden[0]) && garden[nx][ny] == '.' {
			walk(garden, nx, ny, steps+1)
		}
	}
	garden[x][y] = '.'
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var garden [][]rune
	var startX, startY int
	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := make([]rune, len(line))
		for x, char := range line {
			if char == 'S' {
				startX, startY = y, x
			}
			row[x] = char
		}
		garden = append(garden, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	walk(garden, startX, startY, 0)
	fmt.Println(len(uniqueGardenPlots))
}

