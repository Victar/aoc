
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (int, int, int) {
	parts := strings.Fields(line)
	start, _ := strconv.Atoi(parts[0])
	source, _ := strconv.Atoi(parts[1])
	length, _ := strconv.Atoi(parts[2])
	return start, source, length
}

func mapValue(maps []string, value int) int {
	for _, m := range maps {
		start, source, length := parseLine(m)
		if value >= source && value < source+length {
			return start + (value - source)
		}
	}
	return value // If not mapped, the source number corresponds to the same destination number
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan() // Read seeds line
	seedsLine := strings.TrimPrefix(scanner.Text(), "seeds: ")
	seedsStr := strings.Fields(seedsLine)
	seeds := make([]int, len(seedsStr))
	for i, seed := range seedsStr {
		seeds[i], _ = strconv.Atoi(seed)
	}

	maps := make([][]string, 7)
	for i := range maps {
		scanner.Scan() // Read the map title
		scanner.Scan() // Read the first map line
		maps[i] = append(maps[i], scanner.Text())
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			maps[i] = append(maps[i], line)
		}
	}

	minLocation := -1

	for _, seed := range seeds {
		soil := mapValue(maps[0], seed)
		fertilizer := mapValue(maps[1], soil)
		water := mapValue(maps[2], fertilizer)
		light := mapValue(maps[3], water)
		temperature := mapValue(maps[4], light)
		humidity := mapValue(maps[5], temperature)
		location := mapValue(maps[6], humidity)
		if minLocation == -1 || location < minLocation {
			minLocation = location
		}
	}

	fmt.Println(minLocation)
}

