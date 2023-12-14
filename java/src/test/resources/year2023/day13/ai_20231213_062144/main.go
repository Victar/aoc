
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if vertical := findVerticalReflection(line); vertical >= 0 {
			sum += vertical
		} else {
			horizontal, err := findHorizontalReflection(file)
			if err != nil {
				panic(err)
			}
			sum += 100 * horizontal
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(sum)
}

func findVerticalReflection(pattern string) int {
	for i, j := 0, len(pattern)-1; i < j; i, j = i+1, j-1 {
		if pattern[i] != pattern[j] {
			return -1
		}
	}
	return len(pattern) / 2
}

func findHorizontalReflection(file *os.File) (int, error) {
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	size := len(lines[0])
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		for k := 0; k < size; k++ {
			if lines[i][k] != lines[j][k] {
				return -1, nil
			}
		}
	}
	return len(lines) / 2, scanner.Err()
}
