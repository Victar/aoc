package main

import (
	"adventofcode/util"
	"strconv"
	"strings"
)

func main() {
	runSilver()
	runGold()
}

func runAny(reverse bool) {
	lines, err := util.ReadFile("year2022/day5/input.txt")
	if err != nil {
		panic(err)
	}
	var size = len(lines[0])/4 + 1
	var all = make([][]string, size)
	for i := 0; i < size; i++ {
		all[i] = make([]string, 0)
	}
	var needToAdd = true
	for _, line := range lines {
		if strings.HasPrefix(line, " 1 ") {
			needToAdd = false
		}
		if needToAdd {
			addToSlice(all, line+" ")
		}
		if strings.HasPrefix(line, "move") {
			doMove(all, line+" ", reverse)
		}
	}
	printAll(all, true)
}

func doMove(all [][]string, input string, reverse bool) {
	parts := strings.Fields(input)
	var count, _ = strconv.Atoi(parts[1])
	var from, _ = strconv.Atoi(parts[3])
	var to, _ = strconv.Atoi(parts[5])
	all[from-1] = all[from-1]
	take := all[from-1][len(all[from-1])-count:]
	all[from-1] = all[from-1][0 : len(all[from-1])-count]
	if reverse {
		var takeReverse []string
		for i := len(take) - 1; i >= 0; i-- {
			takeReverse = append(takeReverse, take[i])
		}
		all[to-1] = append(all[to-1], takeReverse...)
	} else {
		all[to-1] = append(all[to-1], take...)
	}
}

func printAll(all [][]string, resultOnly bool) {
	if !resultOnly {
		for i := 0; i < len(all); i++ {
			println(strings.Join(all[i], ", "))
		}
	}
	for i := 0; i < len(all); i++ {
		print(all[i][len(all[i])-1])
	}
	println()
}
func addToSlice(all [][]string, input string) {
	for i := 0; i < len(all); i++ {
		var cur = input[i*4 : (i+1)*4]
		if strings.Contains(cur, "[") {
			all[i] = append([]string{cur}, all[i]...)
		}
	}
}
func runSilver() {
	runAny(true)
}

func runGold() {
	runAny(false)
}
