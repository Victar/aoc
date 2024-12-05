package main

import (
	"adventofcode/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var DAY = "5"

func main() {
	runAny(false)
	runAny(true)

}

func runAny(isGold bool) {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}

	var rules []string
	var updates []string

	readingRules := true

	for _, line := range lines {
		if line == "" {
			readingRules = false
		}
		if readingRules {
			rules = append(rules, line)
		} else {
			updates = append(updates, line)
		}
	}
	result := 0
	for _, update := range updates {
		pagesStr := strings.Split(update, ",")
		fixed := false
		swapCount := 0
		for {
			fixed, pagesStr, swapCount = checkAndSwap(pagesStr, rules, swapCount)
			if fixed {
				break
			}
		}
		if (swapCount > 0 && isGold) || (swapCount == 0 && !isGold) {
			middlePageIndex := len(pagesStr) / 2
			middlePage, _ := strconv.Atoi(pagesStr[middlePageIndex])
			result += middlePage
		}
	}
	fmt.Println(result)
}

func checkAndSwap(pages []string, rules []string, swapsCount int) (bool, []string, int) {
	pageMap := make(map[string]bool)
	for _, pageNum := range pages {
		pageMap[pageNum] = true
	}
	for _, rule := range rules {
		ruleAr := strings.Split(rule, "|")
		left, right := ruleAr[0], ruleAr[1]
		if pageMap[left] && pageMap[right] {
			leftIndex := slices.Index(pages, left)
			rightIndex := slices.Index(pages, right)
			if leftIndex > rightIndex {
				pages[leftIndex], pages[rightIndex] = pages[rightIndex], pages[leftIndex]
				swapsCount++
				return false, pages, swapsCount
			}
		}
	}
	return true, pages, swapsCount
}
