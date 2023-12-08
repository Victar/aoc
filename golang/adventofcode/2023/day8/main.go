package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

var DAY = "8"

func main() {
	runSilver()
	runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	instructions := ""
	nodes := make(map[string][2]string)
	for _, line := range lines {
		if instructions == "" {
			instructions = line
		} else if len(line) > 1 {

			parts := strings.Fields(line)
			nodes[parts[0]] = [2]string{parts[2][1:4], parts[3][:3]}
		}
	}
	currentNode := "AAA"
	steps, instructionLength := 0, len(instructions)

	for currentNode != "ZZZ" {
		direction := instructions[steps%instructionLength]
		if direction == 'R' {
			currentNode = nodes[currentNode][1]
		} else if direction == 'L' {
			currentNode = nodes[currentNode][0]
		}
		steps++
	}
	fmt.Println(steps)
}

func runGold() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	instructions := ""
	nodes := make(map[string][2]string)
	var nodesToCheck []string
	var steps []int

	for _, line := range lines {
		if instructions == "" {
			instructions = line
		} else if len(line) > 1 {
			parts := strings.Fields(line)
			nodes[parts[0]] = [2]string{parts[2][1:4], parts[3][:3]}
			if parts[0][2] == 'A' {
				nodesToCheck = append(nodesToCheck, parts[0])
			}
		}
	}
	for _, nodesToCheck := range nodesToCheck {
		steps = append(steps, countSteps(instructions, nodesToCheck, nodes))
	}
	fmt.Println(findLCM(steps))
}

func countSteps(instructions string, currentNode string, nodes map[string][2]string) int {
	steps, instructionLength := 0, len(instructions)
	for currentNode[2] != 'Z' {
		direction := instructions[steps%instructionLength]
		if direction == 'R' {
			currentNode = nodes[currentNode][1]
		} else if direction == 'L' {
			currentNode = nodes[currentNode][0]
		}
		steps++
	}
	return steps
}

// Function to find the greatest common divisor (GCD) using Euclidean algorithm
func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the least common multiple (LCM) of a slice of integers
func findLCM(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		gcd := findGCD(lcm, nums[i])
		lcm = (lcm * nums[i]) / gcd
	}
	return lcm
}
