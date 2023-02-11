package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()

	cards.saveToFile("my_cards")

	cards = newDeckFromFile("./my_cards")

	fmt.Println("\nShuffle...")

	cards.shuffle()

	cards.print()

	cards.deal(5)

	hand, remainingDeck := cards.deal(5)

	fmt.Println("\nDealing...\n-------------\nHand:")
	hand.print()

	fmt.Println("\nRemaining deck:")
	remainingDeck.print()
}
