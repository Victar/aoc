
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

	var platform [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		platform = append(platform, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	totalLoad := calcTotalLoad(platform)
	fmt.Println(totalLoad)
}

func calcTotalLoad(platform [][]rune) int {
	totalLoad := 0
	for col := range platform[0] {
		for row := len(platform) - 1; row >= 0; row-- {
			if platform[row][col] == 'O' {
				totalLoad += len(platform) - row
			}
		}
	}
	return totalLoad
}
