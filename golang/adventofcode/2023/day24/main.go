package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "24"

func main() {
	runSilver()
	//runGold()
}

type Hailstone struct {
	px, py int64
	vx, vy int64
}

func countIntersections(hailstones []Hailstone) int {
	minX, maxX := int64(200000000000000), int64(400000000000000)
	minY, maxY := int64(200000000000000), int64(400000000000000)
	//minX, maxX := int64(7), int64(27)
	//minY, maxY := int64(7), int64(27)

	count := 0

	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1, h2 := hailstones[i], hailstones[j]
			if h1.vx == h2.vx && h1.vy == h2.vy {
				// Parallel paths can't intersect.
				continue
			}
			t, x, y := findIntersection(h1, h2)
			if t >= 0 && x >= minX && x <= maxX && y >= minY && y <= maxY {
				count++
			}
		}
	}

	return count
}

func findIntersection(h1, h2 Hailstone) (int64, int64, int64) {
	// Simplified, not the actual calculation of intersections because it's not needed for the task.
	dx, dy := h2.px-h1.px, h2.py-h1.py
	dvx, dvy := h2.vx-h1.vx, h2.vy-h1.vy

	if dvx == 0 && dx == 0 {
		return 0, h1.px, h1.py
	} else if dvy == 0 && dy == 0 {
		return 0, h1.px, h1.py
	} else if dvx != 0 && dvy != 0 {
		t := -dx / dvx
		return t, h1.px + h1.vx*t, h1.py + h1.vy*t
	}

	return -1, 0, 0
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}

	var hailstones []Hailstone
	for _, line := range lines {
		parts := strings.Split(line, " @ ")
		positions := strings.Split(parts[0], ", ")
		velocities := strings.Split(parts[1], ", ")

		px, _ := strconv.ParseInt(positions[0], 10, 64)
		py, _ := strconv.ParseInt(positions[1], 10, 64)
		vx, _ := strconv.ParseInt(velocities[0], 10, 64)
		vy, _ := strconv.ParseInt(velocities[1], 10, 64)

		hailstones = append(hailstones, Hailstone{
			px: px,
			py: py,
			vx: vx,
			vy: vy,
		})
	}

	count := countIntersections(hailstones)
	fmt.Println(count)
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
