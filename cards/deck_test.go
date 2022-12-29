package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	const EXPECTED_LENGTH = 52
	if len(d) != EXPECTED_LENGTH {
		t.Errorf("TestNewDeck: Expected deck to have a length of %v, recieved %v", EXPECTED_LENGTH, len(d))
	}

	const EXPECTED_FIRST_ELEMENT = "Two of Clubs"
	const EXPECTED_LAST_ELEMENT = "Ace of Spades"

	if d[0] != EXPECTED_FIRST_ELEMENT {
		t.Errorf("TestNewDeck: Expected first element to be %v, recieved %v", EXPECTED_FIRST_ELEMENT, d[0])
	}
	if d[len(d)-1] != EXPECTED_LAST_ELEMENT {
		t.Errorf("TestNewDeck: Expected last element to be %v, recieved %v", EXPECTED_LAST_ELEMENT, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in loaded deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
