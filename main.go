package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

	//Declaring a slice
	fmt.Println(cards.toString())

	//To save data to a file
	cards.saveToFile("my_cards")

	//Reading contents from a file and printing it
	cards = newDeckFromFile("my_cards")
	cards.print()

	//Call deal function with the deck of cards and number of hands to be dealt
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	//Iterating over slice object in the receiver function
	cards.print()
	fmt.Println(cards)

	//Type conversion Example
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

}
