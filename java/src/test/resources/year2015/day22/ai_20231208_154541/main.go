
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define your spell and effect system here

type Player struct {
	hitPoints int
	mana      int
}

type Boss struct {
	hitPoints int
	damage    int
}

type Game struct {
	player Player
	boss   Boss
	// Add mana spent and effects here
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	hitPointsLine := scanner.Text()
	scanner.Scan()
	damageLine := scanner.Text()

	hitPointsStr := strings.Split(hitPointsLine, ": ")[1]
	damageStr := strings.Split(damageLine, ": ")[1]

	hitPoints, _ := strconv.Atoi(hitPointsStr)
	damage, _ := strconv.Atoi(damageStr)

	// Initialize game state
	game := Game{
		player: Player{
			hitPoints: 50, // The player's initial hit points
			mana:      500, // The player's initial mana
		},
		boss: Boss{
			hitPoints: hitPoints,
			damage:    damage,
		},
		// Initialize any additional state needed
	}

	// Implement the simulation of the game and find the least amount of mana

	// Just a placeholder print to avoid syntax error; replace this with the actual result
	fmt.Println("Mana spent to win:", simulateFight(game))
}

// Define a function to simulate the fight and return the least amount of mana used to win

// Placeholder simulation function, you'll need to implement the actual logic
func simulateFight(g Game) int {
	// Your logic to simulate the game and find the minimum mana spent to win goes here
	return 0
}

