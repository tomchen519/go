package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("cards")
	cards := readFileToDeck("cards")
	cards.shuffle()
	cards.print()
}
