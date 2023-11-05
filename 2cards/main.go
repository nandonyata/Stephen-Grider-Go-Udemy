package main

func main() {
	cards := newDeck()
	// cards.saveToFile("my_deck")

	cards.shuffle()
	cards.print()

	// value := newDeckFromFile("my_deck")

	// fmt.Println(value)
}
