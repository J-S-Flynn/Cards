package main

func main() {

	cards := deck{"Ace of Spades", newCard()}

	cards = append(cards, "Ace of Clubs")
	cards = append(cards, "Ace of Dimonds")

	cards.print()
}

func newCard() string {

	return "Ace of Dimonds"
}
