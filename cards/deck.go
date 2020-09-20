package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// receiver, d is like 'this' in JS
func (d deck) print() {
	// with for loops in Go,
	// each loop we need to declear the type of data
	for i, card := range d {
		fmt.Println(i, card)
	}
}
