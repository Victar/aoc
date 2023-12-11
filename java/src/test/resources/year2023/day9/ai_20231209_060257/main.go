
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var histories [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valuesStr := strings.Fields(scanner.Text())
		var history []int
		for _, valueStr := range valuesStr {
			value, _ := strconv.Atoi(valueStr)
			history = append(history, value)
		}
		histories = append(histories, history)
	}

	sum := 0
	for _, history := range histories {
		sum += extrapolateNextValue(history)
	}

	fmt.Println(sum)
}

func extrapolateNextValue(history []int) int {
	for {
		allZeroes := true
		differences := make([]int, len(history)-1)
		for i := 0; i < len(history)-1; i++ {
			differences[i] = history[i+1] - history[i]
			if differences[i] != 0 {
				allZeroes = false
			}
		}
		if allZeroes {
			break
		}
		history = differences
	}

	// Calculate the next extrapolated value by taking the last value and adding the latest difference.
	nextValue := history[len(history)-1] + history[len(history)-2]
	return nextValue
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
