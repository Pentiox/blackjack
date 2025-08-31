package main

import "blackjack/config"

// Player is a blackjack player.
type Player struct {
	Chips int
	Hands []*Hand
}

// NewPlayer creates a new blackjack player.
func NewPlayer() *Player {
	return &Player{
		Chips: config.PlayerStartingChips,
		Hands: []*Hand{},
	}
}

// IsBusted checks if the player is busted (out of chips).
func (p *Player) IsBusted() bool {
	return p.Chips <= 0
}

// Players is a list of blackjack players.
type Players []*Player

// NewPlayers creates a new list of blackjack players.
func NewPlayers() Players {

	players := Players{}

	for range config.PlayersPerGame {
		players = append(players, NewPlayer())
	}

	return players
}
