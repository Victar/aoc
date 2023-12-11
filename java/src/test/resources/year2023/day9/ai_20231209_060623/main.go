package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
func extrapolateNext(seq Sequence) int {
	next := seq[len(seq)-1]

	for !isAllZeros(seq) {
		seq = computeDifferences(seq)
	}

	return next + seq[0]
}

// sumOfNextValues computes the sum of the next values in all sequences.
func sumOfNextValues(sequences []Sequence) int {
	sum := 0
	for _, seq := range sequences {
		sum += extrapolateNext(seq)
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sequences []Sequence

	for scanner.Scan() {
		line := scanner.Text()
		seq := parseLine(line)
		sequences = append(sequences, seq)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	answer := sumOfNextValues(sequences)
	fmt.Println(answer)
}
