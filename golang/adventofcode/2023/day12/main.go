package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "12"

func main() {
	runAny(1)
	runAny(5)
}

var DP = map[string]int{}

type record struct {
	line    string
	numbers []int
}

func runAny(count int) {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	records := []record{}
	for _, line := range lines {
		parts := strings.Split(line, " ")

		line := strings.Repeat(parts[0]+"?", count)
		numbersS := strings.Repeat(parts[1]+",", count)
		line = line[:len(line)-1]
		numbersS = numbersS[:len(numbersS)-1]

		numbersStr := strings.Split(numbersS, ",")
		numbers := make([]int, len(numbersStr))

		for i, ns := range numbersStr {
			numbers[i], _ = strconv.Atoi(ns)
		}
		records = append(records, record{line, numbers})
	}
	ans := 0
	for _, record := range records {
		options := countOptions(record, 0, 0, 0)
		ans += options
		DP = make(map[string]int)
	}
	fmt.Println(ans)
}

func countOptions(record record, i int, bi int, current int) int {
	dpKey := strconv.Itoa(i) + "_" + strconv.Itoa(bi) + "_" + strconv.Itoa(current)
	value, ok := DP[dpKey]
	if ok {
		return value
	}
	if i == len(record.line) {
		if bi == len(record.numbers) && current == 0 {
			return 1
		} else if bi == len(record.numbers)-1 && current == record.numbers[bi] {
			return 1
		}
		return 0
	}
	ans := 0
	if record.line[i] == '?' || record.line[i] == '.' {
		if current == 0 {
			ans += countOptions(record, i+1, bi, 0)
		} else if current > 0 && bi < len(record.numbers) && record.numbers[bi] == current {
			ans += countOptions(record, i+1, bi+1, 0)
		}
	}
	if record.line[i] == '?' || record.line[i] == '#' {
		if bi < len(record.numbers) && current < record.numbers[bi] {
			ans += countOptions(record, i+1, bi, current+1)
		}
	}
	DP[dpKey] = ans
	return ans
}
