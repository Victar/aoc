package main

import (
	"adventofcode/util"
	"fmt"
	"regexp"
	"strconv"
)

var DAY = "3"

func main() {
	runAny(false)
	runAny(true)
}

func runAny(isGold bool) {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	oneLine := ""
	for _, line := range lines {
		oneLine += line
	}
	if isGold {
		dontDoRe := regexp.MustCompile(`don't\(\).*?do\(\)`)
		oneLine = dontDoRe.ReplaceAllString(oneLine, "")
	}
	sum := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(oneLine, -1)
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}
	fmt.Println(sum)
}
