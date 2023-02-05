package main

func main() {
	cards := newDeck()

	hand, remainingDeck := cards.deal(5)

	// Second approach
	// hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
