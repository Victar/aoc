
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y, z int
}

type Brick struct {
	from Coord
	to   Coord
}

func readBricks(filename string) ([]Brick, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bricks []Brick
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, "~")
		if len(coords) != 2 {
			continue
		}

		start := strings.Split(coords[0], ",")
		end := strings.Split(coords[1], ",")
		if len(start) != 3 || len(end) != 3 {
			continue
		}

		x1, y1, z1 := parseCoord(start)
		x2, y2, z2 := parseCoord(end)
		bricks = append(bricks, Brick{
			from: Coord{x: x1, y: y1, z: z1},
			to:   Coord{x: x2, y: y2, z: z2},
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bricks, nil
}

func parseCoord(s []string) (int, int, int) {
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	z, _ := strconv.Atoi(s[2])
	return x, y, z
}

func main() {
	bricks, err := readBricks("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v\n", err)
	}

	couldFall := make(map[int][]int) // Maps brick index to bricks directly above it
	for i, brick := range bricks {
		for j, other := range bricks {
			if i != j && isSupporting(brick, other) {
				couldFall[j] = append(couldFall[j], i)
			}
		}
	}

	safeCount := 0
	for i, supporters := range couldFall {
		if canDisintegrate(supporters, couldFall) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func isSupporting(below, above Brick) bool {
	// Check if 'below' brick can support 'above' brick after everything settles
	for x := max(below.from.x, above.from.x); x <= min(below.to.x, above.to.x); x++ {
		for y := max(below.from.y, above.from.y); y <= min(below.to.y, above.to.y); y++ {
			if above.from.z-1 == below.from.z || above.from.z-1 == below.to.z {
				return true
			}
		}
	}
	return false
}

func canDisintegrate(supporters []int, couldFall map[int][]int) bool {
	for _, supporter := range supporters {
		if len(couldFall[supporter]) == 0 {
			continue
		}

		allHaveOtherSupport := true
		for _, other := range couldFall[supporter] {
			if len(couldFall[other]) == 1 { // if only supported by current brick
				allHaveOtherSupport = false
				break
			}
		}
		if !allHaveOtherSupport {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
