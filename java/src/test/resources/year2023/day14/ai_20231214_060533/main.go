
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculateTotalLoad(rows))
}

func calculateTotalLoad(rows []string) int {
	n := len(rows)
	totalLoad := 0

	for i, row := range rows {
		for j := 0; j < len(row); j++ {
			if row[j] == 'O' {
				totalLoad += n - i
			}
		}
	}

	return totalLoad
}
