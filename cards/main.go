package main

func main() {
	//cards := newDeck()
	//fmt.Println("PRINT CARDS")
	//cards.printCards()
	//cards.deckToString()
	//fmt.Println("PRINT LONG STRING")
	//fmt.Println(cards)
	//cards2 := newDeck()
	//fmt.Println(cards2.deckToString())
	//cards3 := newDeck()
	//cards3.saveToFile("MY_DECK")
	//hand, _ := deal(cards3, 2)
	//hand.saveToFile("MY_HAND")
	//myHand := readDeckFromFile("MY_HAND")
	//myHand.printCards()
	//myHand.shuffle()
	//myHand.printCards()

	deck2 := newDeck()
	deck2.printCards()
	deck2.shuffle()
	deck2.printCards()

}
