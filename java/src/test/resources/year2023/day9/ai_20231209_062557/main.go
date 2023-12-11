
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var nums []int
		for _, field := range line {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		data = append(data, nums)
	}
	return data, scanner.Err()
}

func differences(seq []int) []int {
	diff := make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		diff[i] = seq[i+1] - seq[i]
	}
	return diff
}

func allZeroes(seq []int) bool {
	for _, n := range seq {
		if n != 0 {
			return false
		}
	}
	return true
}

func extrapolateNext(seq []int) int {
	for !allZeroes(seq) {
		seq = differences(seq)
	}
	return seq[len(seq)-1] + seq[len(seq)-2]
}

func extrapolatePrev(seq []int) int {
	for !allZeroes(seq) {
		seq = differences(seq)
	}
	lastDiff := seq[len(seq)-1]
	for i := len(seq) - 1; i > 0; i-- {
		seq[i] = seq[i-1]
	}
	seq[0] = -lastDiff
	return seq[0] + seq[1]
}

func main() {
	inputData, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	var sumNext, sumPrev int
	for _, line := range inputData {
		sumNext += extrapolateNext(line)
		sumPrev += extrapolatePrev(line)
	}

	fmt.Println("Sum of extrapolated next values:", sumNext)
	fmt.Println("Sum of extrapolated previous values:", sumPrev)
}
