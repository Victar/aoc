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
	runBoth()
}

func runBoth() {
	lines, err := util.ReadFile("year2024/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	registers := make([]int64, 3)
	program := []int64{}
	registers[0], _ = strconv.ParseInt(strings.TrimSpace(lines[0][12:]), 10, 64)
	registers[1], _ = strconv.ParseInt(strings.TrimSpace(lines[1][12:]), 10, 64)
	registers[2], _ = strconv.ParseInt(strings.TrimSpace(lines[2][12:]), 10, 64)
	instructions := strings.Split(lines[4][9:], ",")
	for _, instr := range instructions {
		num, _ := strconv.ParseInt(strings.TrimSpace(instr), 10, 64)
		program = append(program, num)
	}
	registerCopy := make([]int64, len(registers))
	_ = copy(registerCopy, registers)
	result := executeProgram(registerCopy, program)
	fmt.Println(result)

	step := int64(0)
	found := false
	goldAnswer := int64(0)
	//Brute force with manual iteration adjustment
	accuracy := 16 // 247839653009594
	accuracyShow := 11
	start := int64(247839653000000)
	stepAdd := int64(1)

	iterLimit := 100000000
	iter := 0
	for !found {
		iter++
		startUp := start + step
		startDown := start - step
		if startDown < 0 || iter == iterLimit {
			fmt.Println("iter", iter)
			break
		}
		programUp, indexUp := executeWithIndex(startUp, registers, program)
		programDown, indexDown := executeWithIndex(startDown, registers, program)
		if indexUp >= accuracyShow {
			fmt.Println("up  ", programUp, program, indexUp, startUp)
		}
		if indexDown >= accuracyShow {
			fmt.Println("down", programDown, program, indexDown, startDown)
		}
		if indexDown >= accuracy || indexUp >= accuracy {
			found = true
		}
		step = step + stepAdd
	}
	fmt.Println(goldAnswer)
}

func findEqualIndex(currentProgram []int64, program []int64) int {
	if len(currentProgram) != len(program) {
		return -1
	}
	result := 0
	for index, _ := range currentProgram {
		if currentProgram[index] == program[index] {
			result++
		}
	}
	return result
}
func executeWithIndex(numberToCheck int64, registers []int64, program []int64) ([]int64, int) {
	registerCopy := make([]int64, len(registers))
	_ = copy(registerCopy, registers)
	registerCopy[0] = numberToCheck
	currentProgram := executeProgram(registerCopy, program)
	eqIndex := findEqualIndex(currentProgram, program)
	return currentProgram, eqIndex
}

func executeProgram(registers []int64, program []int64) []int64 {
	ip := 0
	output := []int64{}
	for {
		if ip >= len(program) {
			break
		}
		opcode := program[ip]
		operand := program[ip+1]
		ip += 2
		switch opcode {
		case 0: // adv
			denom := int64(1 << comboOperandValue(operand, registers))
			registers[0] /= denom
		case 1: // bxl
			registers[1] ^= operand
		case 2: // bst
			registers[1] = comboOperandValue(operand, registers) % 8
		case 3: // jnz
			if registers[0] != 0 {
				ip = int(operand)
			}
		case 4: // bxc
			registers[1] ^= registers[2]
		case 5: // out
			output = append(output, comboOperandValue(operand, registers)%8)
		case 6: // bdv
			denom := int64(1 << comboOperandValue(operand, registers))
			registers[1] = registers[0] / denom
		case 7: // cdv
			denom := int64(1 << comboOperandValue(operand, registers))
			registers[2] = registers[0] / denom
		default:
			log.Fatalf("Invalid opcode: %d", opcode)
		}
	}
	return output
}

func comboOperandValue(op int64, registers []int64) int64 {
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
