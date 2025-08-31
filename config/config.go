package config

// Global configuration constants for all blackjack games.
const (

	// BustLimit is the maximum value a player's hand can reach before busting.
	//
	// This is typically set to 21 in standard blackjack.
	BustLimit int = 25

	// MaxHands is the maximum number of hands that can be played in a single game across all players.
	MaxHands int = 1000000

	// NumGames is the total number of blackjack games played.
	NumGames int = 1

	// NumGamesParallel is the total number of blackjack games played in parallel.
	NumGamesParallel int = 32

	// PlayerHands is the number of hands each player is dealt simultaneously.
	PlayerHands int = 3

	// PlayerStartingChips is the number of chips each player starts with.
	PlayerStartingChips int = 10000

	// PlayersPerGame is the number of players aside from the dealer in a given game.
	PlayersPerGame int = 5

	// StandLimit is the minimum value a player must reach to stand.
	StandLimit int = 19
)
