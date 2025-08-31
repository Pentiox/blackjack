package deck

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
