package deck

import (
	"blackjack/config"
	"fmt"
	"math/rand"
	"strings"
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

	// value is the hand value.
	value Value
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

		// This can be optimized by avoiding recomputation in some cases.
		// For example, we can cache the hand value until we encounter an ace.
		hand.value = hand.valueBest()

		if hand.value >= Value(config.StandLimit) {
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

// IsBlackjack checks if the hand is a blackjack.
//
// The value of a hand with two cards can never exceed a bust limit greater than 21.
// As such, we redefine a blackjack hand as one that is equal to the bust limit using the least amount of cards.
func (h *Hand) IsBlackjack() bool {
	return h.value == Value(config.BustLimit) && len(h.Cards) == BlackjackHandSizes
}

// IsBusted checks if the hand is busted.
func (h *Hand) IsBusted() bool {
	return h.value > Value(config.BustLimit)
}

// IsEights checks if the hand contains at least two eights.
func (h *Hand) IsEights() bool {

	eights := 0

	for _, card := range h.Cards {
		if card.Rank == RankEight {
			eights++
		}
	}

	return eights >= 2
}

// IsEqual checks if the hand is equal to another hand.
func (h *Hand) IsEqual(hand *Hand) bool {
	return h.value == hand.value
}

// IsBetter checks if the hand is better than another hand.
func (h *Hand) IsGreater(hand *Hand) bool {
	return h.value > hand.value
}

// IsKingAndQueen checks if the hand contains both a king and a queen.
func (h *Hand) IsKingAndQueen() bool {

	hasKing := false
	hasQueen := false

	for _, card := range h.Cards {
		switch card.Rank {
		case RankQueen:
			hasQueen = true
		case RankKing:
			hasKing = true
		}
	}

	return hasKing && hasQueen
}

// String returns a string representation of the hand.
func (h *Hand) String() string {

	var cards []string

	for _, card := range h.Cards {
		cards = append(cards, card.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(cards, ", "))
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
