package main

func main() {
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_card")
	cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// greeting := "hi there"
	// fmt.Println([]byte(greeting))
}

func newCard() string {
	return "five of dia"
}
