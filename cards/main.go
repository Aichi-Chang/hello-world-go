package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// is similar to this, := allows Go to determin the data type
	// card := newCard()
	// fmt.Println(card)

	cards := []string{"Ace if Diamonds", newCard()}
	// append does not modify the existing slides, it returns a new slices
	cards = append(cards, "Six of Spades")

	// with for loops in Go,
	// each loop we need to declear the type of data
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// when return, we need to tell Go what type of data we want it to return
func newCard() string {
	return "Five of Diamonds"
}
