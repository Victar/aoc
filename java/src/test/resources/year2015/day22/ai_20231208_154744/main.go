
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

const (
	MagicMissileCost = 53
	DrainCost        = 73
	ShieldCost       = 113
	PoisonCost       = 173
	RechargeCost     = 229
)

type Spell struct {
	name   string
	cost   int
	damage int
	heal   int
	effect *Effect
}

type Effect struct {
	name   string
	timer  int
	armor  int
	damage int
	mana   int
}

type Stats struct {
	hitPoints int
	damage    int
	mana      int
	armor     int
}

type Game struct {
	player      Stats
	boss        Stats
	spentMana   int
	activeSpells map[string]*Effect
}

var spells = []Spell{
	{"Magic Missile", MagicMissileCost, 4, 0, nil},
	{"Drain", DrainCost, 2, 2, nil},
	{"Shield", ShieldCost, 0, 0, &Effect{"Shield", 6, 7, 0, 0}},
	{"Poison", PoisonCost, 0, 0, &Effect{"Poison", 6, 0, 3, 0}},
	{"Recharge", RechargeCost, 0, 0, &Effect{"Recharge", 5, 0, 0, 101}},
}

func minManaToWin(player, boss Stats) int {
	initialGame := Game{
		player:      player,
		boss:        boss,
		activeSpells: make(map[string]*Effect),
	}
	return fight(initialGame, math.MaxInt32, true)
}

func fight(game Game, minManaSpent int, playerTurn bool) int {
	if game.spentMana >= minManaSpent {
		return minManaSpent
	}
	for effectName, effect := range game.activeSpells {
		game.boss.hitPoints -= effect.damage
		game.player.mana += effect.mana
		game.activeSpells[effectName].timer--
		if game.activeSpells[effectName].timer == 0 {
			if effect.name == "Shield" {
				game.player.armor -= effect.armor
			}
			delete(game.activeSpells, effectName)
		}
	}
	if game.boss.hitPoints <= 0 {
		return game.spentMana
	}

	if playerTurn {
		minMana := minManaSpent
		for _, spell := range spells {
			if game.player.mana < spell.cost || (spell.effect != nil && game.activeSpells[spell.name] != nil) {
				continue
			}
			nextGame := game
			nextGame.player.mana -= spell.cost
			nextGame.spentMana += spell.cost
			if spell.effect == nil {
				nextGame.boss.hitPoints -= spell.damage
				nextGame.player.hitPoints += spell.heal
			} else {
				nextGame.activeSpells[spell.name] = &Effect{*spell.effect}
			}
			minMana = fight(nextGame, minMana, false)
		}
		return minMana
	} else {
		damage := game.boss.damage - game.player.armor
		if damage < 1 {
			damage = 1
		}
		game.player.hitPoints -= damage
		if game.player.hitPoints <= 0 {
			return minManaSpent
		}
		return fight(game, minManaSpent, true)
	}
}

func readInput(filename string) (Stats, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Stats{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var boss Stats
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			return Stats{}, err
		}
		switch parts[0] {
		case "Hit Points":
			boss.hitPoints = value
		case "Damage":
			boss.damage = value
		}
	}
	return boss, scanner.Err()
}

func main() {
	boss, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	player := Stats{
		hitPoints: 50,
		mana:      500,
	}

	minMana := minManaToWin(player, boss)
	fmt.Println(minMana)
}
