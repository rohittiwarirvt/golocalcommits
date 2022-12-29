package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newCard()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected the first Item to be Ace of Spades  but Got %v", d[0])
	}

	if (d[len(d)-1]) != "Four of Clubs" {
		t.Errorf("Expected the lat Item to be Four of Clubs but Got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testing")

	deck := newCard()

	deck.saveToFile("_testing")

	loadedDeck := newDeckFromFile("_testing")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_testing")
}
