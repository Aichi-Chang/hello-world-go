package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// %v is a placeholder for len(d)
		// when reject something to a string, we need to add %v
		t.Errorf("Expected dack length of 16, but got %v", len(d))
	}

}
