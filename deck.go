package main

import (
	"fmt"
	"os"
	"strings"
)

// create type deck
// ..which is slice of strings

type deck []string

func (d deck) print() { //reciever on deck type
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "two", "three", "four", "five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) { // takes deck as argument
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}
func newDeckFromFile(fileName string) deck {
	bSlice, err := os.ReadFile(fileName)
	if err != nil {

		fmt.Println("Error", err)
		os.Exit(1)
	}
	s := strings.Split(string(bSlice), ",")
	return deck(s)
}
