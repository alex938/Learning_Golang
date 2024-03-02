package main

func main() {
	//var cards string = new_card()
	//var cards := "Ace of Spades"
	//fmt.Printf("%s", cards)

	//create a deck from file or via function:

	//slice example:
	//var cards []string = deck{new_card(), "Ace of Diamonds"}

	//1.
	//cards := new_deck()
	//2.
	//cards := load_deck("my_cards")
	//3.
	cards := start_game()

	//shuffle the cards
	cards.shuffle()

	//save cards to a file "my_cards"
	cards.save_to_file("my_cards")

	dealt_hand, remaining_deck := deal(cards, 5)

	dealt_hand.print()
	remaining_deck.print()

}

func new_card() string {
	return "Five of Diamonds"
}
