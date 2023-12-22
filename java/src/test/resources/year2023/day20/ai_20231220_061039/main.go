
// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pulse bool

const (
	High Pulse = true
	Low  Pulse = false
)

type Module interface {
	ReceivePulse(Pulse, map[string]Module)
}

type Broadcaster struct {
	destinations []string
}

func (b *Broadcaster) ReceivePulse(p Pulse, modules map[string]Module) {
	for _, dest := range b.destinations {
		modules[dest].ReceivePulse(p, modules)
	}
}

type FlipFlop struct {
	name         string
	isOn         bool
	destinations []string
}

func (f *FlipFlop) ReceivePulse(p Pulse, modules map[string]Module) {
	if p == Low {
		f.isOn = !f.isOn
		nextPulse := High
		if !f.isOn {
			nextPulse = Low
		}
		for _, dest := range f.destinations {
			modules[dest].ReceivePulse(nextPulse, modules)
		}
	}
}

type Conjunction struct {
	name               string
	inputs             []string
	lastReceivedPulses map[string]Pulse
	destinations       []string
}

func (c *Conjunction) ReceivePulse(p Pulse, modules map[string]Module) {
	allHigh := true
	for _, input := range c.inputs {
		if c.lastReceivedPulses[input] == Low {
			allHigh = false
			break
		}
	}

	outPulse := High
	if allHigh {
		outPulse = Low
	}

	for _, dest := range c.destinations {
		modules[dest].ReceivePulse(outPulse, modules)
	}

	// Update the last received pulse memory
	for _, input := range c.inputs {
		c.lastReceivedPulses[input] = Low // Reset all to low
	}
	c.lastReceivedPulses[p] = High
}

func newConjunction(name string, destinations, inputs []string) Module {
	conj := &Conjunction{
		name:               name,
		destinations:       destinations,
		inputs:             inputs,
		lastReceivedPulses: make(map[string]Pulse),
	}
	for _, input := range inputs {
		conj.lastReceivedPulses[input] = Low
	}
	return conj
}

func parseModule(line string, modules map[string]Module) {
	parts := strings.Split(line, " -> ")
	name := strings.TrimSpace(parts[0])
	destinations := strings.Split(parts[1], ", ")
	for i, dest := range destinations {
		destinations[i] = strings.TrimSpace(dest)
	}
	var mod Module
	if strings.HasPrefix(name, "%") {
		name = strings.TrimPrefix(name, "%")
		mod = &FlipFlop{name: name, destinations: destinations}
	} else if strings.HasPrefix(name, "&") {
		name = strings.TrimPrefix(name, "&")
		mod = newConjunction(name, destinations, strings.Fields(name))
	} else if name == "broadcaster" {
		mod = &Broadcaster{destinations: destinations}
	}

	if mod != nil {
		modules[name] = mod
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	modules := make(map[string]Module)
	for scanner.Scan() {
		parseModule(scanner.Text(), modules)
	}

	const buttonPresses = 1000
	lowPulses := 0
	highPulses := 0

	for i := 0; i < buttonPresses; i++ {
		// Each press sends a low pulse to the broadcaster, so count that
		lowPulses++

		// Simulate sending the pulse through the network
		broadcaster := modules["broadcaster"]
		broadcaster.ReceivePulse(Low, modules)

		// TODO: The implementation is incomplete and likely incorrect.
		// Implement proper pulse propagation and handling logic following task rules.
	}

	// Optional: Report individual counts of low and high pulses.
	fmt.Println("Low Pulses:", lowPulses, "High Pulses:", highPulses)

	// Print the result of multiplying the number of low pulses sent by the number of high pulses sent.
	result := lowPulses * highPulses
	fmt.Println(result)
}
