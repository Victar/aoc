
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the types of modules
type Module interface {
	Pulse(bool)
	AddDestination(Module)
}

// FlipFlop module
type FlipFlop struct {
	name         string
	on           bool
	destinations []Module
}

func (f *FlipFlop) Pulse(highPulse bool) {
	if !highPulse {
		f.on = !f.on
		for _, dest := range f.destinations {
			dest.Pulse(f.on)
		}
	}
}

func (f *FlipFlop) AddDestination(m Module) {
	f.destinations = append(f.destinations, m)
}

// Conjunction module
type Conjunction struct {
	name         string
	inputMemory  map[string]bool
	destinations []Module
}

func (c *Conjunction) Pulse(highPulse bool) {
	allHigh := true
	for _, pulse := range c.inputMemory {
		allHigh = allHigh && pulse
	}
	for _, dest := range c.destinations {
		dest.Pulse(!allHigh)
	}
}

func (c *Conjunction) SetInputMemory(moduleName string, highPulse bool) {
	c.inputMemory[moduleName] = highPulse
}

func (c *Conjunction) AddDestination(m Module) {
	c.destinations = append(c.destinations, m)
}

// Broadcaster module
type Broadcaster struct {
	destinations []Module
}

func (b *Broadcaster) Pulse(highPulse bool) {
	for _, dest := range b.destinations {
		dest.Pulse(highPulse)
	}
}

func (b *Broadcaster) AddDestination(m Module) {
	b.destinations = append(b.destinations, m)
}

func newFlipFlop(name string) *FlipFlop {
	return &FlipFlop{
		name:         name,
		on:           false,
		destinations: make([]Module, 0),
	}
}

func newConjunction(name string, inputs []string) *Conjunction {
	inputMemory := make(map[string]bool)
	for _, input := range inputs {
		inputMemory[input] = false // Start with the assumption of low pulse
	}
	return &Conjunction{
		name:         name,
		inputMemory:  inputMemory,
		destinations: make([]Module, 0),
	}
}

func newBroadcaster() *Broadcaster {
	return &Broadcaster{
		destinations: make([]Module, 0),
	}
}

func parseLine(line string, modules map[string]Module) {
	parts := strings.Split(line, " -> ")
	moduleName := strings.TrimSpace(parts[0])
	moduleDestinations := strings.Split(parts[1], ", ")

	var m Module
	switch moduleName[0] {
	case '%':
		m = newFlipFlop(moduleName[1:])
	case '&':
		m = newConjunction(moduleName[1:], moduleDestinations)
	default:
		m = newBroadcaster()
	}

	for _, destName := range moduleDestinations {
		dest := modules[destName]
		if dest == nil {
			// If module doesn't exist yet, create and add it
			if destName[0] == '%' {
				dest = newFlipFlop(destName[1:])
			} else if destName[0] == '&' {
				dest = newConjunction(destName[1:], nil)
			} else {
				panic("Unsupported module type")
			}
			modules[destName] = dest
		}
		m.AddDestination(dest)
	}

	modules[moduleName] = m
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	modules := make(map[string]Module)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line, modules)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	broadcaster := modules["broadcaster"].(*Broadcaster)
	lowPulses := 0
	highPulses := 0

	// Simulate button presses
	for i := 0; i < 1000; i++ {
		lowPulses++
		broadcaster.Pulse(false)
		// Iterate over all modules and update the state as necessary
		for _, m := range modules {
			if f, ok := m.(*FlipFlop); ok {
				if f.on {
					highPulses++
				} else {
					lowPulses++
				}
			} else if c, ok := m.(*Conjunction); ok {
				for inputName, pulse := range c.inputMemory {
					c.SetInputMemory(inputName, pulse)
					if pulse {
						highPulses++
					} else {
						lowPulses++
					}
				}
			}
		}
	}

	// Output the result
	fmt.Println(lowPulses * highPulses)
}
