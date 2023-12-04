
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
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	totalPoints := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			fmt.Println("Invalid data format:", line)
			return
		}

		winNums := strings.Fields(parts[0])
		playerNums := strings.Fields(parts[1])
		winMap := make(map[string]bool)
		for _, n := range winNums {
			winMap[n] = true
		}

		cardPoints := 0
		for _, n := range playerNums {
			if winMap[n] {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}
		
		totalPoints += cardPoints
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(totalPoints)
}
