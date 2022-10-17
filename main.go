package main

func main() {
	// cards := newDeck()
	// cards.saveDeckToFile("deck_1")
	cards := newDeckFromFile("deck_1")

	cards.shuffle()
	cards.print()
}
