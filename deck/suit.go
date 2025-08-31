package deck

import "math/rand"

// Suit is a blackjack card's suit.
type Suit string

// List of blackjack card suits.
const (
	Clubs    Suit = "Clubs"
	Diamonds Suit = "Diamonds"
	Hearts   Suit = "Hearts"
	Spades   Suit = "Spades"
)

// Suits is a list of blackjack card suits.
var Suits = []Suit{Clubs, Diamonds, Hearts, Spades}

// NewSuit picks a new blackjack card suit.
func NewSuit(rng *rand.Rand) Suit {
	return Suits[rng.Intn(len(Suits))]
}
