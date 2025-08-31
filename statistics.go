package main

import (
	"log"
	"sync/atomic"
)

// Statistics is the statistical data for a single blackjack game.
type Statistics struct {
	BlackjacksDealer atomic.Int64
	BlackjacksPlayer atomic.Int64
	BustsDealer      atomic.Int64
	BustsPlayer      atomic.Int64
	WinsDealer       atomic.Int64
	WinsPlayer       atomic.Int64
	TotalHands       atomic.Int64
	TotalPushes      atomic.Int64
}

// Add combines the statistics from another Statistics object into the current one.
func (s *Statistics) Add(stats *Statistics) {
	s.BlackjacksDealer.Add(stats.BlackjacksDealer.Load())
	s.BlackjacksPlayer.Add(stats.BlackjacksPlayer.Load())
	s.BustsDealer.Add(stats.BustsDealer.Load())
	s.BustsPlayer.Add(stats.BustsPlayer.Load())
	s.WinsDealer.Add(stats.WinsDealer.Load())
	s.WinsPlayer.Add(stats.WinsPlayer.Load())
	s.TotalHands.Add(stats.TotalHands.Load())
	s.TotalPushes.Add(stats.TotalPushes.Load())
}

// Print displays the current statistics.
func (s *Statistics) Print() {
	log.Printf("Dealer blackjacks: %d", s.BlackjacksDealer.Load())
	log.Printf("Player blackjacks: %d", s.BlackjacksPlayer.Load())
	log.Printf("Dealer busts: %d", s.BustsDealer.Load())
	log.Printf("Player busts: %d", s.BustsPlayer.Load())
	log.Printf("Dealer wins: %d", s.WinsDealer.Load())
	log.Printf("Player wins: %d", s.WinsPlayer.Load())
	log.Printf("Total hands: %d", s.TotalHands.Load())
	log.Printf("Total pushes: %d", s.TotalPushes.Load())
}
