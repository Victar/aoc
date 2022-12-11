package main

import (
	"adventofcode/util"
	"sort"
	"strconv"
	"strings"
)

func main() {
	runSilver()
	runGold()
}
func runSilver() {
	runAny(false, 20)
}

func runGold() {
	runAny(true, 10000)
}

type Monkey struct {
	id                int
	currentItems      []int
	operation         string
	testNumber        int
	testTrueMonkeyId  int
	testFalseMonkeyId int
	inspectionCount   int
}

func newMonkey(sName, sStart, sOperation, sTest, sTestTrue, sTestFalse string) *Monkey {
	id, _ := strconv.Atoi(sName[7:8])
	operation := sOperation[18:]
	testNumber, _ := strconv.Atoi(sTest[21:])
	testTrueMonleyId, _ := strconv.Atoi(sTestTrue[29:])
	testFalseMonkeyId, _ := strconv.Atoi(sTestFalse[30:])
	startSlice := strings.Split(sStart[18:], ", ")
	currentItems := []int{}
	for _, s := range startSlice {
		sInt, _ := strconv.Atoi(s)
		currentItems = append(currentItems, sInt)
	}
	return &Monkey{
		id:                id,
		operation:         operation,
		testNumber:        testNumber,
		testFalseMonkeyId: testFalseMonkeyId,
		testTrueMonkeyId:  testTrueMonleyId,
		currentItems:      currentItems,
	}
}

func (monkey *Monkey) appendWorry(currWorry int) {
	items := append(monkey.currentItems, currWorry)
	monkey.currentItems = items
}

func (monkey *Monkey) processRound(monkeys map[int]*Monkey, gold bool, commonDivider int) {
	for i := 0; i < len(monkey.currentItems); i++ {
		currWorry := monkey.getNewWorry(monkey.currentItems[i])
		if gold {
			currWorry = currWorry % commonDivider
		} else {
			currWorry = currWorry / 3
		}
		if currWorry%monkey.testNumber == 0 {
			monkeys[monkey.testTrueMonkeyId].appendWorry(currWorry)
		} else {
			monkeys[monkey.testFalseMonkeyId].appendWorry(currWorry)
		}
		monkey.inspectionCount++
	}
	monkey.currentItems = []int{}
}

func (monkey *Monkey) getNewWorry(current int) int {
	postFix := monkey.operation[7:]
	var number int
	if postFix == "old" {
		number = current
	} else {
		number, _ = strconv.Atoi(postFix)
	}
	if strings.Contains(monkey.operation, "*") {
		return current * number
	}
	return current + number
}

func runAny(gold bool, cycleCount int) {
	lines, err := util.ReadFile("year2022/day11/input.txt")
	if err != nil {
		panic(err)
	}
	monkeys := make(map[int]*Monkey)
	commonDivider := 1
	for i := 0; i < len(lines); i = i + 7 {
		curMonkey := newMonkey(lines[i], lines[i+1], lines[i+2], lines[i+3], lines[i+4], lines[i+5])
		monkeys[curMonkey.id] = curMonkey
		commonDivider *= curMonkey.testNumber
	}
	for j := 0; j < cycleCount; j++ {
		for i := 0; i < len(monkeys); i++ {
			monkeys[i].processRound(monkeys, gold, commonDivider)
		}
	}

	inspections := make([]int, len(monkeys))
	for i, m := range monkeys {
		inspections[i] = m.inspectionCount
	}
	sort.Ints(inspections)
	println(inspections[len(inspections)-1] * inspections[len(inspections)-2])
}
