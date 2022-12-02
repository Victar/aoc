package main

import (
	"adventofcode/util"
	"sort"
	"strconv"
)

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	ar, err := util.ReadFile("year2022/day1/input.txt")
	if err != nil {
		panic(err)
	}
	var max = 0
	var current = 0
	for _, s := range ar {
		if len(s) == 0 {
			if current > max {
				max = current
			}
			current = 0
		} else {
			cur, _ := strconv.Atoi(s)
			current += cur
		}
	}
	println(max)
}

func runGold() {
	ar, err := util.ReadFile("year2022/day1/input.txt")
	if err != nil {
		println(err)
	}
	var all []int
	var current = 0
	for _, s := range ar {
		if len(s) == 0 {
			all = append(all, current)
			current = 0
		} else {
			cur, _ := strconv.Atoi(s)
			current += cur
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(all)))
	println(all[0] + all[1] + all[2])
}
