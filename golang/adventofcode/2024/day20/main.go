package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "20"

func main() {
	//runSilver() // slow brute-force
	runAny(2)
	runAny(20)

}

func runAny(distance int) {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grid := util.NewGridEmpty()
	for _, line := range lines {
		grid.AddRow(line)
	}
	var start, end util.Point
	sizeR, sizeC := grid.RowColSize()
	for r := 0; r < sizeR; r++ {
		for c := 0; c < sizeC; c++ {
			if grid.At(r, c) == 'S' {
				start = util.Point{r, c}
				grid.SetRune(r, c, '.')
			}
			if grid.At(r, c) == 'E' {
				end = util.Point{r, c}
				grid.SetRune(r, c, '.')
			}
		}
	}

	var directions = []util.Point{util.Directions[util.RIGHT], util.Directions[util.DOWN], util.Directions[util.LEFT], util.Directions[util.UP]}
	curPoint := start
	traceMap := make(map[util.Point]int)
	traceMap[start] = 0
	for curPoint != end {
		for _, direction := range directions {
			nextPoint := curPoint.AddPoint(direction)
			if _, exists := traceMap[nextPoint]; !exists && grid.IsValidPoint(nextPoint) && grid.AtPoint(nextPoint) == '.' {
				traceMap[nextPoint] = traceMap[curPoint] + 1
				curPoint = nextPoint
			}
		}
	}

	count := 0
	for tracePoint := range traceMap {
		cheatPointsMap := cheatPoints(tracePoint, traceMap, distance)
		for cheatPoint := range cheatPointsMap {
			if traceMap[cheatPoint]-traceMap[tracePoint]-tracePoint.Distance(cheatPoint) >= 100 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func cheatPoints(point util.Point, traceMap map[util.Point]int, distance int) map[util.Point]bool {
	cheatPointsMap := make(map[util.Point]bool)
	for r := -distance; r <= distance; r++ {
		for c := -distance; c <= distance; c++ {
			cheatPoint := point.AddPoint(util.Point{r, c})
			if _, exists := traceMap[cheatPoint]; exists && point.Distance(cheatPoint) <= distance {
				cheatPointsMap[cheatPoint] = true
			}
		}
	}
	return cheatPointsMap
}

func runSilver() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var grid = util.NewGridEmpty()
	for _, line := range lines {
		grid.AddRow(line)
	}
	var start, end util.Point

	rSize, cSize := grid.RowColSize()
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			if grid.At(r, c) == 'S' {
				start = util.Point{r, c}
				grid.SetRune(r, c, '.')
			}
			if grid.At(r, c) == 'E' {
				end = util.Point{r, c}
				grid.SetRune(r, c, '.')

			}
		}
	}
	_, original := grid.BFS(start, end)
	orSize := len(original)
	silverAns := 0
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			if grid.At(r, c) == '#' {
				grid.SetRune(r, c, '.')
				_, path := grid.BFS(start, end)
				if len(path) <= orSize-100 {
					silverAns++
				}
				grid.SetRune(r, c, '#')
			}
		}
	}
	fmt.Println(silverAns)
}
