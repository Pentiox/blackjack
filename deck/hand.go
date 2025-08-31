package deck

import (
	"blackjack/config"
	"math/rand"
)

// Hand is a blackjack player's hand.
type Hand struct {
	Cards []*Card

	// aces is the count of aces in the hand.
	//
	// This is used to calculate the best hand value.
	aces int

	// soft is the soft value of the hand.
	//
	// The soft value is the total value of the hand, treating aces as 11.
	soft Value
}

// NewHand creates a new blackjack hand.
func NewHand(rng *rand.Rand) *Hand {

	hand := &Hand{}

	for {

		card := NewCard(rng)

		hand.Cards = append(hand.Cards, card)
		hand.soft += card.Value

		switch card.Rank {
		// Jokers do not change the hand value, so we can skip recomputing it.
		case RankJoker:
			continue
		case RankAce:
			hand.aces++
		}

		if hand.valueBest() >= Value(config.StandLimit) {
			break
		}
	}

	return hand
}

// valueBest returns the best value of the player's hand.
//
// This considers all possible values by promoting aces from 1 to 11.
func (h *Hand) valueBest() Value {

	best := h.soft

	for i := 0; i <= h.aces; i++ {

		value := h.soft + Value(10*i)

		if value > Value(config.BustLimit) {
			continue
		}

		if value > best {
			best = value
		}
	}

	return best
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
