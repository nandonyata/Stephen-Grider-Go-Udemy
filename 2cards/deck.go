package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

// every type deck var can use this func (like class in JS)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, "<", card)
	}
}

// just a normal function that returns new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, v := range cardSuits {
		for _, w := range cardValues {
			cards = append(cards, w + " of " + v)
		}
	}

	return cards
}

// just a normal function with 2 return
func deal (d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// var result string

	result := strings.Join(d, ",")

	return result
} 

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	 value, err := os.ReadFile(fileName)

	 if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	 }

	//  string(value) //Ace of Spades,Two of Spades ...
	s := strings.Split(string(value), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newRandomIndex := rand.Intn(len(d) - 1)
		d[i], d[newRandomIndex] = d[newRandomIndex], d[i]
	}
}