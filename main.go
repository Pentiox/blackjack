package main

import (
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"

	"blackjack/config"
)

func main() {

	// Start a timer to measure how long each game takes.
	start := time.Now()
	defer func() {
		log.Printf("Simulating all games took %s", time.Since(start))
	}()

	// Statistics is the overall statistics over all blackjack games.
	//
	// The mutex is used to synchronize the blackjack games running in parallel.
	statistics := &Statistics{}
	statisticsMu := sync.Mutex{}

	group := errgroup.Group{}
	group.SetLimit(config.NumGamesParallel)

	for i := range config.NumGames {

		group.Go(func() error {

			log.Printf("Game %d started at %s", i, time.Since(start))

			game := NewGame()

			result := game.Play()

			statisticsMu.Lock()
			statistics.Add(result)
			statisticsMu.Unlock()

			log.Printf("Game %d took %s", i, time.Since(start))

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		log.Fatalf("group wait: %s", err)
	}

	statistics.Print()
}
