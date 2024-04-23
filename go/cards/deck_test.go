package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	length_expected := 20
	if len(d) != length_expected {
		t.Errorf("Expected deck length of %v, but got %v", length_expected, len(d))
	}

	first_card := "Ace of Spades"
	if d[0] != first_card {
		t.Errorf("Expected first card to be %v, but got %v", first_card, d[0])
	}

	last_card := "Five of Clubs"
	if d[len(d)-1] != last_card {
		t.Errorf("Expected last card to be %v, but got %v", last_card, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	expected_length := 20
	if len(loadedDeck) != expected_length {
		t.Errorf("Expected %v cards in deck, but got %v", expected_length, len(loadedDeck))
	}

	os.Remove(filename)
}
