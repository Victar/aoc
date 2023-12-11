
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Function to obtain all combinations of `items` of length `length`.
func combinations(items []int, length int) <-chan []int {
	c := make(chan []int)
	go func() {
		defer close(c)
		generateCombinations(items, length, nil, c)
	}()
	return c
}

// Recursive helper function to generate combinations.
func generateCombinations(items []int, length int, prefix []int, c chan<- []int) {
	if length == 0 {
		c <- append([]int(nil), prefix...)
		return
	}
	for i, item := range items {
		generateCombinations(items[i+1:], length-1, append(prefix, item), c)
	}
}

// Function to check whether it's possible to evenly split remaining items into two groups of `target` weight.
func canSplitEvenly(items []int, count, target int) bool {
	for combination := range combinations(items, count) {
		if sum(combination) == target {
			remaining := remove(items, combination)
			if findSumCombination(remaining, target) {
				return true
			}
		}
	}
	return false
}

func findSumCombination(items []int, target int) bool {
	for setSize := 1; setSize <= len(items); setSize++ {
		for combination := range combinations(items, setSize) {
			if sum(combination) == target {
				return true
			}
		}
	}
	return false
}

// Function to calculate the sum of elements in a slice.
func sum(items []int) int {
	total := 0
	for _, x := range items {
		total += x
	}
	return total
}

// Function to remove elements from original slice.
func remove(original, toRemove []int) []int {
	remaining := make([]int, len(original))
	copy(remaining, original)
	for _, item := range toRemove {
		for i, origItem := range remaining {
			if item == origItem {
				remaining = append(remaining[:i], remaining[i+1:]...)
				break
			}
		}
	}
	return remaining
}

// Function to calculate the quantum entanglement of a group of packages.
func quantumEntanglement(items []int) int64 {
	result := int64(1)
	for _, item := range items {
		result *= int64(item)
	}
	return result
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input.txt: %s\n", err)
		os.Exit(1)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	var packages []int
	for _, line := range lines {
		weight, _ := strconv.Atoi(line)
		packages = append(packages, weight)
	}

	totalWeight := sum(packages)
	groupWeight := totalWeight / 3
	idealQE := int64(^uint64(0) >> 1) // Max positive int64 value as a starting best quantum entanglement.
	
	// Find the smallest possible size for the first group.
	for size := 1; size < len(packages); size++ {
		for combo := range combinations(packages, size) {
			if sum(combo) == groupWeight {
				if canSplitEvenly(remove(packages, combo), size, groupWeight) {
					qe := quantumEntanglement(combo)
					if qe < idealQE {
						idealQE = qe
					}					
				}
			}
		}
		// Stop if we found at least one suitable combination.
		if idealQE < int64(^uint64(0) >> 1) {
			break
		}
	}

	fmt.Println(idealQE)
}
