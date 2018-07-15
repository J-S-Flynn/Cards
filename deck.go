package main // package name of main for an exicutable programme

//imports we will need to work with in this programme
//in this perticuler programme only the fmt, and io/ioutil package is needed for print
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := 0; i < len(d); i++ {

		num := r.Intn(len(d) - 1)

		d[i], d[num] = d[num], d[i]
	}

	for k := 0; k < 5; k++ {
		for i := 0; i < len(d); i++ {

			numOne := r.Intn(len(d) - 1)
			numTwo := r.Intn(len(d) - 1)

			d[numOne], d[numTwo] = d[numTwo], d[numOne]
		}
	}
}

//this is a good func to look at as it displays how go allows
//for the defining of multiple return types. this func is deck deck, but there is no reason we could
//not have a deck int, or int string etc.
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

//calles the jion method on a slice of strings to return a single string, we use the method jion from the strings
//the reciver can be called on aby var of type deck and will concatinate the vales of each slice into a single string to
//be returned.
func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

//saves the deck , simple reciiver given any var of type deck it will call the toString func
//and then save the content of the string returned as a byte slice to the file via io/ioutil writeFile method
//the method takes in a file name(string) a slice of bytes and the permissions allowed to do it in this case 0666
//more readingwill need to be done in regards to the permisions,
func (d deck) saveDeck(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//func opens a file given a file name as a string then checks to see if file is there
//if it is it siply converst the silce of bytes withing the file back in to a string and then splits that string
//using the diliniator which in this case is a comma ",". then casts the slice of strings to type deck and retutns the value
func openFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error, Error, Error : danger Will Robinson : ", err)
		os.Exit(1)
	}

	hand := strings.Split(string(bs), ",")
	return deck(hand)
}
