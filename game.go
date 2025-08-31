package main

import (
	"math/rand"
	"time"
)

// Game is the state of a blackjack game.
type Game struct {
	Dealer     *Player
	Players    Players
	Statistics *Statistics

	seed int64
}

// NewGame initializes a new blackjack game.
func NewGame() *Game {
	return &Game{
		Dealer:     NewPlayer(),
		Players:    NewPlayers(),
		Statistics: &Statistics{},
		seed:       time.Now().UnixNano(),
	}
}

// isOver checks if the game is over.
//
// The game is over when all players are busted.
func (g *Game) isOver() bool {

	for _, player := range g.Players {
		if !player.IsBusted() {
			return false
		}
	}

	return true
}

// Play runs the blackjack game loop.
func (g *Game) Play() *Statistics {

	for !g.isOver() {

		g.dealCards()
	}

	return g.Statistics
}

// dealCards deals cards to all players and the dealer.
func (g *Game) dealCards() {

	for _, player := range g.Players {

		g.dealCardsToPlayer(player)
	}
}

func (g *Game) dealCardsToPlayer(player *Player) {

	player.Hands = NewHands(rand.New(rand.NewSource(g.seed)))
}
