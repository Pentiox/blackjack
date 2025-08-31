package main

import (
	"log"
)

// Statistics is the statistical data for a single blackjack game.
type Statistics struct {
	BlackjacksDealer int
	BlackjacksPlayer int
	BustsDealer      int
	BustsPlayer      int
	WinsDealer       int
	WinsPlayer       int
	TotalHands       int
	TotalPushes      int
}

// Add combines the statistics from another Statistics object into the current one.
func (s *Statistics) Add(stats *Statistics) {
	s.BlackjacksDealer += stats.BlackjacksDealer
	s.BlackjacksPlayer += stats.BlackjacksPlayer
	s.BustsDealer += stats.BustsDealer
	s.BustsPlayer += stats.BustsPlayer
	s.WinsDealer += stats.WinsDealer
	s.WinsPlayer += stats.WinsPlayer
	s.TotalHands += stats.TotalHands
	s.TotalPushes += stats.TotalPushes
}

// Print displays the current statistics.
func (s *Statistics) Print() {
	log.Printf("Dealer blackjacks: %d", s.BlackjacksDealer)
	log.Printf("Player blackjacks: %d", s.BlackjacksPlayer)
	log.Printf("Dealer busts: %d", s.BustsDealer)
	log.Printf("Player busts: %d", s.BustsPlayer)
	log.Printf("Dealer wins: %d", s.WinsDealer)
	log.Printf("Player wins: %d", s.WinsPlayer)
	log.Printf("Total hands: %d", s.TotalHands)
	log.Printf("Total pushes: %d", s.TotalPushes)
}
