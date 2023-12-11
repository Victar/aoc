
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

	sum := 0
	for scanner.Scan() {
		numbers := parseNumbers(scanner.Text())
		nextValue := findNextValue(numbers)
		sum += nextValue
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func parseNumbers(line string) []int {
	parts := strings.Split(line, " ")
	numbers := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		numbers[i] = num
	}
	return numbers
}

func findNextValue(numbers []int) int {
	for {
		if isAllZeroes(numbers) {
			break
		}
		numbers = differences(numbers)
	}
	return extrapolate(numbers, len(numbers))
}

func isAllZeroes(numbers []int) bool {
	for _, v := range numbers {
		if v != 0 {
			return false
		}
	}
	return true
}

func differences(numbers []int) []int {
	diffs := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		diffs[i] = numbers[i+1] - numbers[i]
	}
	return diffs
}

func extrapolate(numbers []int, length int) int {
	extrapolated := make([]int, length+1)
	for i, number := range numbers {
		extrapolated[i] = number
	}
	for i := len(numbers) - 1; i >= 0; i-- {
		extrapolated[i+1] += extrapolated[i]
	}
	return extrapolated[len(extrapolated)-1]
}
