package main

import (
	"log"
)

// Statistics is the statistical data for a single blackjack game.
type Statistics struct {
	BlackjacksDealer    int
	BlackjacksPlayer    int
	BustsDealer         int
	BustsPlayer         int
	EightsDealer        int
	EightsPlayer        int
	JokersDealer        int
	JokersPlayer        int
	KingAndQueensDealer int
	KingAndQueensPlayer int
	TotalHands          int
	TotalPushes         int
	WinsDealer          int
	WinsPlayer          int
}

// Add combines the statistics from another Statistics object into the current one.
func (s *Statistics) Add(stats *Statistics) {
	s.BlackjacksDealer += stats.BlackjacksDealer
	s.BlackjacksPlayer += stats.BlackjacksPlayer
	s.BustsDealer += stats.BustsDealer
	s.BustsPlayer += stats.BustsPlayer
	s.EightsDealer += stats.EightsDealer
	s.EightsPlayer += stats.EightsPlayer
	s.JokersDealer += stats.JokersDealer
	s.JokersPlayer += stats.JokersPlayer
	s.KingAndQueensDealer += stats.KingAndQueensDealer
	s.KingAndQueensPlayer += stats.KingAndQueensPlayer
	s.TotalHands += stats.TotalHands
	s.TotalPushes += stats.TotalPushes
	s.WinsDealer += stats.WinsDealer
	s.WinsPlayer += stats.WinsPlayer
}

// Print displays the current statistics.
func (s *Statistics) Print() {
	log.Printf("Dealer blackjacks: %d", s.BlackjacksDealer)
	log.Printf("Player blackjacks: %d", s.BlackjacksPlayer)
	log.Printf("Dealer busts: %d", s.BustsDealer)
	log.Printf("Player busts: %d", s.BustsPlayer)
	log.Printf("Dealer eights: %d", s.EightsDealer)
	log.Printf("Player eights: %d", s.EightsPlayer)
	log.Printf("Dealer jokers: %d", s.JokersDealer)
	log.Printf("Player jokers: %d", s.JokersPlayer)
	log.Printf("Dealer kings and queens: %d", s.KingAndQueensDealer)
	log.Printf("Player kings and queens: %d", s.KingAndQueensPlayer)
	log.Printf("Total hands: %d", s.TotalHands)
	log.Printf("Total pushes: %d", s.TotalPushes)
	log.Printf("Dealer wins: %d", s.WinsDealer)
	log.Printf("Player wins: %d", s.WinsPlayer)
}
