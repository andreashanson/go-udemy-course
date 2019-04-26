package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Dimonds", "Hearts", "Spades", "Cloves"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) deckToString() string {
	s := []string(d)
	return strings.Join(s, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.deckToString()), 0666)
}

func readDeckFromFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := string(bs)
	ss := strings.Split(s, ",")
	return deck(ss)
}

func (d deck) shuffle() {

}
