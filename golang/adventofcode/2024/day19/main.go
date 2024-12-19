package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

var DAY = "19"

var DP = make(map[string]int)

func main() {
	runBoth()
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}

	patterns := strings.Split(lines[0], ", ")
	silverAns := 0
	goldAns := 0

	for i := 2; i < len(lines); i++ {
		curAns := countWays(lines[i], patterns)
		if curAns > 0 {
			silverAns++
			goldAns += curAns
		}
	}
	fmt.Println(silverAns)
	fmt.Println(goldAns)
}

func countWays(design string, patterns []string) int {
	if design == "" {
		return 1
	}

	if v, found := DP[design]; found {
		return v
	}

	total := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design, pattern) {
			remaining := design[len(pattern):]
			total += countWays(remaining, patterns)
		}
	}
	DP[design] = total
	return total
}
