package main

import (
	"adventofcode/util"
)

var DAY = "10"

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
	rSize, cSize := grid.RowColSize()
	silverAnswer := 0
	goldAnswer := 0
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			cell := grid.At(r, c)
			if cell == '0' {
				visited := make(map[util.Point]bool)
				goldAnswer += findScore(grid, util.NewPoint(r, c), '0', visited)
				curScore := len(visited)
				silverAnswer += curScore
			}
		}
	}
	println(silverAnswer)
	println(goldAnswer)
}

func findScore(grid *util.Grid, startPoint util.Point, current rune, visited map[util.Point]bool) int {
	if current == '9' {
		visited[startPoint] = true
		return 1
	}
	directions := []util.Point{util.NewPoint(-1, 0), util.NewPoint(1, 0), util.NewPoint(0, -1), util.NewPoint(0, 1)}
	nextCur := current + 1
	ans := 0
	for _, dir := range directions {
		nextPoint := startPoint.AddPoint(dir)
		if grid.IsValidPoint(nextPoint) && grid.AtPoint(nextPoint) == nextCur {
			ans += findScore(grid, nextPoint, nextCur, visited)
		}
	}
	return ans
}
