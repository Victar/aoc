
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
		log.Fatalf("unable to open input file: %v", err)
	}
	defer file.Close()

	var waysProducts int64 = 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatalf("invalid input format: %s", line)
		}
		time, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatalf("invalid time value: %v", err)
		}
		distance, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatalf("invalid distance value: %v", err)
		}

		var ways int64 = 0
		for holdTime := int64(1); holdTime < time; holdTime++ {
			travelTime := time - holdTime
			traveled := holdTime * travelTime
			if traveled > distance {
				ways++
			}
		}
		waysProducts *= ways
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(waysProducts)
}
