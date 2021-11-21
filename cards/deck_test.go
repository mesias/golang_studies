package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("expected deck lenght of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	deckfile := "_decktesting"

	os.Remove(deckfile)

	deck := newDeck()
	deck.saveToFile(deckfile)

	loadedDeck := newDeckFromFile(deckfile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(deckfile)

}
