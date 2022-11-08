package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

// test to create a deck, save it to a file and then load it back in
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// delete the file if it exists
	os.Remove("_decktesting")

	// create a new deck
	d := newDeck()

	// save the deck to a file
	d.saveToFile("_decktesting")

	// load the deck from the file
	loadedDeck := newDeckFromFile("_decktesting")

	// check the deck length
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	// delete the file
	os.Remove("_decktesting")
}

// test to shuffle a deck
// should create a new deck, shuffle it and then check that the first card is not the same as the first card in the new deck
func TestShuffle(t *testing.T) {
	// create a new deck
	d := newDeck()

	// shuffle the deck
	d.shuffle()

	// check the first card is not the same as the first card in the new deck
	if d[0] == "Ace of Spades" {
		t.Errorf("Expected first card to be different, but got %v", d[0])
	}
}
