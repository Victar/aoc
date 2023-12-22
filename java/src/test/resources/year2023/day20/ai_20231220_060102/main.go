
package main

import (
	"bufio"
	"os"
	"strings"
)

type Module interface {
	ReceivePulse(pulse bool) []string
}

type FlipFlop struct {
	Name     string
	State    bool
	DestMods []string
}

func (ff *FlipFlop) ReceivePulse(pulse bool) (next []string) {
	if !pulse {
		ff.State = !ff.State
		for _, dest := range ff.DestMods {
			next = append(next, dest)
		}
	}
	return next
}

type Conjunction struct {
	Name      string
	Inputs    map[string]bool
	DestMods  []string
	AllHigh   bool
	FirstPulse bool
}

func (c *Conjunction) ReceivePulse(pulse bool) (next []string) {
	if c.FirstPulse {
		c.AllHigh = pulse
		c.FirstPulse = false
	} else {
		c.AllHigh = c.AllHigh && pulse
	}

	sendPulse := !c.AllHigh
	for _, dest := range c.DestMods {
		next = append(next, dest)
	}

	c.AllHigh = !sendPulse // Reset the state.
	return next
}

type Broadcaster struct {
	DestMods []string
}

func (b *Broadcaster) ReceivePulse(pulse bool) []string {
	return b.DestMods
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
		parts := strings.Split(line, " -> ")
		namePart := strings.TrimSpace(parts[0])
		destPart := strings.Split(parts[1], ",")
		for i := range destPart {
			destPart[i] = strings.TrimSpace(destPart[i])
		}

		if namePart[0] == '%' {
			modules[namePart[1:]] = &FlipFlop{Name: namePart[1:], DestMods: destPart}
		} else if namePart[0] == '&' {
			modules[namePart[1:]] = &Conjunction{
				Name:      namePart[1:],
				Inputs:    make(map[string]bool),
				DestMods:  destPart,
				AllHigh:   false,
				FirstPulse: true,
			}
		} else if namePart == "broadcaster" {
			modules["broadcaster"] = &Broadcaster{DestMods: destPart}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	countLow := 1000 + 1 // Button is also considered as sending low pulses
	countHigh := 0
	// Push the button 1000 times.
	for i := 0; i < 1000; i++ {
		// Define a queue to process pulses in order.
		var queue []string = modules["broadcaster"].ReceivePulse(false)
		for len(queue) > 0 {
			mod := queue[0]
			queue = queue[1:]

			pulse := false
			if f, ok := modules[mod].(*FlipFlop); ok {
				pulse = !f.State
			}

			if pulse {
				countHigh++
			} else {
				countLow++
			}

			queue = append(queue, modules[mod].ReceivePulse(pulse)...)
		}
	}

	// Output result
	println(countLow * countHigh)
}
