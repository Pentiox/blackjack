package main

import (
	"blackjack/config"
	"math/rand"
)

// Hand is a blackjack player's hand.
type Hand struct {
	Cards []*Card
}

// NewHand creates a new blackjack hand.
func NewHand(rng *rand.Rand) *Hand {
	return &Hand{
		Cards: NewCards(rng),
	}
}

// Hands is a list of blackjack hands.
type Hands []*Hand

// NewHands creates a new list of blackjack hands.
func NewHands(rng *rand.Rand) Hands {

	hands := Hands{}

	for range config.PlayerHands {
		hands = append(hands, NewHand(rng))
	}

	return hands
}
