
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	var times, distances []int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		// Extract time and distance from the line
		time, _ := strconv.Atoi(parts[1])
		distance, _ := strconv.Atoi(parts[3])
		times = append(times, time)
		distances = append(distances, distance)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Compute the number of ways to beat the record for each race
	product := 1
	for i := 0; i < len(times); i++ {
		ways := 0
		for chargeTime := 1; chargeTime < times[i]; chargeTime++ {
			travelTime := times[i] - chargeTime
			if distance := chargeTime * travelTime; distance > distances[i] {
				ways++
			}
		}
		product *= ways
	}

	// Print the final product
	fmt.Println(product)
}
