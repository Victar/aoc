
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
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sequences [][]int
	for scanner.Scan() {
		line := scanner.Text()
		numStrings := strings.Split(line, " ")
		nums := make([]int, 0, len(numStrings))
		for _, numString := range numStrings {
			num, err := strconv.Atoi(numString)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			nums = append(nums, num)
		}
		sequences = append(sequences, nums)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	sum := 0
	for _, seq := range sequences {
		nextValue := getNextValue(seq)
		sum += nextValue
	}

	fmt.Println(sum)
}

func getNextValue(seq []int) int {
	for {
		diffSeq := getDifferenceSequence(seq)
		if allZeros(diffSeq) {
			break
		}
		seq = diffSeq
	}
	return seq[len(seq)-1] + seq[len(seq)-2]
}

func getDifferenceSequence(seq []int) []int {
	diffSeq := make([]int, len(seq)-1)
	for i := 1; i < len(seq); i++ {
		diffSeq[i-1] = seq[i] - seq[i-1]
	}
	return diffSeq
}

func allZeros(slice []int) bool {
	for _, n := range slice {
		if n != 0 {
			return false
		}
	}
	return true
}
