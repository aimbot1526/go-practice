package main

import "fmt"

type Player struct {
	health int
}

func takeDamage(player *Player) {
	fmt.Println("Taking damage")
	explosion := 10
	player.health -= explosion
}

func main() {
	newPlayer := &Player{
		health: 100,
	}
	fmt.Println("Player health right now", newPlayer)
	takeDamage(newPlayer)
	fmt.Println("Player after explosion", newPlayer)
}
