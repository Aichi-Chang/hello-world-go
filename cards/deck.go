package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// append does not modify the existing slides, it returns a new slices
			// that's why we need to put cards first, like JS [...a, b, c]
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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// receiver
func (d deck) saveToFile(filename string) error {
	// 0666 means everyone can read and write the file (permission from the WriteFile func)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
