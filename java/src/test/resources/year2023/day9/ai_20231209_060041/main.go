
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := parseNumbers(scanner.Text())
		nextValue := extrapolateNextValue(numbers)
		sum += nextValue
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func parseNumbers(line string) []int {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		numbers[i], _ = strconv.Atoi(field)
	}
	return numbers
}

func extrapolateNextValue(numbers []int) int {
	for {
		differences := calculateDifferences(numbers)
		if isZeroSlice(differences) {
			break
		}
		numbers = differences
	}
	return numbers[len(numbers)-1] + differencesAtEnd(numbers)
}

func calculateDifferences(numbers []int) []int {
	differences := make([]int, len(numbers)-1)
	for i := 0; i < len(numbers)-1; i++ {
		differences[i] = numbers[i+1] - numbers[i]
	}
	return differences
}

func isZeroSlice(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
}

func differencesAtEnd(numbers []int) int {
	last := numbers[len(numbers)-1]
	if len(numbers) > 1 {
		secondLast := numbers[len(numbers)-2]
		return last - secondLast
	}
	return last
}
