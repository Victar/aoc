
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Brick struct {
	start       Point
	end         Point
	supports    map[string]struct{}
	supportedBy map[string]struct{}
}

type Point struct {
	x, y, z int
}

func main() {
	bricks := readBricks("input.txt")
	settledBricks := settleBricks(bricks)
	safeBricks := countSafeToDisintegrate(settledBricks)
	fmt.Println(safeBricks)
}

func readBricks(filename string) map[string]*Brick {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bricks := make(map[string]*Brick)
	scanner := bufio.NewScanner(file)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		parts := strings.Split(line, "~")
		start := parsePoint(parts[0])
		end := parsePoint(parts[1])

		id := "Brick" + strconv.Itoa(lineNumber)
		bricks[id] = &Brick{
			start:       start,
			end:         end,
			supports:    make(map[string]struct{}),
			supportedBy: make(map[string]struct{}),
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return bricks
}

func parsePoint(s string) Point {
	coords := strings.Split(s, ",")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	z, _ := strconv.Atoi(coords[2])
	return Point{x: x, y: y, z: z}
}

func settleBricks(bricks map[string]*Brick) map[string]*Brick {
	for _, brick := range bricks {
		for otherID, otherBrick := range bricks {
			if brick == otherBrick {
				continue
			}
			if isSupported(brick, otherBrick) {
				brick.supportedBy[otherID] = struct{}{}
				otherBrick.supports[otherID] = struct{}{}
			}
		}
	}
	return bricks
}

func isSupported(brick, otherBrick *Brick) bool {
	// Check if 'otherBrick' is directly under 'brick'
	ex, ey, ez := expandBrick(brick)
	for x := ex.start; x <= ex.end; x++ {
		for y := ey.start; y <= ey.end; y++ {
			for z := brick.start.z - 1; z >= otherBrick.end.z; z-- {
				if _, ok := otherBrick.supports[fmt.Sprintf("%d,%d,%d", x, y, z)]; ok {
					return true // Supported by otherBrick at this position
				}
			}
		}
	}
	return false
}

func expandBrick(brick *Brick) (xRange, yRange, zRange Point) {
	return Point{start: min(brick.start.x, brick.end.x), end: max(brick.start.x, brick.end.x)},
		Point{start: min(brick.start.y, brick.end.y), end: max(brick.start.y, brick.end.y)},
		Point{start: min(brick.start.z, brick.end.z), end: max(brick.start.z, brick.end.z)}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countSafeToDisintegrate(bricks map[string]*Brick) (safeCount int) {
	for id, brick := range bricks {
		if isSafeToDisintegrate(brick, bricks) {
			safeCount++
			fmt.Println("Safe to disintegrate:", id)
		}
	}
	return safeCount
}

func isSafeToDisintegrate(brick *Brick, bricks map[string]*Brick) bool {
	for supportedID := range brick.supports {
		requiresBrick := true
		for otherID := range bricks[supportedID].supportedBy {
			if otherID == brick && len(bricks[supportedID].supportedBy) == 1 {
				requiresBrick = true
				break
			} else if otherID != brick && len(bricks[supportedID].supportedBy) > 1 {
				requiresBrick = false
				break
			}
		}
		if requiresBrick {
			return false // There's at least one brick that requires this one for support
		}
	}
	return true
}

func main() {
	fmt.Println("Parsing bricks and checking for safe disintegration...")
	numSafeToDisintegrate := countSafeToDisintegrate()
	fmt.Println(numSafeToDisintegrate)
}
