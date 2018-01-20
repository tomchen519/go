package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := range d {
		j := random.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(ofile string) error {
	return ioutil.WriteFile(ofile, []byte(d.toString()), 0666)
}

func readFileToDeck(ifile string) deck {
	bs, err := ioutil.ReadFile(ifile)
	if err != nil {
		fmt.Println("Error occurred: ", err)
		os.Exit(1)
		return nil
	}

	return deck(strings.Split(string(bs), ","))
}

func load() {
	return
}
