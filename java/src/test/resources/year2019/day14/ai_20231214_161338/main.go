
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reactions := readInput("input.txt")
	oreNeeded := calculateOreForFuel(reactions, 1)
	fmt.Println(oreNeeded)
}

func readInput(filename string) map[string]reaction {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reactions := make(map[string]reaction)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " => ")
		outputParts := strings.Split(parts[1], " ")
		amount, _ := strconv.Atoi(outputParts[0])
		name := outputParts[1]

		inputs := make(map[string]int)
		for _, input := range strings.Split(parts[0], ", ") {
			inputParts := strings.Split(input, " ")
			amount, _ := strconv.Atoi(inputParts[0])
			name := inputParts[1]
			inputs[name] = amount
		}

		reactions[name] = reaction{amount, inputs}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return reactions
}

func calculateOreForFuel(reactions map[string]reaction, fuelAmount int) int {
	needs := map[string]int{"FUEL": fuelAmount}
	extras := make(map[string]int)
	var oreNeeded int

	for len(needs) > 0 {
		for chem, qty := range needs {
			if chem == "ORE" {
				oreNeeded += qty
				delete(needs, chem)
				continue
			}

			if extra, exists := extras[chem]; exists && extra > 0 {
				if extra >= qty {
					extras[chem] -= qty
					delete(needs, chem)
					continue
				} else {
					qty -= extra
					delete(extras, chem)
				}
			}

			reaction := reactions[chem]
			times := (qty + reaction.output - 1) / reaction.output

			for input, inputQty := range reaction.inputs {
				needs[input] += inputQty * times
			}

			produced := reaction.output * times
			extras[chem] += produced - qty
			delete(needs, chem)
		}
	}

	return oreNeeded
}

type reaction struct {
	output int
	inputs map[string]int
}

