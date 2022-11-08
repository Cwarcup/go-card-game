package main

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 5)
	hand.saveToFile("my_cards")

	// create a new deck from a file
	newCards := newDeckFromFile("my_cards")
	newCards.print()

}
