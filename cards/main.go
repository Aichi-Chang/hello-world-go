package main

func main() {
	// the deck is defined in deck.go
	// we also need to run deck.go while running main.go
	// the command would be 'go run main.go deck.go'
	cards := newDeck()
	// append does not modify the existing slides, it returns a new slices
	// that's why we need to put cards first, like JS [...a, b, c]
	// cards = append(cards, "Six of Spades")

	cards.print()
}
