package main

import (
	"adventofcode/util"
)

var DAY = "4"

func main() {
	runBoth()
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var grid = util.NewGridEmpty()
	for _, line := range lines {
		grid.AddRow(line)
	}
	countSilver := 0
	countGold := 0
	for r := range grid.Grid {
		for c := range grid.Grid[r] {
			if grid.At(r, c) == 'X' {
				countSilver += validSilver(grid, r, c)
			}
			if grid.At(r, c) == 'A' {
				countGold += validGold(grid, r, c)
			}
		}
	}
	println(countSilver)
	println(countGold)
}

func validSilver(grid *util.Grid, r, c int) int {
	count := 0
	curPoint := util.Point{r, c}
	for _, dir := range util.DIRECTIONS_ALL {
		if grid.IsValidPoint(curPoint.AddPoint(dir.TimesPoint(3))) {
			if grid.AtPoint(curPoint.AddPoint(dir.TimesPoint(1))) == 'M' && grid.AtPoint(curPoint.AddPoint(dir.TimesPoint(2))) == 'A' && grid.AtPoint(curPoint.AddPoint(dir.TimesPoint(3))) == 'S' {
				count++
			}
		}
	}
	return count
}

func validGold(grid *util.Grid, r, c int) int {
	count := 0
	patterns := [][]rune{{'M', 'M', 'S', 'S'}, {'M', 'S', 'M', 'S'}, {'S', 'S', 'M', 'M'}, {'S', 'M', 'S', 'M'}}
	directions := []util.Direction{util.LEFT_UP, util.LEFT_DOWN, util.RIGHT_UP, util.RIGHT_DOWN}
	curPoint := util.Point{r, c}
	if grid.IsValidPoint(curPoint.AddDirection(util.LEFT_UP)) && grid.IsValidPoint(curPoint.AddDirection(util.RIGHT_DOWN)) {
		for _, pattern := range patterns {
			valid := true
			for i, direction := range directions {
				if grid.AtPoint(curPoint.AddDirection(direction)) != pattern[i] {
					valid = false
				}
			}
			if valid {
				count++
			}
		}
	}
	return count
}
