package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Condition struct {
	Unknowns      []int
	DamagedGroups []int
	Arrangements  int
}

func main() {
	input, err := os.Open("/Users/vkad2506/AdventOfCode/java/src/test/resources/year2023/day12/input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	totalArrangements := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		unknowns, damaged := parseLine(parts[0], parts[1])
		condition := Condition{Unknowns: unknowns, DamagedGroups: damaged}
		condition.solve()
		totalArrangements += condition.Arrangements
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	fmt.Println(totalArrangements)
}

func parseLine(springs, groups string) ([]int, []int) {
	unknownGroups := strings.Split(groups, ",")
	numGroups := make([]int, len(unknownGroups))

	for i, g := range unknownGroups {
		numGroups[i], _ = strconv.Atoi(g)
	}

	unknowns := countUnknowns(springs)
	return unknowns, numGroups
}

func countUnknowns(springs string) []int {
	unknownCounts := []int{}
	count := 0
	for _, r := range springs {
		switch r {
		case '?':
			count++
		case '.', '#':
			if count > 0 {
				unknownCounts = append(unknownCounts, count)
				count = 0
			}
		}
	}
	if count > 0 {
		unknownCounts = append(unknownCounts, count)
	}
	return unknownCounts
}

func (c *Condition) solve() {
	c.recurse(0, c.DamagedGroups, []int{})
}

func (c *Condition) recurse(groupIndex int, remainingGroups []int, groupSizes []int) {
	if groupIndex >= len(c.Unknowns) {
		if len(remainingGroups) == 0 {
			valid := true
			for i := 0; i < len(groupSizes)-1; i++ {
				if groupSizes[i+1]-groupSizes[i] <= 1 {
					valid = false
				}
			}
			if valid {
				c.Arrangements++
			}
		}
		return
	}

	size := c.Unknowns[groupIndex]
	for i := 0; i < len(remainingGroups); i++ {
		if size >= remainingGroups[i] {
			nextGroups := make([]int, len(remainingGroups)-(i+1))
			copy(nextGroups, remainingGroups[i+1:])
			newGroupSizes := append([]int(nil), groupSizes...)
			for j := remainingGroups[i]; j <= size; j++ {
				newGroupSizes = append(newGroupSizes, j)
				c.recurse(groupIndex+1, nextGroups, newGroupSizes)
				if j < size {
					newGroupSizes = append(newGroupSizes, j+1)
				}
			}
			break
		}
	}
}
