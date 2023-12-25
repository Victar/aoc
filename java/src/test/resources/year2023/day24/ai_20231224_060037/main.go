
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Hailstone struct {
	Px, Py           int
	Vx, Vy           int
	XInRange, YInRange bool
}

func main() {
	hailstones, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}
	count := countIntersections(hailstones)
	fmt.Println(count)
}

func parseInput(filename string) ([]Hailstone, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var hailstones []Hailstone
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " @ ")
		position := strings.Split(parts[0], ", ")
		velocity := strings.Split(parts[1], ", ")

		px, _ := strconv.Atoi(position[0])
		py, _ := strconv.Atoi(position[1])
		vx, _ := strconv.Atoi(velocity[0])
		vy, _ := strconv.Atoi(velocity[1])

		hailstone := Hailstone{
			Px: px, Py: py,
			Vx: vx, Vy: vy,
			XInRange: isInTestRange(px),
			YInRange: isInTestRange(py),
		}
		hailstones = append(hailstones, hailstone)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return hailstones, nil
}

func isInTestRange(position int) bool {
	return position >= 200000000000000 && position <= 400000000000000
}

func countIntersections(hailstones []Hailstone) int {
	intersectionCount := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if willIntersect(hailstones[i], hailstones[j]) {
				intersectionCount++
			}
		}
	}
	return intersectionCount
}

func willIntersect(a, b Hailstone) bool {
	if a.Vx == b.Vx && a.Px != b.Px {
		return false
	} else if a.Vy == b.Vy && a.Py != b.Py {
		return false
	}
	
	dx := b.Px - a.Px
	dy := b.Py - a.Py
	vx_diff := b.Vx - a.Vx
	vy_diff := b.Vy - a.Vy

	if vx_diff == 0 && isInTestRange(a.Px) {
		if dy%vy_diff == 0 && isInTestRange(a.Py+dy) {
			return true
		}
	} else if vy_diff == 0 && isInTestRange(a.Py) {
		if dx%vx_diff == 0 && isInTestRange(a.Px+dx) {
			return true
		}
	} else if vx_diff != 0 && vy_diff != 0 {
		if dx%vx_diff == 0 && dy%vy_diff == 0 && dx/vx_diff == dy/vy_diff {
			px := a.Px + dx/vx_diff*a.Vx
			py := a.Py + dy/vy_diff*a.Vy
			if isInTestRange(px) && isInTestRange(py) {
				return true
			}
		}
	}
	return false
}
