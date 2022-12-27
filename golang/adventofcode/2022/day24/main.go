package main

import (
	"adventofcode/util"
	"fmt"
	"image"
	"strconv"
)

var DAY = "24"

var mapField [][]byte
var mapBlizzardsList = make(map[int][]*Blizzard)
var xSize int
var ySize int

func main() {
	runBoth()
}

type Blizzard struct {
	point image.Point
	dir   byte
}

type PointField struct {
	point image.Point
	prev  *PointField
	step  int
}

func (pf PointField) getState() string {
	return strconv.Itoa(pf.point.X) + ":" + strconv.Itoa(pf.point.Y) + ":" + strconv.Itoa(pf.step)
}
func runBoth() {
	lines, err := util.ReadFile("year2022/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	ySize = len(lines)
	xSize = len(lines[0])

	var blizzardsList []*Blizzard
	for y := 0; y < ySize-1; y++ {
		for x := 0; x < xSize-1; x++ {
			var curr = lines[y][x]
			if curr == '>' || curr == '<' || curr == 'v' || curr == '^' {
				blizzardsList = append(blizzardsList, &Blizzard{
					point: image.Point{x, y},
					dir:   curr,
				})
			}
		}
	}
	mapBlizzardsList[0] = blizzardsList
	start := PointField{point: image.Point{1, 0}, step: 0}
	end := PointField{point: image.Point{xSize - 2, ySize - 1}, step: 0}
	steps := findPath(&start, &end)
	fmt.Printf("Silver: %d\n", steps)
	startTwo := PointField{point: image.Point{xSize - 2, ySize - 1}, step: steps}
	steps = findPath(&startTwo, &start)
	startThree := PointField{point: image.Point{1, 0}, step: steps}
	steps = findPath(&startThree, &end)
	fmt.Printf("Gold: %d", steps)

}

func findPath(start, end *PointField) (steps int) {
	queue := make([]*PointField, 0)
	visited := make(map[string]bool)
	queue = append(queue, start)
	visited[start.getState()] = true
	for len(queue) > 0 {
		size := len(queue)
		//println(size)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.point.X == end.point.X && curr.point.Y == end.point.Y {
				return curr.step
			}
			neighbors := []PointField{
				{point: curr.point.Add(image.Point{0, 0}), step: curr.step + 1, prev: curr},
				{point: curr.point.Add(image.Point{0, +1}), step: curr.step + 1, prev: curr},
				{point: curr.point.Add(image.Point{0, -1}), step: curr.step + 1, prev: curr},
				{point: curr.point.Add(image.Point{+1, 0}), step: curr.step + 1, prev: curr},
				{point: curr.point.Add(image.Point{-1, 0}), step: curr.step + 1, prev: curr},
			}
			for _, next := range neighbors {
				next := next
				if isValidAtStep(&next) && !visited[next.getState()] {
					queue = append(queue, &next)
					visited[next.getState()] = true
				}
			}
		}
	}
	return -1
}

func isValidAtStep(pointField *PointField) bool {

	if pointField.point.X < 0 || pointField.point.X >= xSize || pointField.point.Y < 0 || pointField.point.Y >= ySize {
		return false
	}

	blizzardsList := getBlizzardsList(pointField.step)
	for _, blizzard := range blizzardsList {
		if pointField.point.X == blizzard.point.X && pointField.point.Y == blizzard.point.Y {
			return false
		}
	}

	if (pointField.point.X == 1 && pointField.point.Y == 0) || (pointField.point.X == xSize-2 && pointField.point.Y == ySize-1) {
		return true
	}
	if pointField.point.X == 0 || pointField.point.X == xSize-1 || pointField.point.Y == 0 || pointField.point.Y == ySize-1 {
		return false
	}
	return true
}

func getBlizzardsList(step int) []*Blizzard {
	if blizzards, ok := mapBlizzardsList[step]; ok {
		return blizzards
	} else {
		blizzardsListPrevious := getBlizzardsList(step - 1)
		blizzardsListCopy := make([]*Blizzard, len(blizzardsListPrevious))
		for i, blizzard := range blizzardsListPrevious {
			blizzardsListCopy[i] = blizzard
		}
		for _, blizzard := range blizzardsListCopy {
			switch blizzard.dir {
			case '^':
				blizzard.point.Y--
			case 'v':
				blizzard.point.Y++
			case '<':
				blizzard.point.X--
			case '>':
				blizzard.point.X++
			}
			if blizzard.point.X < 1 {
				blizzard.point.X = xSize - 2
			}
			if blizzard.point.X >= xSize-1 {
				blizzard.point.X = 1
			}
			if blizzard.point.Y < 1 {
				blizzard.point.Y = ySize - 2
			}
			if blizzard.point.Y >= ySize-1 {
				blizzard.point.Y = 1
			}
		}
		mapBlizzardsList[step] = blizzardsListCopy
		return blizzardsListCopy
	}

}
