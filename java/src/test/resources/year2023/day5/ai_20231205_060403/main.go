
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Store initial seed values
	scanner.Scan()
	seedsLine := scanner.Text()
	seedsLine = strings.TrimPrefix(seedsLine, "seeds: ")
	seeds := convertToIntSlice(seedsLine)

	// Read mappings from the file
	maps := make(map[string][][3]int)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		mapName, mappingsLine := parts[0], parts[1]
		maps[mapName] = parseMappings(mappingsLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minLocation := findMinLocation(seeds, maps)

	fmt.Println(minLocation)
}

func convertToIntSlice(line string) []int {
	strNums := strings.Fields(line)
	nums := make([]int, len(strNums))
	for i, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = num
	}
	return nums
}

func parseMappings(mappingsLine string) [][3]int {
	parts := convertToIntSlice(mappingsLine)
	mappings := make([][3]int, 0, len(parts)/3)
	for i := 0; i < len(parts); i += 3 {
		mappings = append(mappings, [3]int{parts[i], parts[i+1], parts[i+2]})
	}
	return mappings
}

func findMapping(num int, mappings [][3]int) int {
	for _, mapping := range mappings {
		destStart, srcStart, length := mapping[0], mapping[1], mapping[2]
		if num >= srcStart && num < srcStart+length {
			return destStart + (num - srcStart)
		}
	}
	return num
}

func findMinLocation(seeds []int, maps map[string][][3]int) int {
	minLoc := int(^uint(0) >> 1) // Set to maximum int value
	for _, seed := range seeds {
		soil := findMapping(seed, maps["seed-to-soil map"])
		fertilizer := findMapping(soil, maps["soil-to-fertilizer map"])
		water := findMapping(fertilizer, maps["fertilizer-to-water map"])
		light := findMapping(water, maps["water-to-light map"])
		temperature := findMapping(light, maps["light-to-temperature map"])
		humidity := findMapping(temperature, maps["temperature-to-humidity map"])
		location := findMapping(humidity, maps["humidity-to-location map"])
		if location < minLoc {
			minLoc = location
		}
	}
	return minLoc
}
