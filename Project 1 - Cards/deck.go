package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new deck type
//which is a slice of strings

type deck []string

func new_deck() deck {
	cards := deck{}

	var card_suites []string = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var card_values []string = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range card_suites {
		for _, values := range card_values {
			cards = append(cards, values+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, j := range d {
		fmt.Printf("%d %s\n", i+1, j)
	}
}

func deal(d deck, hand_size int) (deck, deck) {
	return d[:hand_size], d[hand_size:]
}

func (d deck) deck_to_string() string {
	//cast to strings
	return strings.Join([]string(d), ",")
}

func (d deck) save_to_file(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.deck_to_string()), 0666)
}

func load_deck(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s1 := string(bs)
	s2 := strings.Split(s1, ",")
	return deck(s2)
}

func (d deck) shuffle() {
	t := time.Now().UnixNano()
	source := rand.NewSource(t)
	r := rand.New(source)

	for i := range d {
		new_pos := r.Intn(len(d) - 1)
		d[i], d[new_pos] = d[new_pos], d[i]
	}
}

/* Bad practice example! avoid self or this!
type laptopSize float64

func (this laptopSize) getSizeOfLaptop() laptopSize {
    return this
}
*/

/*
func (d deck) new_card() string {
	return "new card added"
}
*/
