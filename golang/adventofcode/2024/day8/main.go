package main

import (
	"adventofcode/util"
)

var DAY = "8"

func main() {
	runAny(false)
	runAny(true)
}

func runAny(isGold bool) {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var grid = util.NewGridEmpty()
	for _, line := range lines {
		grid.AddRaw(line)
	}
	rSize, cSize := grid.RowColLength()
	antennaLocations := make(map[rune][]util.Point)
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			cell := grid.At(r, c)
			if cell != '.' {
				antennaLocations[cell] = append(antennaLocations[cell], util.NewPoint(r, c))
			}
		}
	}
	antinodes := make(map[util.Point]bool)
	for _, locations := range antennaLocations {
		for i := 0; i < len(locations)-1; i++ {
			for j := i + 1; j < len(locations); j++ {
				points := []util.Point{}
				if isGold {
					points = append(points, findPointsGold(grid, locations[i], locations[j])...)
				} else {
					points = append(points, findPointsSilver(grid, locations[i], locations[j])...)
				}
				for _, point := range points {
					antinodes[point] = true
				}
			}
		}
	}
	//grid.PrintDebugWithDots(antinodes)
	println(len(antinodes))
}

func findPointsSilver(grid *util.Grid, startPoint, endPoint util.Point) []util.Point {
	points := []util.Point{}
	deltaPoint := endPoint.MinusPoint(startPoint)
	antinodePrev := startPoint.MinusPoint(deltaPoint)
	antinodeNext := endPoint.AddPoint(deltaPoint)
	if grid.IsValidPoint(antinodePrev) {
		points = append(points, antinodePrev)
	}
	if grid.IsValidPoint(antinodeNext) {
		points = append(points, antinodeNext)
	}
	return points
}

func findPointsGold(grid *util.Grid, startPoint, endPoint util.Point) []util.Point {
	points := []util.Point{}
	deltaPoint := endPoint.MinusPoint(startPoint)
	antinodePrev := startPoint
	antinodeNext := endPoint
	for grid.IsValidPoint(antinodeNext) {
		points = append(points, antinodeNext)
		antinodeNext = antinodeNext.AddPoint(deltaPoint)
	}
	for grid.IsValidPoint(antinodePrev) {
		points = append(points, antinodePrev)
		antinodePrev = antinodePrev.MinusPoint(deltaPoint)
	}
	return points
}
