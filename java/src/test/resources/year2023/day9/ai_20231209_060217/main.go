
// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file containing the input data.
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line.
	scanner := bufio.NewScanner(file)
	sumExtrapolated := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		nums := convertToSliceInt64(strings.Fields(line))
		nextValue := predictNextValue(nums)
		sumExtrapolated += nextValue
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print the answer.
	fmt.Println(sumExtrapolated)
}

// predictNextValue takes a sequence of int64 and predicts the next value.
func predictNextValue(nums []int64) int64 {
	for {
		differences := calculateDifferences(nums)
		if allZeros(differences) {
			break
		}
		nums = differences
	}
	extrapolatedValue := nums[len(nums)-1] + nums[len(nums)-2]
	return extrapolatedValue
}

// calculateDifferences calculates the differences between the consecutive numbers in a slice.
func calculateDifferences(nums []int64) []int64 {
	var differences []int64
	for i := 1; i < len(nums); i++ {
		differences = append(differences, nums[i]-nums[i-1])
	}
	return differences
}

// allZeros checks if all elements in a slice of int64 are zeros.
func allZeros(nums []int64) bool {
	for _, v := range nums {
		if v != 0 {
			return false
		}
	}
	return true
}

// convertToSliceInt64 takes a slice of strings and converts it to a slice of int64.
func convertToSliceInt64(strs []string) []int64 {
	var result []int64
	for _, s := range strs {
		number, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, number)
	}
	return result
}
