package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "14"

const (
	maxR    = 103 // height
	maxC    = 101 // width
	seconds = 100
)

type Robot struct {
	position util.Point
	velocity util.Point
}

func main() {
	runBoth()
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var robots []Robot
	for _, line := range lines {
		robots = append(robots, parseRobot(line))
	}
	foundSilver := false
	foundGold := false
	answerSilver := 0
	answerGold := 0
	second := 0
	for !foundSilver || !foundGold {
		second++
		moveRobots(robots)
		if second == seconds {
			foundSilver = true
			answerSilver = countQuadrant(robots)
		}
		if isTreeBasedOnSegment(robots) {
			printTree(robots)
			foundGold = true
			answerGold = second
		}
	}
	fmt.Println(answerSilver)
	fmt.Println(answerGold)
}

func isTreeBasedOnSegment(robots []Robot) bool {
	return strings.Contains(robotsToString(robots), "xxxxxxxxxx")
}

func printTree(robots []Robot) {
	print(robotsToString(robots))
}

func robotsToString(robots []Robot) string {
	visited := map[util.Point]bool{}
	for _, robot := range robots {
		visited[robot.position] = true
	}
	var sb strings.Builder
	for r := 0; r < maxR; r++ {
		for c := 0; c < maxC; c++ {
			if visited[util.Point{r, c}] {
				sb.WriteByte('x')
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func countQuadrant(robots []Robot) int {
	midR, midC := maxR/2, maxC/2
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.position.R != midR && robot.position.C != midC {
			switch {
			case robot.position.R < midR && robot.position.C < midC:
				q1++
			case robot.position.R >= midR && robot.position.C < midC:
				q2++
			case robot.position.R < midR && robot.position.C >= midC:
				q3++
			case robot.position.R >= midR && robot.position.C >= midC:
				q4++
			}
		}
	}
	return q1 * q2 * q3 * q4
}
func moveRobots(robots []Robot) {
	for i := range robots {
		robots[i].position = robots[i].position.AddPointInBorder(robots[i].velocity, maxR, maxC)
	}
}

func parseRobot(line string) Robot {
	parts := strings.Split(line, " ")
	posPart := parts[0]
	velPart := parts[1]
	posValues := strings.Split(strings.TrimPrefix(posPart, "p="), ",")
	velValues := strings.Split(strings.TrimPrefix(velPart, "v="), ",")
	c, _ := strconv.Atoi(posValues[0])
	r, _ := strconv.Atoi(posValues[1])
	vc, _ := strconv.Atoi(velValues[0])
	vr, _ := strconv.Atoi(velValues[1])
	return Robot{util.Point{r, c}, util.Point{vr, vc}}
}
