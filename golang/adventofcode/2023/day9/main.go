package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
	"strings"
)

var DAY = "9"

func main() {
	runAny(false)
	runAny(true)
}

func runAny(isGold bool) {

	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var sequences []Sequence

	for _, line := range lines {
		seq := parseLine(line)
		sequences = append(sequences, seq)
	}

	answer := sumOfNextValues(sequences, isGold)
	fmt.Println(answer)
}

// Sequence represents an integer slice.
type Sequence []int

// Parses a line to a Sequence.
func parseLine(line string) Sequence {
	fields := strings.Fields(line)
	seq := make(Sequence, len(fields))
	for i, f := range fields {
		val, err := strconv.Atoi(f)
		if err != nil {
			fmt.Printf("Failed to parse value: %s\n", f)
			panic(err)
		}
		seq[i] = val
	}
	return seq
}

// computeDifferences takes a Sequence and returns a sequence of differences
// between consecutive elements.
func computeDifferences(seq Sequence) Sequence {
	differences := make(Sequence, len(seq)-1)
	for i := 1; i < len(seq); i++ {
		differences[i-1] = seq[i] - seq[i-1]
	}
	return differences
}

// isAllZeros checks if all the elements in the Sequence are zero.
func isAllZeros(seq Sequence) bool {
	for _, v := range seq {
		if v != 0 {
			return false
		}
	}
	return true
}

// extrapolateNext computes the next number in the sequence based on differences.
func extrapolateNext(seq Sequence, isGold bool) int {
	if isGold {
		reverseSlice(seq)
	}
	next := seq[len(seq)-1]
	for !isAllZeros(seq) {
		seq = computeDifferences(seq)
		next += seq[len(seq)-1]
	}
	return next
}
func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// sumOfNextValues computes the sum of the next values in all sequences.
func sumOfNextValues(sequences []Sequence, isGold bool) int {
	sum := 0
	for _, seq := range sequences {
		sum += extrapolateNext(seq, isGold)
	}
	return sum
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
