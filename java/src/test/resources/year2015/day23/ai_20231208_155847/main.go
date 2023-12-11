
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read instructions from the file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var instructions []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Initialize the registers
	registers := map[string]int{"a": 0, "b": 0}

	// Execute the instructions
	for i := 0; i < len(instructions); {
		instr := strings.Fields(instructions[i])
		switch instr[0] {
		case "hlf":
			registers[instr[1]] /= 2
			i++
		case "tpl":
			registers[instr[1]] *= 3
			i++
		case "inc":
			registers[instr[1]]++
			i++
		case "jmp":
			offset, _ := strconv.Atoi(instr[1])
			i += offset
		case "jie":
			offset, _ := strconv.Atoi(instr[2])
			if registers[instr[1][0:1]]%2 == 0 {
				i += offset
			} else {
				i++
			}
		case "jio":
			offset, _ := strconv.Atoi(instr[2])
			if registers[instr[1][0:1]] == 1 {
				i += offset
			} else {
				i++
			}
		}
	}

	// Print the result
	fmt.Println(registers["b"])
}
