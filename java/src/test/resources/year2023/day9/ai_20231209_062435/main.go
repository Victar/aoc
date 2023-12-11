
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

	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		numbers, err := convertToNumbers(line)
		if err != nil {
			panic(err)
		}
		
		nextValue := findNextValue(numbers)
		sum += nextValue
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func convertToNumbers(fields []string) ([]int, error) {
	numbers := make([]int, len(fields))
	for i, field := range fields {
		number, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		numbers[i] = number
	}
	return numbers, nil
}

func findNextValue(numbers []int) int {
	for {
		diffs := differences(numbers)
		if allZeros(diffs) {
			break
		}
		numbers = diffs
	}

	lastNumber := numbers[len(numbers)-1]
	nextDifference := numbers[len(numbers)-2] - lastNumber

	return lastNumber + nextDifference
}

func differences(numbers []int) []int {
	diffs := make([]int, 0, len(numbers)-1)
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		diffs = append(diffs, diff)
	}
	return diffs
}

func allZeros(numbers []int) bool {
	for _, number := range numbers {
		if number != 0 {
			return false
		}
	}
	return true
}
