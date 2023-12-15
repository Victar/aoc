
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Chemical struct {
	name     string
	quantity int64
}

type Reaction struct {
	output Chemical
	inputs []Chemical
}

func parseInput(filename string) (map[string]Reaction, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reactions := make(map[string]Reaction)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " => ")
		inputsPart := parts[0]
		outputPart := parts[1]

		var inputs []Chemical
		for _, in := range strings.Split(inputsPart, ", ") {
			pieces := strings.Split(in, " ")
			quantity, _ := strconv.ParseInt(pieces[0], 10, 64)
			inputs = append(inputs, Chemical{pieces[1], quantity})
		}

		outputPieces := strings.Split(outputPart, " ")
		outputQuantity, _ := strconv.ParseInt(outputPieces[0], 10, 64)
		reactions[outputPieces[1]] = Reaction{
			output: Chemical{outputPieces[1], outputQuantity},
			inputs: inputs,
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reactions, nil
}

// calculateORE calculates the amount of ORE needed to produce the requested quantity of a given chemical.
func calculateORE(reactions map[string]Reaction, chemical string, quantity int64, surplus map[string]int64) int64 {
	if chemical == "ORE" {
		return quantity
	}

	// If there is surplus chemical, use that first.
	if surplus[chemical] > 0 {
		used := min(quantity, surplus[chemical])
		quantity -= used
		surplus[chemical] -= used
	}

	reaction := reactions[chemical]
	times := (quantity + reaction.output.quantity - 1) / reaction.output.quantity

	var ore int64
	for _, input := range reaction.inputs {
		ore += calculateORE(reactions, input.name, input.quantity*times, surplus)
	}

	// Add any surplus produced to the surplus map
	surplus[chemical] += times*reaction.output.quantity - quantity

	return ore
}

// min is a helper function to return the minimum of two int64 numbers.
func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	reactions, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	surplus := make(map[string]int64)
	oreNeeded := calculateORE(reactions, "FUEL", 1, surplus)

	fmt.Println(oreNeeded)
}
