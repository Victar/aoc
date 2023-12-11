
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read instructions from the file
	instructions := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	// Execute the program with a starting value of 1 for register a
	registers := map[string]int{"a": 1, "b": 0}
	executeProgram(instructions, registers)

	// Print the value of register b
	fmt.Println(registers["b"])
}

func executeProgram(instructions []string, registers map[string]int) {
	ip := 0 // Instruction pointer
	for ip < len(instructions) {
		fields := strings.Fields(instructions[ip])
		switch fields[0] {
		case "hlf":
			registers[fields[1]] /= 2
		case "tpl":
			registers[fields[1]] *= 3
		case "inc":
			registers[fields[1]]++
		case "jmp":
			offset, _ := strconv.Atoi(fields[1])
			ip += offset
			continue
		case "jie":
			reg, offset := fields[1][:len(fields[1])-1], fields[2]
			if registers[reg]%2 == 0 {
				o, _ := strconv.Atoi(offset)
				ip += o
				continue
			}
		case "jio":
			reg, offset := fields[1][:len(fields[1])-1], fields[2]
			if registers[reg] == 1 {
				o, _ := strconv.Atoi(offset)
				ip += o
				continue
			}
		}
		ip++
	}
}
