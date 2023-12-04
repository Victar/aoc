package main

import (
	"adventofcode/util"
	"fmt"
	"strconv"
)

var DAY = "3"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var sum int64
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
								number, start := extractNumber(lines, ny, nx)
								n, err := strconv.ParseInt(number, 10, 64)
								if err == nil {
									sum += n
								}
								maskNumber(lines, ny, start, len(number))
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var sum int64
	for y, line := range lines {
		for x := 0; x < len(line); x++ {

			if line[x] == '.' {
				continue
			}
			if line[x] == '*' {
				var nums []int64
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						ny, nx := y+dy, x+dx
						if ny >= 0 && ny < len(lines) && nx >= 0 && nx < len(lines[ny]) {
							if !isSymbol(lines[ny][nx]) && lines[ny][nx] != '.' {
								number, start := extractNumber(lines, ny, nx)
								n, err := strconv.ParseInt(number, 10, 64)
								if err == nil {
									nums = append(nums, n)
								}
								maskNumber(lines, ny, start, len(number))
							}
						}
					}
				}
				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}
	fmt.Println(sum)
}

func isSymbol(c byte) bool {
	return c != '.' && !('0' <= c && c <= '9')
}

func extractNumber(lines []string, y, x int) (result string, start int) {
	var number = ""
	xx := x
	start = x
	for x >= 0 && lines[y][x] >= '0' && lines[y][x] <= '9' {
		number = string(lines[y][x]) + number
		x--
		start = x
	}
	start++

	x = xx + 1
	for x < len(lines[y]) && lines[y][x] >= '0' && lines[y][x] <= '9' {
		number = number + string(lines[y][x])
		x++
	}
	return number, start
}

func maskNumber(lines []string, y, x, length int) {
	for i := 0; i < length; i++ {
		lineChars := []rune(lines[y])
		lineChars[x+i] = '.'
		lines[y] = string(lineChars)
	}
}
