package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck which is a slice of strings
type deck []string

//Add two slices to add all 52 different card combinations,if we are returning value/values from any function you need to specify the type right after function name as shown below "deck"
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	//Use _ to avoid the usage of the index variable and the error that occurs due to non usage
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Writing a receiver function,d refers to the original card object from where print is invoked,ie, cards object in main.go file
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//deal function which deals the card, with return type
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Receiver Function to convert a string slice into a single string seperated by commas
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Save string to file is a receiver function of type deck,0666 =>anyone can read/write into this file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//function to read data from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("The Error is:", err)
		os.Exit(1)
	}
	fmt.Println(string(bs)) // Ace of Spades,Two of Spades,Three of Spades
	s := strings.Split(string(bs), ",")
	return deck(s)
}

//Function to shuffle cards
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
