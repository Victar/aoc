package main

import (
	"adventofcode/util"
)

var DAY = "25"

var snafuToDigits = map[byte]int{
	'1': 1,
	'2': 2,
	'0': 0,
	'-': -1,
	'=': -2,
}
var digitsToSnafu = map[int]byte{
	1:  '1',
	2:  '2',
	0:  '0',
	-1: '-',
	-2: '=',
}

func main() {
	runSilver()
}

func runSilver() {
	lines, err := util.ReadFile("year2022/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var answer = 0
	for _, line := range lines {
		answer += snafuToNumber(line)
	}
	println(numberToSnafu(answer))

}

func numberToSnafu(number int) string {
	var answer []byte
	for number > 0 {
		var reminder = (number+2)%5 - 2
		answer = append([]byte{digitsToSnafu[reminder]}, answer...)
		number = (number - reminder) / 5
	}
	return string(answer)
}

func snafuToNumber(input string) int {
	var answer = 0
	var pow = 1
	for i := len(input) - 1; i >= 0; i-- {
		answer += snafuToDigits[input[i]] * pow
		pow *= 5
	}
	return answer
}
