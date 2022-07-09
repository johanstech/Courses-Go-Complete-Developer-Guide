package main

import "fmt"

func main() {
	//? Long way of declaring variable
	// var card string = "Ace of Spades"
	//? Walrus operator := is only used when declaring a variable, not reassigning
	// card := newCard()

	cards := deck{newCard(), "Ace of Hearts"}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
