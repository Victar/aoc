
package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	NoPulse = 0
	LowPulse = 1
	HighPulse = 2
)

type Module interface {
	ReceivePulse(pulse int) []string
}

type FlipFlop struct {
	Name string
	State bool
	Destinations []string
}

func NewFlipFlop(name string, destinations []string) *FlipFlop {
	return &FlipFlop{
		Name: name,
		State: false,
		Destinations: destinations,
	}
}

func (ff *FlipFlop) ReceivePulse(pulse int) []string {
	if pulse == LowPulse {
		ff.State = !ff.State
		if ff.State {
			return ff.Destinations
		}
	}
	return []string{}
}

type Conjunction struct {
	Name string
	Inputs map[string]int
	Destinations []string
}

func NewConjunction(name string, destinations []string) *Conjunction {
	return &Conjunction{
		Name: name,
		Inputs: make(map[string]int),
		Destinations: destinations,
	}
}

func (c *Conjunction) ReceivePulse(pulse int) []string {
	c.Inputs[moduleName] = pulse
	allHigh := true
	for _, p := range c.Inputs {
		if p != HighPulse {
			allHigh = false
			break
		}
	}
	if allHigh {
		return c.Destinations
	}
	return []string{}
}

type Broadcaster struct {
	Destinations []string
}

func (b *Broadcaster) ReceivePulse(pulse int) []string {
	return b.Destinations
}

func main() {
	modules := make(map[string]Module)
	readModules(modules)

	buttonPresses := 1000
	lowPulses, highPulses := simulate(buttonPresses, modules)
	result := lowPulses * highPulses
	println(result)
}

func simulate(presses int, modules map[string]Module) (int, int) {
	var lowPulses, highPulses int

	broadcaster, _ := modules["broadcaster"].(*Broadcaster)

	for i := 0; i < presses; i++ {
		queue := broadcaster.ReceivePulse(LowPulse)
		lowPulses++

		for len(queue) > 0 {
			moduleName := queue[0]
			queue = queue[1:]

			destinations := modules[moduleName].ReceivePulse(LowPulse)
			if moduleName[0] == '%' {
				highPulses += len(destinations)
			} else {
				lowPulses += len(destinations)
			}

			queue = append(queue, destinations...)
		}
	}

	return lowPulses, highPulses
}

func readModules(modules map[string]Module) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		nameDest := strings.Split(parts[0], " ")

		destinations := strings.Split(parts[1], ", ")

		switch nameDest[0][0] {
		case '%':
			// Flip-Flop module
			modules[nameDest[1]] = NewFlipFlop(nameDest[1], destinations)
		case '&':
			// Conjunction module
			modules[nameDest[1]] = NewConjunction(nameDest[1], destinations)
		default:
			// Broadcaster module
			modules["broadcaster"] = &Broadcaster{Destinations: destinations}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
