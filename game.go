package main

import (
	"math/rand"
	"time"

	"blackjack/config"
	"blackjack/deck"
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

	if g.Statistics.TotalHands.Load() <= config.MaxHands {
		return false
	}

	if !g.Players.IsBusted() {
		return false
	}

	return true
}

// Play runs the blackjack game loop.
func (g *Game) Play() *Statistics {

	for !g.isOver() {

		g.dealCards()

		g.resolveHands()
	}

	return g.Statistics
}

// dealCards deals cards to all players and the dealer.
func (g *Game) dealCards() {

	for _, player := range g.Players {
		g.dealCardsToPlayer(player)
	}

	g.dealCardsToPlayer(g.Dealer)
}

// dealCardsToPlayer deals cards to a single player.
func (g *Game) dealCardsToPlayer(player *Player) {

	// Generate a new random number generator for the player.
	rng := rand.New(rand.NewSource(g.seed))

	// Create new hands for the player.
	player.Hands = deck.NewHands(rng)
}

// resolveHands resolves the hands for all players and the dealer.
func (g *Game) resolveHands() {

	for i := range config.PlayerHands {
		for j := range g.Players {
			g.resolveHand(g.Dealer.Hands[i], g.Players[j].Hands[i])
		}
	}

}

// resolveHand resolves the outcome of a single hand between the dealer and a player.
func (g *Game) resolveHand(handDealer, handPlayer *deck.Hand) {

	if handPlayer.IsBusted() {
		g.Statistics.BustsPlayer.Add(1)
	}

	if handDealer.IsBusted() {
		g.Statistics.BustsDealer.Add(1)
	}

	if handDealer.IsGreater(handPlayer) || handDealer.IsEqual(handPlayer) {
		g.resolveWinDealer(handDealer)
	}

	g.resolveWinPlayer(handPlayer)
}

// resolveWinDealer resolves the win for the dealer.
func (g *Game) resolveWinDealer(hand *deck.Hand) {
	g.Statistics.WinsDealer.Add(1)
}

// resolveWinPlayer resolves the win for a player.
func (g *Game) resolveWinPlayer(hand *deck.Hand) {
	g.Statistics.WinsPlayer.Add(1)
}
