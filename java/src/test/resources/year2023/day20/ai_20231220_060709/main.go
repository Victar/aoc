
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ModuleType int

const (
	FlipFlop ModuleType = iota
	Conjunction
	Broadcast
)

type Module struct {
	Type         ModuleType
	Name         string
	Destinations []string
	State        bool // ON/OFF for flip-flop, HIGH/LOW memory for conjunction
	Memories     map[string]bool
}

func main() {
	// Read input and prepare modules
	modules := make(map[string]*Module)
	broadcasterName, err := readInput("input.txt", modules)
	if err != nil {
		panic(err)
	}

	// Simulate the button press 1000 times
	const presses = 1000
	var lowPulses, highPulses int64
	for i := 0; i < presses; i++ {
		queue := []string{broadcasterName}
		pulses := map[string]bool{broadcasterName: false}
		
		// Process queue
		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]
			curPulse := pulses[next]
			if curPulse {
				highPulses++
			} else {
				lowPulses++
			}
			queue = modules[next].process(pulses, queue)
		}
	}

	// Calculate final result and print
	result := lowPulses * highPulses
	fmt.Println(result)
}

func readInput(filename string, modules map[string]*Module) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var broadcasterName string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.TrimSpace(line), " -> ")
		moduleName := strings.Trim(parts[0], "%&")
		moduleDests := strings.Split(parts[1], ", ")

		module := getOrCreateModule(modules, moduleName[1:])
		for _, destName := range moduleDests {
			module.Destinations = append(module.Destinations, strings.TrimSpace(destName))
			getOrCreateModule(modules, destName) // Ensures all destinations exist as modules.
		}

		switch moduleName[0] {
		case '%':
			module.Type = FlipFlop
		case '&':
			module.Type = Conjunction
			module.Memories = make(map[string]bool)
			for _, dest := range module.Destinations {
				module.Memories[dest] = false
			}
		default:
			module.Type = Broadcast
			broadcasterName = moduleName
		}
	}
	return broadcasterName, scanner.Err()
}

func getOrCreateModule(modules map[string]*Module, name string) *Module {
	module, exists := modules[name]
	if !exists {
		module = &Module{Name: name}
		modules[name] = module
	}
	return module
}

func (m *Module) process(pulses map[string]bool, queue []string) []string {
	switch m.Type {
	case FlipFlop:
		if !pulses[m.Name] {
			m.State = !m.State
			pulseType := m.State
			for _, dest := range m.Destinations {
				pulses[dest] = pulseType
				queue = append(queue, dest)
			}
		}
	case Conjunction:
		allHigh := true
		m.Memories[m.Name] = pulses[m.Name]
		for _, memory := range m.Memories {
			allHigh = allHigh && memory
		}
		pulseType := !allHigh
		for _, dest := range m.Destinations {
			m.Memories[dest] = pulses[m.Name]
			pulses[dest] = pulseType
			queue = append(queue, dest)
		}
	case Broadcast:
		for _, dest := range m.Destinations {
			pulses[dest] = pulses[m.Name]
			queue = append(queue, dest)
		}
	}
	return queue
}
