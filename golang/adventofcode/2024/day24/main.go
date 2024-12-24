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
	runGold()
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
	fmt.Println(output)

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
	divider := 0
	for i, line := range lines {
		if line == "" {
			divider = i
			break
		}
	}
	configurations := lines[divider+1:]
	swaps := checkAdder(configurations)
	sort.Strings(swaps)
	fmt.Println(strings.Join(swaps, ","))
}

func checkAdder(configurations []string) []string {
	var currentCarryWire string
	var swaps []string
	bit := 0

	for {
		xWire := fmt.Sprintf("x%02d", bit)
		yWire := fmt.Sprintf("y%02d", bit)
		zWire := fmt.Sprintf("z%02d", bit)

		if bit == 0 {
			currentCarryWire = findGate(xWire, yWire, "AND", configurations)
		} else {
			abXorGate := findGate(xWire, yWire, "XOR", configurations)
			abAndGate := findGate(xWire, yWire, "AND", configurations)

			cinAbXorGate := findGate(abXorGate, currentCarryWire, "XOR", configurations)
			if cinAbXorGate == "" {
				swaps = append(swaps, abXorGate, abAndGate)
				configurations = swapOutputWires(abXorGate, abAndGate, configurations)
				bit = 0
				continue
			}

			if cinAbXorGate != zWire {
				swaps = append(swaps, cinAbXorGate, zWire)
				configurations = swapOutputWires(cinAbXorGate, zWire, configurations)
				bit = 0
				continue
			}

			cinAbAndGate := findGate(abXorGate, currentCarryWire, "AND", configurations)
			carryWire := findGate(abAndGate, cinAbAndGate, "OR", configurations)
			currentCarryWire = carryWire
		}

		bit++
		if bit >= 45 {
			break
		}
	}
	return swaps
}

func findGate(xWire, yWire, gateType string, configurations []string) string {
	subStrA := fmt.Sprintf("%s %s %s -> ", xWire, gateType, yWire)
	subStrB := fmt.Sprintf("%s %s %s -> ", yWire, gateType, xWire)

	for _, config := range configurations {
		if strings.Contains(config, subStrA) || strings.Contains(config, subStrB) {
			return strings.Split(config, " -> ")[1]
		}
	}
	return ""
}

func swapOutputWires(wireA, wireB string, configurations []string) []string {
	var newConfigurations []string
	for _, config := range configurations {
		parts := strings.Split(config, " -> ")
		inputWires, outputWire := parts[0], parts[1]

		switch outputWire {
		case wireA:
			newConfigurations = append(newConfigurations, fmt.Sprintf("%s -> %s", inputWires, wireB))
		case wireB:
			newConfigurations = append(newConfigurations, fmt.Sprintf("%s -> %s", inputWires, wireA))
		default:
			newConfigurations = append(newConfigurations, config)
		}
	}
	return newConfigurations
}
