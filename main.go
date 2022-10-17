package main

func main() {
	cards := newDeckFromFile("deck_1")

	cards.shuffle()
	cards.print()
}
