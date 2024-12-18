package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "18"

const gridSize = 71
const byteLimit = 1024

func main() {
	runBoth()
	//runGold()
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grid := util.NewGridEmpty()
	gridRow := strings.Repeat(".", gridSize)
	for _ = range gridSize {
		grid.AddRow(gridRow)
	}
	start := util.Point{0, 0}
	end := util.Point{gridSize - 1, gridSize - 1}
	var silverAnswer int
	var goldAnswer string
	for i, line := range lines {
		coords := strings.Split(line, ",")
		c, _ := strconv.Atoi(coords[0])
		r, _ := strconv.Atoi(coords[1])
		grid.SetRune(r, c, '#')
		if i == byteLimit-1 {
			_, path := grid.BFS(start, end)
			silverAnswer = len(path) - 1
			//grid.PrintDebugWithDots(util.ConvertSliceToMap(path))
		} else if i >= byteLimit {
			found, _ := grid.BFS(start, end)
			if !found {
				goldAnswer = line
				break
			}
		}
	}
	fmt.Println(silverAnswer)
	fmt.Println(goldAnswer)
}
