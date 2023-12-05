
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type MapRange struct {
	destStart int
	srcStart  int
	length    int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Read the first line, which is not needed for the solution

	var seeds []int
	maps := make([][]MapRange, 6) // There are 6 maps in total

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "seeds:") {
			seeds = parseIntList(line)
		} else if strings.HasSuffix(line, "map:") {
			currentMap := readMap(scanner)
			maps = append(maps, currentMap)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}

	lowestLocation := findLowestLocation(seeds, maps)
	println(lowestLocation)
}

func readMap(scanner *bufio.Scanner) []MapRange {
	var mapRange []MapRange
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break // Empty line means end of the current map
		}
		parts := strings.Split(line, " ")
		destStart, _ := strconv.Atoi(parts[0])
		srcStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])

		mapRange = append(mapRange, MapRange{destStart, srcStart, length})
	}
	return mapRange
}

func applyMapping(value int, mapRange []MapRange) int {
	for _, m := range mapRange {
		if value >= m.srcStart && value < m.srcStart+m.length {
			offset := value - m.srcStart
			return m.destStart + offset
		}
	}
	return value
}

func findLowestLocation(seeds []int, maps [][]MapRange) int {
	lowest := -1
	for _, seed := range seeds {
		location := seed
		for _, mapping := range maps {
			location = applyMapping(location, mapping)
		}
		if lowest == -1 || location < lowest {
			lowest = location
		}
	}
	return lowest
}

func parseIntList(line string) []int {
	line = strings.Replace(line, "seeds:", "", 1)
	parts := strings.Fields(line)
	var nums []int
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		nums = append(nums, num)
	}
	return nums
}
