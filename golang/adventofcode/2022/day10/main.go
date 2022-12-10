package main

import (
	"adventofcode/util"
	"math"
	"strconv"
	"strings"
)

func main() {
	runSilverAndGold()
}

var x = 1
var s = 0
var cycle = 0
var size = 40
var image [6]string

func runSilverAndGold() {
	lines, err := util.ReadFile("year2022/day10/input.txt")
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(image); i++ {
		image[i] = strings.Repeat(" ", 40)
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "noop") {
			doCycle()
		}
		if strings.HasPrefix(line, "addx") {
			doCycle()
			doCycle()
			add, _ := strconv.Atoi(line[5:])
			x += add
		}
	}
	println(s)
	for i := 0; i < len(image); i++ {
		println(image[i])
	}
}

func doCycle() {
	raw := cycle / size
	pos := cycle % size
	if math.Abs(float64(x-pos)) < 2 {
		cur := image[raw]
		image[raw] = cur[0:pos] + "#" + cur[pos+1:]
	}
	cycle++
	if (cycle+20)%size == 0 {
		s += cycle * x
	}
}
