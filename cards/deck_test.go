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

	firstCard := "Ace of Spades"
	if d[0] != firstCard {
		t.Errorf("Expected first card to be %v, but got %v", firstCard, d[0])
	}

	lastCard := "King of Hearts"
	if d[len(d)-1] != lastCard {
		t.Errorf("Expected last card to be %v, but got %v", lastCard, d[len(d)-1])
	}
}

func TestSaveToFileAndReadDeckFromFile(t *testing.T) {
	filename := "_decktesting.txt"

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loaded := readDeckFromFile(filename)
	if len(loaded) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loaded))
	}

	os.Remove(filename)
}
