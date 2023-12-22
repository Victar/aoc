package main

import (
	"adventofcode/util"
	"fmt"
	"strings"
)

var DAY = "20"

type Module interface {
	ReceivePulse(pulse bool, prevName string) (signal Signal)
	GetModules() []string
}
type Signal struct {
	next       []string
	pulseNext  bool
	nextParent string
}

func (s *Signal) countSignal() {
	if s.pulseNext {
		COUNT_TRUE += len(s.next)
	} else {
		COUNT_FALSE += len(s.next)
	}
}

func (s *Signal) printSignal() {
	pulse := "low"
	if s.pulseNext {
		pulse = "high"
	}
	for _, n := range s.next {
		fmt.Println(s.nextParent, pulse, n)
	}
}

type FlipFlop struct {
	name    string
	state   bool // on or off
	modules []string
}

func (ff *FlipFlop) ReceivePulse(pulse bool, prevName string) (signal Signal) {
	next := []string{}
	if pulse { // high
		return Signal{[]string{}, pulse, ff.name}
	} else { // low
		ff.state = !ff.state
		for _, module := range ff.modules {
			next = append(next, module)
		}
	}
	return Signal{next, ff.state, ff.name}
}

func (ff *FlipFlop) GetModules() []string {
	return ff.modules
}

type Conjunction struct {
	name          string
	modulePulse   map[string]bool
	modules       []string
	prevHigh      int
	prevHighCycle []int
}

func (c *Conjunction) ReceivePulse(pulse bool, prevName string) (signal Signal) {
	c.modulePulse[prevName] = pulse
	allHigh := true
	for _, modPulse := range c.modulePulse {
		if modPulse {
			continue
		} else {
			allHigh = false
		}
	}
	if !allHigh {
		cycle := GLOB_I - c.prevHigh
		c.prevHigh = GLOB_I
		c.prevHighCycle = append(c.prevHighCycle, cycle)
		//fmt.Println(c, GLOB_I)
	}
	return Signal{c.modules, !allHigh, c.name}
}

func (c *Conjunction) GetModules() []string {
	return c.modules
}

type Broadcaster struct {
	name    string
	modules []string
}

func (b *Broadcaster) ReceivePulse(pulse bool, prevName string) (signal Signal) {
	return Signal{b.modules, pulse, b.name}
}

func (b *Broadcaster) GetModules() []string {
	return b.modules
}

var COUNT_TRUE = 0
var COUNT_FALSE = 0
var GLOB_I = 0

func main() {
	runBoth()
}

func runBoth() {
	lines, err := util.ReadFile("year2023/day" + DAY + "/input.txt")
	if err != nil {
		panic(err)
	}
	modules := make(map[string]Module)
	conjs := []*Conjunction{}
	conjsMap := make(map[string]*Conjunction)

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		if parts[0][:1] == "%" {
			modules[parts[0][1:]] = &FlipFlop{name: parts[0][1:], state: false, modules: strings.Split(parts[1], ", ")}
		}
		if parts[0][:1] == "&" {
			modulePulse := make(map[string]bool)
			con := &Conjunction{name: parts[0][1:], modulePulse: modulePulse, modules: strings.Split(parts[1], ", ")}
			modules[parts[0][1:]] = con
			conjs = append(conjs, con)
			conjsMap[parts[0][1:]] = con
		}
		if parts[0] == "broadcaster" {
			modules[parts[0]] = &Broadcaster{name: parts[0], modules: strings.Split(parts[1], ", ")}
		}
	}

	//add destinations for conjunction
	for _, con := range conjs {
		for k, v := range modules {
			for _, vmod := range v.GetModules() {
				if vmod == con.name {
					con.modulePulse[k] = false
				}
			}
		}
	}

	count := 20000

	for i := 0; i < count; i++ {
		if i == 1000 {
			fmt.Println(COUNT_FALSE * COUNT_TRUE)
		}
		GLOB_I = i
		COUNT_FALSE++
		signal := modules["broadcaster"].ReceivePulse(false, "")
		nextQueue := []Signal{signal}
		for len(nextQueue) > 0 {
			signal := nextQueue[0]
			signal.countSignal()
			nextQueue = nextQueue[1:]
			if len(signal.next) > 0 {
				for _, n := range signal.next {
					mod, has := modules[n]
					if has {
						newSignal := mod.ReceivePulse(signal.pulseNext, signal.nextParent)
						nextQueue = append(nextQueue, newSignal)
					}
				}
			}
		}

	}
	conToSearch := []string{"rv", "vp", "cq", "dc"}
	cycles := []int{}
	for _, c := range conToSearch {
		cycles = append(cycles, conjsMap[c].prevHighCycle[len(conjsMap[c].prevHighCycle)-1])
	}
	fmt.Println(findLCM(cycles))
}

// Function to find the greatest common divisor (GCD) using Euclidean algorithm
func findGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to find the least common multiple (LCM) of a slice of integers
func findLCM(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		gcd := findGCD(lcm, nums[i])
		lcm = (lcm * nums[i]) / gcd
	}
	return lcm
}
