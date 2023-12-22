
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pulse int

const (
	Low Pulse = iota
	High
)

type Module interface {
	Receive(p Pulse)
}

type Broadcaster struct {
	destinations []Module
}

type FlipFlop struct {
	name     string
	state    bool
	output   []Module
}

type Conjunct struct {
	name     string
	inputs   map[Module]Pulse
	output   []Module
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{}
}

func (b *Broadcaster) Receive(p Pulse) {
	for _, dest := range b.destinations {
		dest.Receive(p)
	}
}

func NewFlipFlop(name string) *FlipFlop {
	return &FlipFlop{name: name, state: false}
}

func (f *FlipFlop) Receive(p Pulse) {
	if p == Low {
		f.state = !f.state
		nextPulse := High
		if f.state {
			nextPulse = Low
		}
		for _, dest := range f.output {
			dest.Receive(nextPulse)
		}
	}
}

func NewConjunct(name string) *Conjunct {
	return &Conjunct{name: name, inputs: make(map[Module]Pulse)}
}

func (c *Conjunct) Receive(p Pulse) {
	for input := range c.inputs {
		c.inputs[input] = p
	}

	allHigh := true
	for _, pulse := range c.inputs {
		if pulse == Low {
			allHigh = false
			break
		}
	}

	nextPulse := Low
	if !allHigh {
		nextPulse = High
	}

	for _, dest := range c.output {
		dest.Receive(nextPulse)
	}
}

func parseModule(line string, modules map[string]Module) {
	parts := strings.Split(line, " -> ")
	moduleName := strings.TrimSpace(parts[0])
	destNames := strings.Split(parts[1], ", ")
	dests := make([]Module, 0)

	for _, destName := range destNames {
		destName = strings.TrimSpace(destName)
		dests = append(dests, modules[destName])
	}

	if strings.HasPrefix(moduleName, "%") {
		name := strings.TrimPrefix(moduleName, "%")
		modules[name] = NewFlipFlop(name)
		for _, dest := range dests {
			f := modules[name].(*FlipFlop)
			f.output = append(f.output, dest)
		}
	} else if strings.HasPrefix(moduleName, "&") {
		name := strings.TrimPrefix(moduleName, "&")
		modules[name] = NewConjunct(name)
		for _, dest := range dests {
			c := modules[name].(*Conjunct)
			for _, src := range dests {
				c.inputs[src] = Low // Initialize with low pulse memory
			}
			c.output = append(c.output, dest)
		}
	} else if moduleName == "broadcaster" {
		b := NewBroadcaster()
		b.destinations = dests
		modules[moduleName] = b
	}
}

func readInput(filename string) (map[string]Module, Module) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	modules := make(map[string]Module)
	var broadcaster Module

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "broadcaster") {
			// Special treatment for the broadcaster
			parseModule(line, modules)
			broadcaster = modules["broadcaster"]
		} else {
			parseModule(line, modules)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return modules, broadcaster
}

func simulate(modules map[string]Module, broadcaster Module, iterations int) (lowCount, highCount int) {
	for i := 0; i < iterations; i++ { // Push the button 1000 times
		broadcaster.Receive(Low) // Pushing the button sends a low pulse
		lowCount++
	}

	// Since the modules interact in sequence, we can track the actual pulses from conjunct modules
	// by examining how many times the `Receive` method is called with a Low or High pulse.
	for _, module := range modules {
		switch m := module.(type) {
		case *FlipFlop:
			lowCount += m.state // If state is on, add 1
			highCount += 1 - (m.state) // If state is off, add 1
		case *Conjunct:
			for inputState := range m.inputs {
				pulse := m.inputs[inputState]
				if pulse == Low {
					lowCount++
				} else {
					highCount++
				}
			}
		}
	}

	return lowCount, highCount
}

func main() {
	filename := "input.txt"
	modules, broadcaster := readInput(filename)

	// Simulate 1000 button pushes
	lowCount, highCount := simulate(modules, broadcaster, 1000)
	result := lowCount * highCount

	fmt.Println(result)
}
