package main

import "testing"

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first index of arr Ace of Spades, but got %v", cards[0])
	}

	if cards[len(cards) - 1] != "Four of Hearts" {
		t.Errorf("Expected last index of arr Four of Hearts, but got %v", cards[len(cards) -1])
	}
}