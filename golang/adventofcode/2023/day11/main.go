package main

import (
	"adventofcode/util"
	"fmt"
	"math"
)

var DAY = "11"

func main() {
	runAny(2)
	runAny(1000000)
}

func runAny(n int) {
	grid, err := util.ReadInput("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	galaxiesOrigin := findGalaxies(grid)
	colExpansion, rowExpansion := findExpandUniverse(grid)
	fmt.Println(findSumManhattanDistance(galaxiesOrigin, colExpansion, rowExpansion, n))
}

func findSumManhattanDistance(galaxies []coords, colExpansion []bool, rowExpansion []bool, n int) int {
	sumDistance := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			distance := manhattanDistance(galaxies[i], galaxies[j], colExpansion, rowExpansion, n)
			sumDistance += distance
		}
	}
	return sumDistance
}

func manhattanDistance(a, b coords, colExpansion []bool, rowExpansion []bool, n int) int {
	addX, addY := 0, 0
	n = n - 1
	if a.x < b.x {
		for ix := a.x; ix < b.x; ix++ {
			if !colExpansion[ix] {
				addX += n
			}
		}
	} else {
		for ix := b.x; ix < a.x; ix++ {
			if !colExpansion[ix] {
				addX += n
			}
		}
	}
	if a.y < b.y {
		for iy := a.y; iy < b.y; iy++ {
			if !rowExpansion[iy] {
				addY += n
			}
		}
	} else {
		for iy := b.y; iy < a.y; iy++ {
			if !rowExpansion[iy] {
				addY += n
			}
		}
	}
	return int(math.Abs(float64(a.x-b.x))+math.Abs(float64(a.y-b.y))) + addX + addY
}

func findExpandUniverse(grid [][]rune) ([]bool, []bool) {
	colExpansion := make([]bool, len(grid[0]))
	rowExpansion := make([]bool, len(grid))
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				rowExpansion[y] = true
				colExpansion[x] = true
			}
		}
	}
	return colExpansion, rowExpansion
}

func findGalaxies(grid [][]rune) []coords {
	var galaxies []coords
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				galaxies = append(galaxies, coords{x: x, y: y})
			}
		}
	}
	return galaxies
}

type coords struct {
	x, y int
}
