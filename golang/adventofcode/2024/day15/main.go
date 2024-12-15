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
	//runGold()
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

	sizeR, sizeC := grid.RowColSize()
	maxSize := sizeR
	if sizeC > sizeR {
		maxSize = sizeC
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
		gps := box.R()*100 + box.C()
		sumGPS += gps
	}
	fmt.Println(sumGPS)

}

func runGold() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
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
			if robotPos.R() == i && robotPos.C() == j {
				curChar = '@'
			}
			sb.WriteByte(curChar)
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}
