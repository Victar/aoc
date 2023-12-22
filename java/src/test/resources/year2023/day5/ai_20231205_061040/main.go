package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, map[string][][3]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Read seed numbers
	seedsRaw := strings.Fields(strings.TrimPrefix(scanner.Text(), "seeds: "))
	seeds := make([]int, len(seedsRaw))
	for i, seed := range seedsRaw {
		seeds[i], _ = strconv.Atoi(seed)
	}

	maps := map[string][][3]int{}

	currentMap := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "map:") {
			currentMap = line
			maps[currentMap] = [][3]int{}
		} else {
			fields := strings.Fields(line)
			if len(fields) != 3 {
				continue
			}
			startDest, _ := strconv.Atoi(fields[0])
			startSrc, _ := strconv.Atoi(fields[1])
			length, _ := strconv.Atoi(fields[2])

			maps[currentMap] = append(maps[currentMap], [3]int{startDest, startSrc, length})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return seeds, maps
}

func applyMapping(value int, mapping [][3]int) int {
	if mapping == nil {
		return value
	}

	for _, m := range mapping {
		destStart, srcStart, length := m[0], m[1], m[2]
		if value >= srcStart && value < srcStart+length {
			return destStart + (value - srcStart)
		}
	}
	return value
}

func main() {
	seeds, maps := readInput("input.txt")

	mapsOrder := []string{
		"seed-to-soil map:",
		"soil-to-fertilizer map:",
		"fertilizer-to-water map:",
		"water-to-light map:",
		"light-to-temperature map:",
		"temperature-to-humidity map:",
		"humidity-to-location map:",
	}

	minLocation := -1
	for _, seed := range seeds {
		current := seed
		for _, mapName := range mapsOrder {
			current = applyMapping(current, maps[mapName])
		}
		if minLocation == -1 || current < minLocation {
			minLocation = current
		}
	}

	fmt.Println(minLocation)
}
