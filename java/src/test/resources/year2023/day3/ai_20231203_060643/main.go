
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read file "input.txt" line by line and store it in a slice of strings.
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var schematic [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		return
	}

	// Function to check if a rune is a symbol other than a period.
	isSymbol := func(r rune) bool {
		return r != '.' && !('0' <= r && r <= '9')
	}

	// Calculate the sum of all part numbers.
	sum := 0
	for y, row := range schematic {
		for x, cell := range row {
			if '0' <= cell && cell <= '9' {
				// Check adjacent cells for symbols.
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 {
							continue
						}
						adjY, adjX := y+dy, x+dx
						// Check if adjacent cell is within the bounds of the schematic.
						if adjY >= 0 && adjY < len(schematic) && adjX >= 0 && adjX < len(row) {
							if isSymbol(schematic[adjY][adjX]) {
								// Extract the full number.
								numStr := ""
								for nx := x; nx < len(row) && ('0' <= row[nx] && row[nx] <= '9'); nx++ {
									numStr += string(row[nx])
									row[nx] = '.' // Mark as visited.
								}
								num, _ := strconv.Atoi(numStr)
								sum += num
								break
							}
						}
					}
				}
			}
		}
	}

	// Print the answer.
	fmt.Println(sum)
}
