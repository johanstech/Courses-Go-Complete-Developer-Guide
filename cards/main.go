package main

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// filename := "my_cards.txt"
	// cards.saveToFile(filename)
	// cards := readDeckFromFile(filename)

	cards.shuffle()
	cards.print()
}
