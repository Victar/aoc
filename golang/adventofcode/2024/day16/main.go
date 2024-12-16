package main

import (
	"adventofcode/util"
	"container/heap"
	"fmt"
)

var DAY = "16"

type State struct {
	pDir  PointDir
	score int
}

type PointDir struct {
	pos util.Point
	dir int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type PriorityQueue []*State

func main() {
	runBoth()
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/sample.txt")
	if err != nil {
		panic(err)
	}
	var start, end util.Point
	var grid = util.NewGridEmpty()
	for _, line := range lines {
		grid.AddRow(line)
	}
	rSize, cSize := grid.RowColSize()
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			cell := grid.At(r, c)
			if cell == 'S' {
				start = util.Point{r, c}
			}
			if cell == 'E' {
				end = util.Point{r, c}
			}
		}
	}

	pointDirCostFinish, pointDirCost := countPath(grid, start, end, 0)               // east
	pointDirCostFinishReverse, pointDirCostReverse := countPath(grid, end, start, 1) // finish to north depends on input
	bestPath := -1
	for _, curBestCost := range pointDirCostFinish {
		if bestPath < 0 || curBestCost < bestPath {
			bestPath = curBestCost
		}
	}

	bestPathPoints := make(map[util.Point]bool)
	for r := 0; r < rSize; r++ {
		for c := 0; c < cSize; c++ {
			for dir := 0; dir < 4; dir++ {
				for dirR := 0; dirR < 4; dirR++ {
					if grid.At(r, c) == '.' {
						pdStartMiddle := PointDir{util.Point{r, c}, dir}
						pdStartMiddleRevers := PointDir{util.Point{r, c}, dirR}
						startMidCost, startMidExists := pointDirCost[pdStartMiddle]
						midEndCost, midEndExists := pointDirCostReverse[pdStartMiddleRevers]
						if startMidExists && midEndExists {
							totalCost := startMidCost + midEndCost
							if totalCost == bestPath {
								//fmt.Println(totalCost, bestPath)
								//fmt.Println(totalCost-bestPath, startMidCost, midEndCost, totalCost, bestPath)
								bestPathPoints[util.Point{r, c}] = true
							}
						}

					}
				}
			}
		}
	}
	fmt.Println(len(pointDirCostReverse))
	fmt.Println("best reverse", pointDirCostFinishReverse)
	fmt.Println(len(pointDirCost))
	fmt.Println("best path", pointDirCostFinish)
	fmt.Println(bestPath)

	grid.PrintDebugWithDotsSymbol(bestPathPoints, '0')
	grid.PrintDebugVisitedOnly(bestPathPoints, '0')

	fmt.Println(len(bestPathPoints) + bestPath/1000 + 1) // Count turns
}

func countPath(grid *util.Grid, start util.Point, end util.Point, direction int) (map[PointDir]int, map[PointDir]int) {
	pointDirCost := make(map[PointDir]int)
	pointDirCostFinish := make(map[PointDir]int)
	var directions = []util.Point{util.Directions[util.RIGHT], util.Directions[util.DOWN], util.Directions[util.LEFT], util.Directions[util.UP]}
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{pDir: PointDir{start, direction}, score: 0})
	visited := make(map[PointDir]int)
	updateVisited := func(pointDir PointDir, score int) bool {
		if s, exists := visited[pointDir]; !exists || score < s {
			visited[pointDir] = score
			return true
		}
		return false
	}
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)
		curPointDir := current.pDir
		if s, exists := pointDirCost[curPointDir]; !exists || current.score < s {
			pointDirCost[curPointDir] = current.score
		}
		if current.pDir.pos == end {
			pointDirCostFinish[curPointDir] = current.score
		}

		for i := -1; i <= 1; i++ {
			newDir := (current.pDir.dir + i + 4) % 4
			newScore := current.score
			if i != 0 {
				newScore += 1000
			}
			newPos := current.pDir.pos.AddPoint(directions[newDir])
			if grid.AtPoint(newPos) != '#' {
				newScore++
				if updateVisited(PointDir{newPos, newDir}, newScore) {
					heap.Push(pq, &State{pDir: PointDir{newPos, newDir}, score: newScore})
				}
			}
		}
	}
	return pointDirCostFinish, pointDirCost
}
