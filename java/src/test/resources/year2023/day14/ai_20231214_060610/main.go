package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day14/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := [][]rune{}
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, ch := range line {
			row[i] = ch
		}
		rows = append(rows, row)
		rowCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalLoad := 0
	for _, row := range rows {
		for _, ch := range row {
			if ch == 'O' {
				totalLoad += rowCount
			}
		}
		rowCount--
	}

	fmt.Println(totalLoad)
}
