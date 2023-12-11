
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spell struct {
	Cost, Damage, Heal, Armor, Mana, Duration int
}

var Spells = map[string]Spell{
	"Magic Missile": {Cost: 53, Damage: 4},
	"Drain":         {Cost: 73, Damage: 2, Heal: 2},
	"Shield":        {Cost: 113, Armor: 7, Duration: 6},
	"Poison":        {Cost: 173, Damage: 3, Duration: 6},
	"Recharge":      {Cost: 229, Mana: 101, Duration: 5},
}

type Effect struct {
	Name     string
	Duration int
}

type State struct {
	PlayerHP      int
	PlayerMana    int
	BossHP        int
	ActiveEffects map[string]Effect
	SpentMana     int
	IsPlayerTurn  bool
}

func (s *State) Copy() *State {
	newEffects := make(map[string]Effect)
	for k, v := range s.ActiveEffects {
		newEffects[k] = v
	}
	return &State{
		PlayerHP:      s.PlayerHP,
		PlayerMana:    s.PlayerMana,
		BossHP:        s.BossHP,
		ActiveEffects: newEffects,
		SpentMana:     s.SpentMana,
		IsPlayerTurn:  s.IsPlayerTurn,
	}
}

func (s *State) ApplyEffects() {
	for name, effect := range s.ActiveEffects {
		spell := Spells[name]
		s.BossHP -= spell.Damage
		s.PlayerMana += spell.Mana
		effect.Duration--
		if effect.Duration == 0 {
			delete(s.ActiveEffects, name)
		} else {
			s.ActiveEffects[name] = effect
		}
	}
}

func (s *State) CanCast(spellName string) bool {
	spell := Spells[spellName]
	_, exists := s.ActiveEffects[spellName]
	return !exists && s.PlayerMana >= spell.Cost
}

func (s *State) Cast(spellName string) {
	spell := Spells[spellName]
	s.SpentMana += spell.Cost
	s.PlayerMana -= spell.Cost
	s.BossHP -= spell.Damage
	s.PlayerHP += spell.Heal
	if spell.Duration > 0 {
		s.ActiveEffects[spellName] = Effect{Name: spellName, Duration: spell.Duration}
	}
}

func (s *State) PlayerArmor() int {
	if effect, exists := s.ActiveEffects["Shield"]; exists {
		return Spells["Shield"].Armor
	}
	return 0
}

func (s *State) BossAttack(damage int) {
	actualDamage := damage - s.PlayerArmor()
	if actualDamage < 1 {
		actualDamage = 1
	}
	s.PlayerHP -= actualDamage
}

func (s *State) IsWin() bool {
	return s.BossHP <= 0
}

func (s *State) IsLose() bool {
	return s.PlayerHP <= 0
}

func MinManaWin(state *State, bossDamage int) int {
	if state.IsWin() {
		return state.SpentMana
	}
	if state.IsLose() {
		return 1<<31 - 1 // MaxInt
	}

	state.ApplyEffects()

	if state.IsWin() {
		return state.SpentMana
	}

	if !state.IsPlayerTurn {
		stateCopy := state.Copy()
		stateCopy.BossAttack(bossDamage)
		stateCopy.IsPlayerTurn = true
		return MinManaWin(stateCopy, bossDamage)
	}

	minMana := 1<<31 - 1
	for spellName := range Spells {
		if state.CanCast(spellName) {
			stateCopy := state.Copy()
			stateCopy.Cast(spellName)
			stateCopy.IsPlayerTurn = false
			mana := MinManaWin(stateCopy, bossDamage)
			if mana < minMana {
				minMana = mana
			}
		}
	}

	return minMana
}

func readInput(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bossHP := 0
	bossDamage := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[1])

		if strings.Contains(parts[0], "Hit Points") {
			bossHP = value
		} else if strings.Contains(parts[0], "Damage") {
			bossDamage = value
		}
	}

	return bossHP, bossDamage
}

func main() {
	bossHP, bossDamage := readInput("input.txt")

	initialState := &State{
		PlayerHP:      50,
		PlayerMana:    500,
		BossHP:        bossHP,
		ActiveEffects: make(map[string]Effect),
		IsPlayerTurn:  true,
	}

	fmt.Println(MinManaWin(initialState, bossDamage))
}
