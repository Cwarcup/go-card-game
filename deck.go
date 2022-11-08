package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// pass in a deck and the number of cards you want to deal
// returns two decks, the first is the hand and the second is the remaining deck
func deal(d deck, handSize int) (deck, deck) {
	// return everything from the start of the deck to the handSize
	// and everything from the handSize to the end of the deck
	return d[:handSize], d[handSize:]
}

// converts a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save the deck to a file on the local machine
func (d deck) saveToFile(filename string) error {
	// convert the deck to a string
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

// read a deck from a file
func newDeckFromFile(filename string) deck {
	// read the file
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// convert the byte slice to a string
	s := strings.Split(string(bs), ",")

	// convert the string to a deck
	return deck(s)
}

// shuffle the deck
// need to accept a deck, shuffle around the cards, and return the deck
func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// for each card in the deck
	for index := range d {
		// generate ran num between 0 and the length of the deck
		newPosition := r.Intn(len(d) - 1)
		// swap the current card with the card at the random index
		d[index], d[newPosition] = d[newPosition], d[index]
	}

	// return the deck
	return d
}
