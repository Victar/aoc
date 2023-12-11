
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func computeNextValue(vals []int) int {
	diff := make([][]int, 0)
	diff = append(diff, vals)

	for {
		lastSeq := diff[len(diff)-1]
		newSeq := make([]int, 0)

		for i := 1; i < len(lastSeq); i++ {
			newSeq = append(newSeq, lastSeq[i]-lastSeq[i-1])
		}

		if allZeros(newSeq) {
			break
		}
		diff = append(diff, newSeq)
	}

	for i := len(diff) - 2; i >= 0; i-- {
		last := diff[i][len(diff[i])-1]
		diff[i] = append(diff[i], last+diff[i+1][len(diff[i+1])-1])
	}

	return diff[0][len(diff[0])-1]
}

func computePrevValue(vals []int) int {
	diff := make([][]int, 0)
	diff = append(diff, vals)

	for {
		lastSeq := diff[len(diff)-1]
		newSeq := make([]int, 0)

		for i := 1; i < len(lastSeq); i++ {
			newSeq = append(newSeq, lastSeq[i]-lastSeq[i-1])
		}

		if allZeros(newSeq) {
			break
		}
		diff = append(diff, newSeq)
	}

	for i := len(diff) - 2; i >= 0; i-- {
		first := diff[i][0]
		diff[i] = append([]int{first - diff[i+1][0]}, diff[i]...)
	}

	return diff[0][0]
}

func allZeros(seq []int) bool {
	for _, num := range seq {
		if num != 0 {
			return false
		}
	}
	return true
}

func processLine(line string) (int, int) {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, f := range fields {
		n, err := strconv.Atoi(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error converting string to int:", err)
			return 0, 0
		}
		nums[i] = n
	}
	return computeNextValue(nums), computePrevValue(nums)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sumNext := 0
	sumPrev := 0
	for scanner.Scan() {
		next, prev := processLine(scanner.Text())
		sumNext += next
		sumPrev += prev
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading file:", err)
		return
	}

	fmt.Println("Sum of next values:", sumNext)
	fmt.Println("Sum of previous values:", sumPrev)
}
