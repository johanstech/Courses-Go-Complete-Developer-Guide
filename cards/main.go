package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	//? Walrus operator := is only used when declaring a variable, not reassigning
	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
