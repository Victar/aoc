package main

import (
	"adventofcode/util"
	"regexp"
	"strconv"
)

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2022/day4/input.txt")
	if err != nil {
		panic(err)
	}
	var score = 0
	r := regexp.MustCompile("[^,-]+")
	for _, line := range lines {
		strarr := r.FindAllString(line, -1)
		var s1b, _ = strconv.Atoi(strarr[0])
		var s1e, _ = strconv.Atoi(strarr[1])
		var s2b, _ = strconv.Atoi(strarr[2])
		var s2e, _ = strconv.Atoi(strarr[3])

		if s1b <= s2b && s1e >= s2e || s2b <= s1b && s2e >= s1e {
			score++
		}
	}
	println(score)

}

func runGold() {
	lines, err := util.ReadFile("year2022/day4/input.txt")
	if err != nil {
		panic(err)
	}
	var score = 0
	r := regexp.MustCompile("[^,-]+")
	for _, line := range lines {
		strarr := r.FindAllString(line, -1)
		var s1b, _ = strconv.Atoi(strarr[0])
		var s1e, _ = strconv.Atoi(strarr[1])
		var s2b, _ = strconv.Atoi(strarr[2])
		var s2e, _ = strconv.Atoi(strarr[3])

		if s1b <= s2b && s2b <= s1e || s2b <= s1b && s1b <= s2e {
			score++
		}
	}
	println(score)
}
