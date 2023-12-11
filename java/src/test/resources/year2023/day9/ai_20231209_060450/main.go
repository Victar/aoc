
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filepath string) ([][]int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		numbers := make([]int, len(fields))

		for i, f := range fields {
			numbers[i], err = strconv.Atoi(f)
			if err != nil {
				return nil, err
			}
		}

		data = append(data, numbers)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func findNext(history []int) int {
	for {
		differences := make([]int, len(history)-1)
		allZeroes := true
		for i := 1; i < len(history); i++ {
			differences[i-1] = history[i] - history[i-1]
			if differences[i-1] != 0 {
				allZeroes = false
			}
		}
		if allZeroes {
			break
		}
		history = differences
	}
	return history[len(history)-1] + history[len(history)-2]
}

func main() {
	histories, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum int
	for _, history := range histories {
		sum += findNext(history)
	}

	fmt.Println(sum)
}
