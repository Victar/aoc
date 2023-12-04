
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer input.Close()

	var sum int64
	scanner := bufio.NewScanner(input)
	var lines [][]byte

	// First pass to read the lines
	for scanner.Scan() {
		line := scanner.Bytes()
		newLine := make([]byte, len(line))
		copy(newLine, line)
		lines = append(lines, newLine)
	}

	// Check for errors during Scan
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input file: %v\n", err)
		os.Exit(1)
	}

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == '.' {
				continue
			}
			if isSymbol(line[x]) {
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						ny, nx := y+dy, x+dx
						if ny >= 0 && ny < len(lines) && nx >= 0 && nx < len(lines[ny]) {
							if !isSymbol(lines[ny][nx]) && lines[ny][nx] != '.' {
								number := extractNumber(lines, ny, nx)
								n, err := strconv.ParseInt(number, 10, 64)
								if err == nil {
									sum += n
								}
								// Masking the used number to avoid double counting
								maskNumber(lines, ny, nx, len(number))
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)
}

func isSymbol(c byte) bool {
	return c == '*' || c == '#' || c == '+' || c == '$'
}

func extractNumber(lines [][]byte, y, x int) string {
	var number strings.Builder
	for x < len(lines[y]) && lines[y][x] >= '0' && lines[y][x] <= '9' {
		number.WriteByte(lines[y][x])
		x++
	}
	return number.String()
}

func maskNumber(lines [][]byte, y, x, length int) {
	for i := 0; i < length; i++ {
		lines[y][x+i] = '.'
	}
}
