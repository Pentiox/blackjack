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

	if g.Statistics.TotalHands <= config.MaxHands {
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

			g.resolveHand(g.Dealer.Hands[i], g.Players[j].Hands[i], j)

			if g.Players[j].IsBusted() {
				g.Players = append(g.Players[:j], g.Players[j+1:]...)
			}
		}
	}
}

// resolveHand resolves the outcome of a single hand between the dealer and a player.
func (g *Game) resolveHand(handDealer, handPlayer *deck.Hand, playerIndex int) {

	g.Statistics.TotalHands++

	if handPlayer.IsBusted() {
		g.Statistics.BustsPlayer++
	}

	if handDealer.IsBusted() {
		g.Statistics.BustsDealer++
	}

	if handDealer.IsGreater(handPlayer) {
		g.resolveWinDealer(handDealer, playerIndex)
	} else if handDealer.IsEqual(handPlayer) {
		g.resolvePush(handDealer, playerIndex)
	} else {
		g.resolveWinPlayer(handPlayer, playerIndex)
	}

	g.resolveJokersDealer(handDealer, playerIndex)
	g.resolveJokersPlayer(handPlayer, playerIndex)
}

// resolveJokersDealer resolves the jokers for the dealer.
func (g *Game) resolveJokersDealer(hand *deck.Hand, playerIndex int) {

	for _, card := range hand.Cards {

		if card.Rank != deck.RankJoker {
			continue
		}

		// Track the number of jokers for the dealer.
		if playerIndex == 0 {
			g.Statistics.JokersDealer++
		}

		g.Players[playerIndex].Chips--
	}
}

// resolveJokersPlayer resolves the jokers for a player.
func (g *Game) resolveJokersPlayer(hand *deck.Hand, playerIndex int) {

	for _, card := range hand.Cards {

		if card.Rank != deck.RankJoker {
			continue
		}

		g.Statistics.JokersPlayer++
		g.Players[playerIndex].Chips--
	}
}

// resolvePush resolves the push for a player.
//
// This assumes the dealer wins the hand.
func (g *Game) resolvePush(hand *deck.Hand, playerIndex int) {
	g.Statistics.TotalPushes++
	g.resolveWinDealer(hand, playerIndex)
}

// resolveWinDealer resolves the win for the dealer.
func (g *Game) resolveWinDealer(hand *deck.Hand, playerIndex int) {

	g.Players[playerIndex].Chips--
	g.Statistics.WinsDealer++

	if hand.IsBlackjack() {
		g.Players[playerIndex].Chips--
		g.Statistics.BlackjacksDealer++
	}

	if hand.IsEights() {
		g.Players[playerIndex].Chips--
		g.Statistics.EightsDealer++
	}

	if hand.IsKingAndQueen() {
		g.Players[playerIndex].Chips--
		g.Statistics.KingAndQueensDealer++
	}
}

// resolveWinPlayer resolves the win for a player.
func (g *Game) resolveWinPlayer(hand *deck.Hand, playerIndex int) {

	g.Players[playerIndex].Chips++
	g.Statistics.WinsPlayer++

	if hand.IsBlackjack() {
		g.Players[playerIndex].Chips++
		g.Statistics.BlackjacksPlayer++
	}

	if hand.IsEights() {
		g.Players[playerIndex].Chips++
		g.Statistics.EightsPlayer++
	}

	if hand.IsKingAndQueen() {
		g.Players[playerIndex].Chips++
		g.Statistics.KingAndQueensPlayer++
	}
}
