package main

import (
	"adventofcode/util"
)

var DAY = "6"

type Position struct {
	PointCur     util.Point
	DirectionCur util.Direction
}

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
		grid.AddRaw(line)
	}
	_, silverAns := isGridCycles(grid)
	println(silverAns)

	sizeR, sizeC := grid.RowColLength()
	goldAns := 0
	for r := 0; r < sizeR; r++ {
		for c := 0; c < sizeC; c++ {
			gridCopy := grid.Copy()
			gridCopy.SetRune(r, c, '#')
			cycled, _ := isGridCycles(gridCopy)
			if cycled {
				goldAns++
			}
		}
	}
	println(goldAns)
}

func isGridCycles(grid *util.Grid) (bool, int) {
	startPoint, curDirection := findStartPosition(grid)
	visitedPositions := make(map[Position]bool)
	visitedPoints := make(map[util.Point]bool)
	visitedPoints[startPoint] = true
	curPoint := startPoint
	for {
		curPosition := Position{curPoint, curDirection}
		if visitedPositions[curPosition] {
			return true, len(visitedPoints)
		} else {
			visitedPositions[curPosition] = true
		}
		nextPoint := curPoint.AddDirection(curDirection)
		if !grid.IsValidPoint(nextPoint) {
			break
		}
		if grid.AtPoint(nextPoint) != '#' {
			curPoint = nextPoint
			visitedPoints[curPoint] = true
		} else {
			curDirection = turnRight(curDirection)
		}
	}
	return false, len(visitedPoints)
}

func findStartPosition(grid *util.Grid) (util.Point, util.Direction) {
	for r, row := range grid.Grid {
		for c, val := range row {
			if val == '^' {
				return util.NewPoint(r, c), util.UP
			}
		}
	}
	return util.NewPoint(-1, -1), util.UP
}

func turnRight(currentDir util.Direction) util.Direction {
	switch currentDir {
	case util.UP:
		return util.RIGHT
	case util.RIGHT:
		return util.DOWN
	case util.DOWN:
		return util.LEFT
	case util.LEFT:
		return util.UP
	default:
		return currentDir
	}
}
