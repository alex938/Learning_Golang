package main

import "testing"

func Test_new_deck(t *testing.T) {
	d := new_deck()
	var found int = 0

	if len(d) != 52 {
		t.Errorf("Exepected deck to be length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card in new deck (before shuffle) to be Ace of Spades, however, first card is %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card in new deck (before shuffle) to be King of Clubs, however, last card is %v", d[len(d)-1])
	}

	for _, card := range d {
		if card == "Ace of Hearts" {
			found = 1
			break
		}
	}

	if found != 1 {
		t.Errorf("Something wrong with deck build - could not find Ace of Hearts")
	}
}
