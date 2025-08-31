package deck

import (
	"blackjack/config"
	"math/rand"
)

// Rank is a blackjack card's rank.
type Rank int64

// List of blackjack card ranks.
const (
	RankJoker Rank = iota
	RankAce
	RankTwo
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
)

// Value is a blackjack card's value.
type Value int

// RankValues maps blackjack card ranks to their values.
var RankValues = map[Rank]Value{
	RankJoker: 0,
	RankAce:   1,
	RankTwo:   2,
	RankThree: 3,
	RankFour:  4,
	RankFive:  5,
	RankSix:   6,
	RankSeven: 7,
	RankEight: 8,
	RankNine:  9,
	RankTen:   10,
	RankJack:  10,
	RankQueen: 10,
	RankKing:  10,
}

// Deck is a deck of playing cards.
var Deck = []Rank{
	RankJoker, RankJoker,
	RankAce, RankAce, RankAce, RankAce,
	RankTwo, RankTwo, RankTwo, RankTwo,
	RankThree, RankThree, RankThree, RankThree,
	RankFour, RankFour, RankFour, RankFour,
	RankFive, RankFive, RankFive, RankFive,
	RankSix, RankSix, RankSix, RankSix,
	RankSeven, RankSeven, RankSeven, RankSeven,
	RankEight, RankEight, RankEight, RankEight,
	RankNine, RankNine, RankNine, RankNine,
	RankTen, RankTen, RankTen, RankTen,
	RankJack, RankJack, RankJack, RankJack,
	RankQueen, RankQueen, RankQueen, RankQueen,
	RankKing, RankKing, RankKing, RankKing,
}

// NewRank picks a new blackjack card rank.
func NewRank(rng *rand.Rand) Rank {
	return Deck[rng.Intn(len(Deck))]
}

// BlackjackHandSize is the number of cards in a blackjack hand.
//
// The smallest possible hand for any bust limit is defined as the
// bust limit divided by the value of the highest card.
var BlackjackHandSizes = func() int {

	size := config.BustLimit / 11

	if config.BustLimit%11 != 0 {
		size++
	}

	return size
}()
