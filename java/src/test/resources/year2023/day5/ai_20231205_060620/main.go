
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
		fmt.Println(err)
		return
	}
	defer file.Close()

	var seeds []int
	mappings := make([]map[int]int, 6)

	scanner := bufio.NewScanner(file)
	scanner.Scan() // scan the seed line
	seeds = parseLine(scanner.Text()[7:])

	for i := range mappings {
		scanner.Scan() // scan the map description line
		currentLine := scanner.Text()[strings.Index(scanner.Text(), ":")+2:]
		mappings[i] = make(map[int]int)
		nums := parseLine(currentLine)
		for j := 0; j < len(nums); j += 3 {
			for k := 0; k < nums[j+2]; k++ {
				mappings[i][nums[j+1]+k] = nums[j]+k
			}
		}
	}

	lowestLocation := -1
	for _, seed := range seeds {
		location := getLocationForSeed(seed, mappings)
		if lowestLocation == -1 || location < lowestLocation {
			lowestLocation = location
		}
	}
	fmt.Println(lowestLocation)
}

func parseLine(line string) []int {
	strings := strings.Split(line, " ")
	nums := make([]int, len(strings))
	for i, str := range strings {
		nums[i], _ = strconv.Atoi(str)
	}
	return nums
}

func getLocationForSeed(seed int, mappings []map[int]int) int {
	stages := [6]string{"seed", "soil", "fertilizer", "water", "light", "temperature"}
	current := seed
	for _, stage := range stages {
		if newNumber, ok := mappings[categoryIndex(stage)][current]; ok {
			current = newNumber
		} else {
			current = current // the number remains the same if not mapped
		}
	}
	// last mapping from temperature to humidity
	if newNumber, ok := mappings[5][current]; ok {
		current = newNumber
	} else {
		current = current
	}
	// then from humidity to location
	return current
}

func categoryIndex(category string) int {
	switch category {
	case "seed":
		return 0
	case "soil":
		return 1
	case "fertilizer":
		return 2
	case "water":
		return 3
	case "light":
		return 4
	case "temperature":
		return 5
	default:
		return -1
	}
}
