package deck

import (
	"fmt"
	"math/rand"
)

// Card is a blackjack playing card.
type Card struct {
	Rank  Rank
	Suit  Suit
	Value Value
}

// String returns a string representation of the card.
func (c *Card) String() string {
	return fmt.Sprintf("%d of %s", c.Rank, c.Suit)
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
