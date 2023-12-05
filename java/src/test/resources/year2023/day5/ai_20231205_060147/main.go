
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt:", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var seedNumbers []int
	var mappings = make(map[string][][3]int)
	var mapOrder = []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ": ")
			mapName := parts[0]
			numbersStr := strings.Fields(parts[1])
			numbers := make([]int, 0, len(numbersStr))
			for _, n := range numbersStr {
				num, _ := strconv.Atoi(n)
				numbers = append(numbers, num)
			}
			if mapName == "seeds" {
				seedNumbers = numbers
			} else {
				for i := 0; i < len(numbers); i += 3 {
					mappings[mapName] = append(mappings[mapName], [3]int{numbers[i], numbers[i+1], numbers[i+2]})
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	transform := func(number int, mapName string) int {
		for _, v := range mappings[mapName] {
			destStart, sourceStart, length := v[0], v[1], v[2]
			if number >= sourceStart && number < sourceStart+length {
				return destStart + (number - sourceStart)
			}
		}
		return number // unchanged if not mapped
	}

	var minLocation = int(^uint(0) >> 1) // max possible int value
	for _, seedNumber := range seedNumbers {
		location := seedNumber
		for _, m := range mapOrder {
			location = transform(location, m)
		}
		if location < minLocation {
			minLocation = location
		}
	}

	fmt.Println(minLocation)
}
