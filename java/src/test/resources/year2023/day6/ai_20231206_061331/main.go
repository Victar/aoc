
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

	var ways int = 1 // Start with 1 as we need to multiply the ways for each race
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		time, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		record, err := strconv.Atoi(fields[3])
		if err != nil {
			panic(err)
		}
		ways *= calculateWays(time, record)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(ways)
}

func calculateWays(time, record int) int {
	count := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		distance := holdTime * (time - holdTime)
		if distance > record {
			count++
		}
	}
	return count
}
