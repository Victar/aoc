
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Fields(string(data))
	packages := make([]int, 0, len(lines))

	for _, line := range lines {
		weight, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		packages = append(packages, weight)
	}

	return packages, nil
}

func sum(slice []int) int {
	total := 0
	for _, v := range slice {
		total += v
	}
	return total
}

func qe(slice []int) int64 {
	product := int64(1)
	for _, v := range slice {
		product *= int64(v)
	}
	return product
}

func findCombination(packages []int, groupWeight int, current []int, startIndex int, minSize int) ([]int, int64) {
	if sum(current) == groupWeight && (minSize == 0 || len(current) <= minSize) {
		return current, qe(current)
	}
	if sum(current) > groupWeight || len(current) > minSize && minSize != 0 {
		return nil, math.MaxInt64
	}

	var bestCombination []int
	minQE := int64(math.MaxInt64)

	for i := startIndex; i < len(packages); i++ {
		next := make([]int, len(current)+1)
		copy(next, current)
		next[len(next)-1] = packages[i]
		comb, combQE := findCombination(packages, groupWeight, next, i+1, minSize)
		if comb != nil && (bestCombination == nil || combQE < minQE || (combQE == minQE && len(comb) < len(bestCombination))) {
			bestCombination = comb
			minQE = combQE
			minSize = len(comb)
		}
	}

	return bestCombination, minQE
}

func main() {
	packages, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	sort.Ints(packages)
	groupWeight := sum(packages) / 3

	_, qe := findCombination(packages, groupWeight, nil, 0, 0)
	fmt.Println(qe)
}
