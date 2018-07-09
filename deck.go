package main // package name of main for an exicutable programme

//imports we will need to work with in this programme
//in this perticuler programme only the fmt, and io/ioutil package is needed for print
import (
	"fmt"
)

// creates new type of deck
type deck []string

//like manu oop we have the ability to define fun that we can then use to

// this is a proper deck
//we define the deck and a return type.
func newDeck() deck {

	cards := deck{}

	suits := []string{" of Hearts", " of Dimonds", " of Clubs", " of Spades"}
	cardNo := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "nine",
		"Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, card := range cardNo {

			cards = append(cards, card+suit)
		}
	}

	return cards
}

//reciver Allows the used of a func by any varible of type
//i.e in this case any var of type deck will have access to these funcs
//reciver is a var that is passed before the func name annotation
func (d deck) print() {

	for _, card := range d {

		fmt.Println(card)
	}
}

//shffle is a little more complicate das it will
func (d deck) shuffle() {
	fmt.Println("shuffle", d)
}

//this is a good func to look at as it displays how go allows
//for the defining of multiple return types. this func is deck deck, but there is no reason we could
//not have a deck int, or int string etc.
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {

	var compleatString string

	for i := 0; i < len(d); i++ {

		compleatString = compleatString + d[i] + " "
	}

	return compleatString
}

func saveDeck() {

}
