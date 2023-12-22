package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "13"

func main() {
	runAny(0)
	runAny(1)
}

func runAny(diff int) {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grids := [][][]rune{}
	grid := [][]rune{}
	for _, line := range lines {
		if line == "" {
			grids = append(grids, grid)
			grid = [][]rune{}
		} else {
			grid = append(grid, []rune(line))
		}
	}
	if len(grid) > 0 {
		grids = append(grids, grid)
	}
	ans := 0
	for _, grid := range grids {
		ans += findReflection(grid, diff)
	}
	println(ans)
}

func findReflection(grid [][]rune, diffExpect int) int {
	rows := len(grid)
	cols := len(grid[0])
	rr := 0
	cr := 0
	for cur := 0; cur < rows-1; cur++ {
		diff := 0
		for r := 0; r < rows; r++ {
			left := cur - r
			right := cur + r + 1
			if 0 <= left && left < right && right < rows {
				diff += stringsDiff(string(grid[left]), string(grid[right]))

			}
		}
		if diff == diffExpect {
			rr = rr + cur + 1
		}
	}

	for cur := 0; cur < cols-1; cur++ {
		diff := 0
		for c := 0; c < cols; c++ {
			left := cur - c
			right := cur + c + 1
			if 0 <= left && left < right && right < cols {
				diff += stringsDiff(getColString(grid, left), getColString(grid, right))
			}
		}
		if diff == diffExpect {
			cr = cr + cur + 1
		}
	}
	return 100*rr + cr
}

func stringsDiff(line1 string, line2 string) int {
	ans := 0
	for i, _ := range line1 {
		if line1[i] != line2[i] {
			ans++
		}
	}
	return ans
}
func getColString(grid [][]rune, col int) string {
	rows := len(grid)
	row := []rune{}
	for r := 0; r < rows; r++ {
		row = append(row, grid[r][col])
	}
	return string(row)
}

func printGrids(grids [][][]rune) {
	for i, grid := range grids {
		fmt.Println(i)
		printGrid(grid)
	}
}

func printGrid(grid [][]rune) {
	for _, line := range grid {
		println(string(line))
	}
}
