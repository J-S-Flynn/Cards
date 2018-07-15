package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	//test to see cards are of correct length
	if len(d) != 52 {
		t.Errorf("ecpected 52, got %v", len(d))
	}

	//boundry test of lower and upper limits of deck
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of hearts , got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected Ace of hearts , got %v", d[len(d)-1])
	}
}
