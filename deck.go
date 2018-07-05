package main

import (
	"fmt"
)

// creates new type of deck

type deck []string

//reciver ?
func (cards deck) print() {

	for i, card := range cards {

		fmt.Println(i, card)
	}
}
