package main

import (
	"log"

	"golang.org/x/sync/errgroup"

	"blackjack/config"
)

func main() {

	statistics := &Statistics{}

	group := errgroup.Group{}
	group.SetLimit(config.NumGamesParallel)

	for range config.NumGames {

		group.Go(func() error {

			game := NewGame()

			result := game.Play()

			statistics.Add(result)

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		log.Fatalf("group wait: %s", err)
	}

	statistics.Print()
}
