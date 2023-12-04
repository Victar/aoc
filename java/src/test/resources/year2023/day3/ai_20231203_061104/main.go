
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	schematic := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}

	sum := 0
	rows := len(schematic)
	cols := len(schematic[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if isSymbol(schematic[i][j]) {
				sum += sumAdjacentNumbers(schematic, i, j, rows, cols)
			}
		}
	}

	fmt.Println(sum)
}

func isSymbol(r rune) bool {
	return r != '.' && !unicode.IsDigit(r)
}

func sumAdjacentNumbers(schematic [][]rune, x, y, rows, cols int) int {
	offsets := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	sum := 0

	for _, offset := range offsets {
		newX, newY := x+offset.dx, y+offset.dy
		if newX >= 0 && newX < rows && newY >= 0 && newY < cols && unicode.IsDigit(schematic[newX][newY]) {
			number, err := strconv.Atoi(string(schematic[newX][newY]))
			if err == nil {
				sum += number
			}
			schematic[newX][newY] = '.' // Mark as visited
		}
	}

	return sum
}

func init() {
	if len(os.Args) > 1 && os.Args[1] == "3/input" {
		filename := os.Args[1]
		os.Rename(filename, "input.txt")
	}
}
