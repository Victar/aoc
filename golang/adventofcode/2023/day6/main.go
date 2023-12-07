package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

var DAY = "6"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	times := strings.Fields(lines[0])
	records := strings.Fields(lines[1])
	var ways = 1
	for i, time := range times {
		timeNumber, err := strconv.Atoi(time)
		if err == nil {
			record, _ := strconv.Atoi(records[i])
			ways *= calculateWays(timeNumber, record)
		}
	}
	println(ways)

}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	times := strings.Join(strings.Fields(lines[0])[1:], "")
	records := strings.Join(strings.Fields(lines[1])[1:], "")
	time, err := strconv.Atoi(times)
	record, err := strconv.Atoi(records)
	println(calculateWays(time, record))
}

func calculateWays(time, record int) int {
	count := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		distance := holdTime * (time - holdTime)
		if distance > record {
			count++
		}
	}
	return count
}
