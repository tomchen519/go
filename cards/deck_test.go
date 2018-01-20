package main

import "testing"

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	if len(testDeck) != 52 {
		t.Errorf("Deck did not generate 52 cards")
	}
}
