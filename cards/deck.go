package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			cards = append(cards, value+" of "+suit)
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

func newDeckFromFile(filename string) deck {
	// readFile will return bytes and error if any
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		// if anything goes run, log the error, and exit the program
		os.Exit(1)
	}

	// split it in to slice of strings and save it in a variable "s"
	s := strings.Split(string(bs), ",")
	// and convert the slice of string to a deck
	return deck(s)
}

// revceiver
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap the position with the original position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
