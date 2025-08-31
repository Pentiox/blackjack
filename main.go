package main

import (
	"log"
	"sync"

	"golang.org/x/sync/errgroup"

	"blackjack/config"
)

func main() {

	// Statistics is the overall statistics over all blackjack games.
	//
	// The mutex is used to synchronize the blackjack games running in parallel.
	statistics := &Statistics{}
	statisticsMu := sync.Mutex{}

	group := errgroup.Group{}
	group.SetLimit(config.NumGamesParallel)

	for range config.NumGames {

		group.Go(func() error {

			game := NewGame()

			result := game.Play()

			statisticsMu.Lock()
			statistics.Add(result)
			statisticsMu.Unlock()

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		log.Fatalf("group wait: %s", err)
	}

	statistics.Print()
}
