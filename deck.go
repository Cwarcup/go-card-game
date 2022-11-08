package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	// creates a new deck of cards
	cards := deck{}

	// create a slice of card suits
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	// create a slice of card values
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// loop through the suits and values to create a deck of cards
	for _, suit := range cardSuits {
		// loop through the values
		for _, value := range cardValues {
			// append the cards to the deck
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// loop through the deck and print each card
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}