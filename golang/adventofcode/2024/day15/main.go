package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "15"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var grid = util.NewGridEmpty()
	movements := ""
	for _, line := range lines {
		if strings.Contains(line, "#") {
			grid.AddRow(line)
		} else {
			movements += strings.TrimSpace(line)
		}
	}
	var robotPos util.Point
	boxes := map[util.Point]bool{}

	rSize, cSize := grid.RowColSize()
	maxSize := rSize
	if rSize > cSize {
		maxSize = cSize
	}
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			cell := grid.At(r, c)
			if cell == '@' {
				robotPos = util.NewPoint(r, c)
				grid.SetRune(r, c, '.')
			} else if cell == 'O' {
				boxes[util.NewPoint(r, c)] = true
				grid.SetRune(r, c, '.')
			}
		}
	}

	moveMap := map[rune]util.Point{
		'^': util.NewPoint(-1, 0),
		'v': util.NewPoint(1, 0),
		'<': util.NewPoint(0, -1),
		'>': util.NewPoint(0, 1),
	}

	for _, move := range movements {
		delta := moveMap[move]
		nextPos := robotPos.AddPoint(delta)
		if grid.AtPoint(nextPos) == '#' {
			continue
		}
		if boxes[nextPos] {
			for i := 1; i < maxSize; i++ {
				boxNextPos := nextPos.AddPoint(delta.TimesPoint(i))
				if grid.IsValidPoint(boxNextPos) && grid.AtPoint(boxNextPos) != '#' && !boxes[boxNextPos] {
					delete(boxes, nextPos)
					boxes[boxNextPos] = true
					robotPos = nextPos
					break
				} else if !boxes[boxNextPos] {
					break
				}
			}
		} else {
			robotPos = nextPos
		}
	}
	printState(grid, boxes, robotPos)

	sumGPS := 0
	for box := range boxes {
		gps := box.R*100 + box.C
		sumGPS += gps
	}
	fmt.Println(sumGPS)

}
func printState(g *util.Grid, boxes map[util.Point]bool, robotPos util.Point) {
	fmt.Printf("Grid row %d columns %d \n", len(g.Grid), len(g.Grid[0]))
	columns := len(g.Grid[0])
	if columns > 10 {
		firstLine := ""
		for i := 0; i < 10; i++ {
			firstLine += strings.Repeat(strconv.Itoa(i), 10)
		}
		println(strings.Repeat(firstLine, columns/100+1)[:columns])
	}
	println(strings.Repeat("0123456789", columns/10+1)[:columns])
	sb := strings.Builder{}
	for i, line := range g.Grid {
		for j, char := range line {
			curChar := byte(char)
			if boxes[util.NewPoint(i, j)] {
				curChar = 'O'
			}
			if robotPos.R == i && robotPos.C == j {
				curChar = '@'
			}
			sb.WriteByte(curChar)
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func runGold() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var originalGrid = util.NewGridEmpty()
	movements := ""
	for _, line := range lines {
		if strings.Contains(line, "#") {
			originalGrid.AddRow(line)
		} else {
			movements += strings.TrimSpace(line)
		}
	}

	gridBig := scaleUpGrid(originalGrid)
	var robotPos util.Point
	rSize, cSize := gridBig.RowColSize()
	grid, robot := map[util.Point]rune{}, util.Point{}

	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			cell := gridBig.At(r, c)
			if cell == '@' {
				robotPos = util.Point{r, c}
				gridBig.SetRune(r, c, '.')
			}
			grid[util.Point{r, c}] = cell
		}
	}
	delta := map[rune]util.Point{
		'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
		'[': {1, 0}, ']': {-1, 0},
	}

	//var isBox func(leftBorder util.Point) (bool, util.Point)
	//var canMoveBoxes func(robotPos util.Point, delta util.Point, boxesToMove map[util.Point]bool) (bool, map[util.Point]bool)
	//var moveBoxes func(boxesToMove map[util.Point]bool, delta util.Point)
	//
	//isBox = func(leftBorder util.Point) (bool, util.Point) {
	//	if boxes[leftBorder] {
	//		return true, leftBorder
	//	}
	//	rightBorder := leftBorder.AddPoint(util.NewPoint(0, -1))
	//	if boxes[rightBorder] {
	//		return true, rightBorder
	//	}
	//	return false, util.NewPoint(-1, -1)
	//}
	//
	//canMoveBoxes = func(robotPos util.Point, delta util.Point, boxesToMove map[util.Point]bool) (bool, map[util.Point]bool) {
	//	nextPos := robotPos.AddPoint(delta)
	//	nextIsBox, nextBox := isBox(nextPos)
	//	if grid.AtPoint(nextPos) == '#' {
	//		return false, boxesToMove
	//	}
	//	if nextIsBox {
	//		boxesToMove[nextBox] = true
	//		return canMoveBoxes(nextBox, delta, boxesToMove)
	//	}
	//	return true, boxesToMove
	//}
	//
	//moveBoxes = func(boxesToMove map[util.Point]bool, delta util.Point) {
	//	for box := range boxesToMove {
	//		delete(boxes, box)
	//		boxes[box.AddPoint(delta)] = true
	//	}
	//}

loop:
	for _, r := range movements {
		queue, boxes := []util.Point{robotPos}, map[util.Point]rune{}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if _, ok := boxes[p]; ok {
				continue
			}
			boxes[p] = grid[p]

			switch n := p.AddPoint(delta[r]); grid[n] {
			case '#':
				continue loop
			case '[', ']':
				queue = append(queue, n.AddPoint(delta[grid[n]]))
				fallthrough
			case 'O':
				queue = append(queue, n)
			}
		}

		for b := range boxes {
			grid[b] = '.'
		}
		for b := range boxes {
			grid[b.AddPoint(delta[r])] = boxes[b]
		}
		robot = robot.AddPoint(delta[r])
	}

	sumGPS := 0
	for p, r := range grid {
		if r == 'O' || r == '[' {
			sumGPS += 100*p.R + p.C
		}
	}
	fmt.Println(sumGPS)
}

func scaleUpGrid(grid *util.Grid) *util.Grid {
	scaledGrid := util.NewGridEmpty()
	rSize, cSize := grid.RowColSize()
	for r := 0; r < rSize; r++ {
		var scaledRow1 string
		for c := 0; c < cSize; c++ {
			cell := grid.At(r, c)
			switch cell {
			case '#':
				scaledRow1 += "##"
			case '.':
				scaledRow1 += ".."
			case 'O':
				scaledRow1 += "[]"
			case '@':
				scaledRow1 += "@."
			}
		}
		scaledGrid.AddRow(scaledRow1)
	}

	return scaledGrid
}
