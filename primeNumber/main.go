package main

import (
	"fmt"
)

func main() {

	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range num {
		if n%2 == 0 {
			fmt.Printf("%d is even \n", n)
		} else {
			fmt.Printf("%d is odd \n", n)
		}
	}
}
