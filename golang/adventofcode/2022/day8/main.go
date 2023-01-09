package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
)

var DAY = "8"

func main() {
	runSilver()
	//runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2022/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	grid := make([][]int, 0)
	for y := 0; y < len(lines); y++ {
		var row []int
		for x := 0; x < len(lines[y]); x++ {
			cur, _ := strconv.Atoi(string(lines[y][x]))
			row = append(row, cur)
		}
		grid = append(grid, row)
	}
	fmt.Printf("%v", grid)
}

func runGold() {
	lines, err := util.ReadFile("year2022/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
