
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Module interface {
	SendPulse(pulse int)
	ReceivedFrom(Module) int
	Process()
}

type FlipFlop struct {
	name       string
	state      int
	outputs    []Module
	outputPulse int
}

type Conjunction struct {
	name       string
	inputsState map[Module]int
	outputs     []Module
}

type Broadcaster struct {
	name    string
	outputs []Module
}

func (ff *FlipFlop) SendPulse(pulse int) {
	if pulse == 0 {
		ff.state = 1 - ff.state
		ff.outputPulse = ff.state
	}
}

func (ff *FlipFlop) ReceivedFrom(Module) int {
	return 1
}

func (ff *FlipFlop) Process() {
	for _, module := range ff.outputs {
		module.SendPulse(ff.outputPulse)
	}
}

func (c *Conjunction) SendPulse(pulse int) {
	for input := range c.inputsState {
		if c.inputsState[input] == 0 {
			c.outputs[0].SendPulse(1)
			return
		}
	}
	c.outputs[0].SendPulse(0)
}

func (c *Conjunction) ReceivedFrom(mod Module) int {
	return c.inputsState[mod]
}

func (c *Conjunction) Process() { /* Conjunction has no processing to do */ }

func (b *Broadcaster) SendPulse(pulse int) {
	for _, module := range b.outputs {
		module.SendPulse(pulse)
	}
}

func (b *Broadcaster) ReceivedFrom(Module) int {
	return 1
}

func (b *Broadcaster) Process() { /* Broadcaster has no processing to do */ }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	modules := make(map[string]Module)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		moduleName := strings.TrimLeft(parts[0], "%&")
		var mod Module
		if strings.HasPrefix(parts[0], "%") {
			mod = &FlipFlop{name: moduleName, state: 0, outputPulse: 1}
		} else if strings.HasPrefix(parts[0], "&") {
			mod = &Conjunction{name: moduleName, inputsState: make(map[Module]int)}
		} else {
			mod = &Broadcaster{name: moduleName}
		}
		modules[moduleName] = mod
		for _, dest := range strings.Split(parts[2], ",") {
			dest = strings.TrimSpace(dest)
			if strings.HasPrefix(dest, "%") || strings.HasPrefix(dest, "&") {
				dest = strings.TrimLeft(dest, "%&")
			}
			if destModule, ok := modules[dest]; ok {
				switch mod := mod.(type) {
				case *FlipFlop:
					mod.outputs = append(mod.outputs, destModule)
				case *Conjunction:
					mod.inputsState[destModule] = 0
					if len(mod.outputs) == 0 {
						mod.outputs = append(mod.outputs, destModule)
					}
				case *Broadcaster:
					mod.outputs = append(mod.outputs, destModule)
				}
			} else {
				fmt.Printf("Destination module %s not found for module %s\n", dest, moduleName)
				return
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	broadcaster, ok := modules["broadcaster"].(*Broadcaster)
	if !ok {
		fmt.Println("Broadcaster module not found.")
		return
	}

	const buttonPresses = 1000
	lowPulses, highPulses := 0, 0

	for i := 0; i < buttonPresses; i++ {
		queue := []Module{broadcaster}
		visited := make(map[Module]bool)

		broadcaster.SendPulse(0)
		lowPulses++ // Initial button press

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			if _, processed := visited[current]; !processed {
				current.Process()
				visited[current] = true

                // SendPulse might have caused state changes, requeue the module
				switch module := current.(type) {
				case *FlipFlop:
					if module.state == 1 {
						highPulses++
					} else {
						lowPulses++
					}
					for _, dest := range module.outputs {
						if !visited[dest] {
							queue = append(queue, dest)
						}
					}
				case *Conjunction:
					module.SendPulse(module.ReceivedFrom(current))
				}
			}
		}

		// Update conjunction states
		for _, mod := range modules {
			conj, isConj := mod.(*Conjunction)
			if isConj {
				for input := range conj.inputsState {
					conj.inputsState[input] = input.ReceivedFrom(conj)
				}
			}
		}
	}

	fmt.Println(lowPulses * highPulses)
}
