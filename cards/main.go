package main

// import "fmt"

func main() {
	// cards := newCard()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

	//cards := newCard()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//card.
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
