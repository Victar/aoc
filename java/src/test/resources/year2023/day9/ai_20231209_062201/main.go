
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := make([]int, 0)
		for _, numStr := range strings.Fields(line) {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		sum += extrapolateNextValue(numbers)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func extrapolateNextValue(nums []int) int {
	for {
		diffs := make([]int, len(nums)-1)
		zeroes := true
		for i := 0; i < len(nums)-1; i++ {
			diffs[i] = nums[i+1] - nums[i]
			if diffs[i] != 0 {
				zeroes = false
			}
		}
		if zeroes {
			break
		}
		nums = diffs
	}
	return nums[len(nums)-1] + nums[len(nums)-2]
}
