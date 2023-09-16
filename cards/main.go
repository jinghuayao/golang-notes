package main

import "fmt"


func main() {
	cards := newDeck()
	// hand, remaingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println()
	// remaingCards.print()
	// fmt.Println()
	// cards.print()

	// fmt.Println(cards.toString())

	cards.saveToFile("my_cards")


	new_cards := newDeckFromFile("my_cards")

	new_cards.print()
	fmt.Println()
	new_cards.shuffle()
	new_cards.print()



	

}


