
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
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalWays int64 = 1 // We will multiply the number of ways, so we start with 1.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 3 {
			fmt.Println("Invalid line format")
			continue
		}

		time, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		distance, err := strconv.ParseInt(parts[3], 10, 64)
		if err != nil {
			panic(err)
		}

		var ways int64
		for holdingTime := time; holdingTime >= 0; holdingTime-- {
			speed := holdingTime
			travelTime := time - holdingTime
			currentDistance := speed * travelTime
			if currentDistance > distance {
				ways += travelTime
			}
		}
		totalWays *= ways
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(totalWays)
}
