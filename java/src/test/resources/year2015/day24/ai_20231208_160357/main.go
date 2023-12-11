
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(path string) ([]int, error) {
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(fileContent), "\n")

	var weights []int
	for _, line := range lines {
		if line == "" {
			continue
		}
		weight, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		weights = append(weights, weight)
	}
	return weights, nil
}

func sum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func findCombination(weights []int, groupWeight int, current []int, best *[]int, bestQE *int64) {
	currentSum := sum(current)
	currentQE := int64(1)

	if currentSum > groupWeight {
		return
	}
	if currentSum == groupWeight {
		for _, w := range current {
			currentQE *= int64(w)
		}
		if len(current) < len(*best) || (len(current) == len(*best) && currentQE < *bestQE) {
			*best = append((*best)[:0:0], current...)
			*bestQE = currentQE
		}
		return
	}
	if len(*best) > 0 && len(current) >= len(*best) {
		// This branch can't possibly be better than our current best.
		return
	}

	for i, w := range weights {
		newCombination := append([]int(nil), current...)
		newCombination = append(newCombination, w)
		findCombination(weights[i+1:], groupWeight, newCombination, best, bestQE)
	}
}

func quantumEntanglement(weights []int) int64 {
	totalWeight := sum(weights)
	if totalWeight%3 != 0 {
		fmt.Println("Invalid input: Total weight is not divisible by 3.")
		return -1
	}
	groupWeight := totalWeight / 3

	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	best := make([]int, math.MaxInt32)
	var bestQE int64 = math.MaxInt64

	// Find a valid combination for the first group.
	findCombination(weights, groupWeight, []int{}, &best, &bestQE)

	return bestQE
}

func main() {
	weights, err := readInput("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	result := quantumEntanglement(weights)
	fmt.Println(result)
}
