package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "12"

func main() {
	runBoth()
}

func runBoth() {
	lines, _ := util.ReadFile("year2024/day" + DAY + "/input.txt")
	var grid = util.NewGridEmpty()
	for _, line := range lines {
		grid.AddRow(line)
	}
	totalPriceSilver := 0
	totalPriceGold := 0
	rSize, cSize := grid.RowColSize()
	visited := make(map[util.Point]bool)
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			cell := grid.At(r, c)
			point := util.Point{r, c}
			if !visited[point] {
				area, perimeter, areaPoints := dfs(grid, point, visited, cell)
				sides := sidesCount(areaPoints)
				totalPriceSilver += area * perimeter
				totalPriceGold += area * sides
			}
		}
	}
	fmt.Println(totalPriceSilver)
	fmt.Println(totalPriceGold)
}

func sidesCount(area map[util.Point]bool) int {
	corners := 0
	for curPoint, _ := range area {
		isCurPointUp := area[curPoint.AddPoint(util.Directions[util.UP])]
		isCurPointRightUp := area[curPoint.AddPoint(util.Directions[util.RIGHT_UP])]
		isCurPointRight := area[curPoint.AddPoint(util.Directions[util.RIGHT])]
		isCurPointRightDown := area[curPoint.AddPoint(util.Directions[util.RIGHT_DOWN])]
		isCurPointDown := area[curPoint.AddPoint(util.Directions[util.DOWN])]
		isCurPointLeftDown := area[curPoint.AddPoint(util.Directions[util.LEFT_DOWN])]
		isCurPointLeft := area[curPoint.AddPoint(util.Directions[util.LEFT])]
		isCurPointLeftUp := area[curPoint.AddPoint(util.Directions[util.LEFT_UP])]
		//fmt.Println(curPoint, isCurPointUp, isCurPointRightUp, isCurPointRight, isCurPointRightDown, isCurPointDown, isCurPointLeftDown, isCurPointLeft, isCurPointLeftUp)
		// isCurPointLeftUp   isCurPointUp   isCurPointRightUp
		// isCurPointLeft     curPoint       isCurPointRight
		// isCurPointLeftDown isCurPointDown isCurPointRightDown
		//

		if !isCurPointLeft && !isCurPointUp {
			// 0 x 0
			// x @ 0
			// 0 0 0
			corners++
		}
		if !isCurPointLeft && !isCurPointDown {
			// 0 0 0
			// x @ 0
			// 0 x 0
			corners++
		}
		if !isCurPointRight && !isCurPointUp {
			// 0 x 0
			// 0 @ x
			// 0 0 0
			corners++
		}
		if !isCurPointRight && !isCurPointDown {
			// 0 0 0
			// 0 @ x
			// 0 x 0
			corners++
		}
		if !isCurPointRightUp && isCurPointUp && isCurPointRight {
			// 0 @ x
			// 0 @ @
			// 0 0 0
			corners++
		}
		if !isCurPointLeftUp && isCurPointUp && isCurPointLeft {
			// x @ 0
			// @ @ 0
			// 0 0 0
			corners++
		}
		if !isCurPointRightDown && isCurPointDown && isCurPointRight {
			// 0 0 0
			// 0 @ @
			// 0 @ x
			corners++
		}
		if !isCurPointLeftDown && isCurPointDown && isCurPointLeft {
			// 0 0 0
			// @ @ 0
			// x @ 0
			corners++
		}
	}
	return corners
}

func dfs(grid *util.Grid, startPoint util.Point, visited map[util.Point]bool, regionType rune) (int, int, map[util.Point]bool) {
	directions := []util.Point{util.Point{-1, 0}, util.Point{1, 0}, util.Point{0, -1}, util.Point{0, 1}}
	stack := []util.Point{startPoint}
	area, perimeter := 0, 0
	areaPoints := map[util.Point]bool{}
	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[point] {
			continue
		}
		visited[point] = true
		areaPoints[point] = true
		area++
		localPerimeter := 4
		for _, dir := range directions {
			nextPoint := point.AddPoint(dir)
			if grid.IsValidPoint(nextPoint) {
				if grid.AtPoint(nextPoint) == regionType {
					stack = append(stack, nextPoint)
					localPerimeter--
				}
			}
		}
		perimeter += localPerimeter
	}
	return area, perimeter, areaPoints
}
