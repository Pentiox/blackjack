package main

import "math/rand"

// Card is a blackjack playing card.
type Card struct {
	Rank  Rank
	Suit  Suit
	Value Value
}

// NewCard creates a new blackjack card.
func NewCard(rng *rand.Rand) *Card {

	card := &Card{
		Rank: NewRank(rng),
		Suit: NewSuit(rng),
	}

	card.Value = RankValues[card.Rank]

	return card
}

// Cards is a list of blackjack cards.
type Cards []*Card

// NewCards creates a new list of blackjack cards.
func NewCards(rng *rand.Rand) Cards {
	return make(Cards, 0)
}
