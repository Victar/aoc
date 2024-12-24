package main

import (
	"adventofcode/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var DAY = "24"

func main() {
	runSilver()
	//runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	var operations [][]string
	values := make(map[string]int)
	for _, line := range lines {
		if strings.Contains(line, " -> ") {
			parts := strings.Fields(line)
			operations = append(operations, parts)
		}
		if strings.Contains(line, ": ") {
			split := strings.Split(line, ": ")
			values[split[0]], _ = strconv.Atoi(split[1])
		}
	}
	runOperations(values, operations)
	output := calcSilver(values)
	//printMap(values)
	fmt.Println(output)

}
func printMap(m map[string]int) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}

func runOperations(values map[string]int, operations [][]string) {
	for len(operations) > 0 {
		var restOperation [][]string
		for _, parts := range operations {
			input1, exist1 := values[parts[0]]
			input2, exist2 := values[parts[2]]
			if exist1 && exist2 {
				dest := parts[4]
				switch parts[1] {
				case "AND":
					values[dest] = input1 & input2
				case "OR":
					values[dest] = input1 | input2
				case "XOR":
					values[dest] = input1 ^ input2
				}
			} else {
				restOperation = append(restOperation, parts)
			}
		}
		operations = restOperation
	}
}

func calcSilver(values map[string]int) int {
	var result int
	for i := 0; ; i++ {
		key := fmt.Sprintf("z%02d", i)
		value, found := values[key]
		if !found {
			break
		}
		result = result | value<<i
	}
	return result
}
func runGold() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		println(line)
	}
}
