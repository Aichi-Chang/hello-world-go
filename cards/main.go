package main

func main() {
	// the deck is defined in deck.go
	// we also need to run deck.go while running main.go
	// the command would be 'go run main.go deck.go'
	cards := newDeck()

	cards.saveToFile("my_cards")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}
