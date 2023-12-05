
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip the "seeds:" line
	seedsLine := strings.Fields(scanner.Text())

	seeds := make([]int, len(seedsLine))
	for i, s := range seedsLine {
		seeds[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, nil, err
		}
	}

	scanner.Scan() // Skip the empty line before mappings
	var mappings [][]int
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) == 0 {
			break
		}
		if strings.Contains(line[0], "-to-") { // it's a header line
			continue
		}
		mapping := make([]int, len(line))
		for i, s := range line {
			mapping[i], err = strconv.Atoi(s)
			if err != nil {
				return nil, nil, err
			}
		}
		mappings = append(mappings, mapping)
	}
	return seeds, mappings, scanner.Err()
}

func mapValues(seeds []int, mappings [][]int) []int {
	for _, m := range mappings {
		for i, seed := range seeds {
			dStart := m[0]
			sStart := m[1]
			length := m[2]

			if seed >= sStart && seed < sStart+length {
				seeds[i] = dStart + (seed - sStart)
			}
		}
	}
	return seeds
}

func main() {
	seeds, mappings, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	lowestLocation := mapValues(seeds, mappings)

	min := lowestLocation[0]
	for _, location := range lowestLocation {
		if location < min {
			min = location
		}
	}
	fmt.Println(min)
}
