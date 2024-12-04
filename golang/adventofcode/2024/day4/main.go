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
		grid.AddRaw(line)
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
	for _, dir := range util.DIRECTIONS_ALL {
		if grid.IsValid(r+3*dir.R, c+3*dir.C) {
			if grid.At(r+dir.R, c+dir.C) == 'M' && grid.At(r+2*dir.R, c+2*dir.C) == 'A' && grid.At(r+3*dir.R, c+3*dir.C) == 'S' {
				count++
			}
		}
	}
	return count
}

func validGold(grid *util.Grid, r, c int) int {
	count := 0
	//M.S
	//.A.
	//M.S
	if grid.IsValid(r-1, c-1) && grid.IsValid(r+1, c+1) {
		if grid.At(r-1, c-1) == 'M' && grid.At(r-1, c+1) == 'M' && grid.At(r+1, c-1) == 'S' && grid.At(r+1, c+1) == 'S' {
			count++
		}
		if grid.At(r-1, c-1) == 'M' && grid.At(r-1, c+1) == 'S' && grid.At(r+1, c-1) == 'M' && grid.At(r+1, c+1) == 'S' {
			count++
		}
		if grid.At(r-1, c-1) == 'S' && grid.At(r-1, c+1) == 'S' && grid.At(r+1, c-1) == 'M' && grid.At(r+1, c+1) == 'M' {
			count++
		}
		if grid.At(r-1, c-1) == 'S' && grid.At(r-1, c+1) == 'M' && grid.At(r+1, c-1) == 'S' && grid.At(r+1, c+1) == 'M' {
			count++
		}
	}

	return count
}
