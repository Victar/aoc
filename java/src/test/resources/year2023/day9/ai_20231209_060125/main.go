
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// readFile reads the input data from the given file.
func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// parseLine converts a space-separated string of numbers into a slice of integers.
func parseLine(line string) ([]int, error) {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		numbers[i] = number
	}
	return numbers, nil
}

// calculateDifferences returns a slice of differences between adjacent elements of the input slice.
func calculateDifferences(numbers []int) []int {
	differences := make([]int, len(numbers)-1)
	for i := range differences {
		differences[i] = numbers[i+1] - numbers[i]
	}
	return differences
}

// allZeros checks if all elements in the slice are zeros.
func allZeros(numbers []int) bool {
	for _, number := range numbers {
		if number != 0 {
			return false
		}
	}
	return true
}

// extrapolateNextValue calculates the next value of the sequence based on the differences.
func extrapolateNextValue(numbers []int) int {
	// Keep calculating differences until we get all zeroes.
	for {
		differences := calculateDifferences(numbers)
		if allZeros(differences) {
			break
		}
		numbers = differences
	}
	// The next value in the original history is the last value plus the last difference (which is 0).
	return numbers[len(numbers)-1]
}

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, line := range lines {
		numbers, err := parseLine(line)
		if err != nil {
			panic(err)
		}
		nextValue := extrapolateNextValue(numbers)
		sum += nextValue
	}

	// Print the sum of extrapolated values.
	fmt.Println(sum)
}
