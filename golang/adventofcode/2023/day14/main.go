package main

import (
	"adventofcode/util"
	"fmt"
)

var DAY = "14"

func main() {
	runSilver()
	runGold()
}
func runSilver() {
	grid, err := util.ReadInput("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	println(rollRocksCycle(grid, false))
}

func runGold() {
	grid, err := util.ReadInput("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	combinationMap := make(map[string]int)
	combinationVal := []int{}
	cycleLength := 0
	cycleStart := 0
	keepSearch := true
	for i := 0; keepSearch; i++ {
		load := rollRocksCycle(grid, true)
		combination := gridToString(grid)
		cl, found := combinationMap[combination]
		if found {
			cycleStart = cl
			cycleLength = i - cl
			keepSearch = false
		} else {
			combinationMap[combination] = i
			combinationVal = append(combinationVal, load)
		}
	}
	fmt.Println(combinationVal[(1000000000-cycleStart-1)%cycleLength+cycleStart])
}

func gridToString(grid [][]rune) string {
	var result string
	for _, row := range grid {
		result += string(row)
	}
	return result
}
func printGrid(runesGrid [][]rune) {
	rows := len(runesGrid)
	for col := 0; col < rows; col++ {
		println(string(runesGrid[col]))
	}
}

// north, then west, then south, then east.
func rollRocksCycle(runesGrid [][]rune, isGold bool) int {
	height := len(runesGrid)
	width := len(runesGrid[0])
	load := 0
	//North
	for col := 0; col < width; col++ {
		for row := 0; row < height; row++ {
			currentCell := runesGrid[row][col]
			if currentCell == 'O' {
				targetRow := row
				for targetRow-1 >= 0 && runesGrid[targetRow-1][col] == '.' {
					targetRow--
				}
				if targetRow != row {
					runesGrid[row][col] = '.'
					runesGrid[targetRow][col] = 'O'
				}
			}
		}
		load = calculateLoad(runesGrid)
	}
	// West
	for col := 0; col < width; col++ {
		for row := 0; row < height; row++ {
			currentCell := runesGrid[row][col]
			if currentCell == 'O' {
				targetCol := col
				for targetCol-1 >= 0 && runesGrid[row][targetCol-1] == '.' {
					targetCol--
				}
				if targetCol != col {
					runesGrid[row][col] = '.'
					runesGrid[row][targetCol] = 'O'
				}
			}
		}
	}
	// South
	for col := 0; col < width; col++ {
		for row := height - 1; row >= 0; row-- {
			currentCell := runesGrid[row][col]
			if currentCell == 'O' {
				targetRow := row
				for targetRow+1 <= height-1 && runesGrid[targetRow+1][col] == '.' {
					targetRow++
				}
				if targetRow != row {
					runesGrid[row][col] = '.'
					runesGrid[targetRow][col] = 'O'
				}
			}
		}
	}
	// East
	for col := width - 1; col >= 0; col-- {
		for row := 0; row < height; row++ {
			currentCell := runesGrid[row][col]
			if currentCell == 'O' {
				targetCol := col
				for targetCol+1 <= width-1 && runesGrid[row][targetCol+1] == '.' {
					targetCol++
				}
				if targetCol != col {
					runesGrid[row][col] = '.'
					runesGrid[row][targetCol] = 'O'
				}
			}
		}
	}
	if isGold {
		load = calculateLoad(runesGrid)
	}
	return load
}

func calculateLoad(runesGrid [][]rune) int {
	load := 0
	height := len(runesGrid[0])

	for row, line := range runesGrid {
		for _, cell := range line {
			if cell == 'O' {
				// Load is the distance from the bottom plus one (the row the rock is on)
				load += (height - row)
			}
		}
	}

	return load
}
