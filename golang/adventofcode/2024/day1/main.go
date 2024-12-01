package main

import (
	"adventofcode/util"
	"sort"
	"strconv"
	"strings"
)

var DAY = "1"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var leftList, rightList []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		leftNum, _ := strconv.Atoi(numbers[0])
		rightNum, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := leftList[i] - rightList[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	println(totalDistance)
}

func runGold() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var leftList, rightList []int

	for _, line := range lines {
		numbers := strings.Fields(line)
		leftNum, _ := strconv.Atoi(numbers[0])
		rightNum, _ := strconv.Atoi(numbers[1])
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	countMap := make(map[int]int)
	for _, right := range rightList {
		countMap[right]++
	}

	totalSimilarity := 0
	for _, left := range leftList {
		if count, found := countMap[left]; found {
			totalSimilarity += left * count
		}
	}
	println(totalSimilarity)
}
