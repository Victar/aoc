
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file.
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := convertLineToInts(line)
		nextValue := getNextValue(numbers)
		sum += nextValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from input file: %v\n", err)
		return
	}

	// Print the sum of the extrapolated values.
	fmt.Println(sum)
}

// convertLineToInts converts a line of space-separated numbers into a slice of ints.
func convertLineToInts(line string) []int {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		number, err := strconv.Atoi(field)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			continue
		}
		numbers[i] = number
	}
	return numbers
}

// getNextValue calculates the next value in the sequence by generating sequences of differences.
func getNextValue(values []int) int {
	for {
		differences := make([]int, 0, len(values)-1)
		isZeroSequence := true
		for i := 0; i < len(values)-1; i++ {
			diff := values[i+1] - values[i]
			differences = append(differences, diff)
			if diff != 0 {
				isZeroSequence = false
			}
		}
		if isZeroSequence {
			break
		}
		values = differences
	}

	// The last element of values now contains the constant difference
	// Add it to the last original number to get the next value.
	lastValue := values[len(values)-1] + values[len(values)-2]
	return lastValue
}
