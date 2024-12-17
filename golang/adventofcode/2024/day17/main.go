package main

import (
	"adventofcode/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var DAY = "17"

func main() {
	runSilver()
	//runGold()
}

func runSilver() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {

		println(line)
	}
	registers := make([]int, 3)
	program := []int{}
	registers[0], _ = strconv.Atoi(strings.TrimSpace(lines[0][12:]))
	registers[1], _ = strconv.Atoi(strings.TrimSpace(lines[1][12:]))
	registers[2], _ = strconv.Atoi(strings.TrimSpace(lines[2][12:]))
	instructions := strings.Split(lines[4][8:], ",")
	for _, instr := range instructions {
		num, _ := strconv.Atoi(strings.TrimSpace(instr))
		program = append(program, num)
	}
	fmt.Println(registers, program)
	result := executeProgram(registers, program)
	fmt.Println(strings.Trim(result, ","))
}

func executeProgram(registers []int, program []int) string {
	ip := 0
	output := ""
	for {
		if ip >= len(program) {
			break
		}
		opcode := program[ip]
		operand := program[ip+1]
		ip += 2
		switch opcode {
		case 0: // adv
			denom := 1 << comboOperandValue(operand, registers)
			registers[0] /= denom
		case 1: // bxl
			registers[1] ^= operand
		case 2: // bst
			registers[1] = comboOperandValue(operand, registers) % 8
		case 3: // jnz
			if registers[0] != 0 {
				ip = operand
			}
		case 4: // bxc
			registers[1] ^= registers[2]
		case 5: // out
			output += fmt.Sprintf("%d,", comboOperandValue(operand, registers)%8)
		case 6: // bdv
			denom := 1 << comboOperandValue(operand, registers)
			registers[1] = registers[0] / denom
		case 7: // cdv
			denom := 1 << comboOperandValue(operand, registers)
			registers[2] = registers[0] / denom
		default:
			log.Fatalf("Invalid opcode: %d", opcode)
		}
	}
	return output
}

func comboOperandValue(op int, registers []int) int {
	switch op {
	case 0, 1, 2, 3:
		return op
	case 4:
		return registers[0]
	case 5:
		return registers[1]
	case 6:
		return registers[2]
	default:
		log.Fatalf("Invalid op: %d", op)
	}
	return 0
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
