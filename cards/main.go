package main

import "os"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("testdeck.txt")
	testdeck := readDeckFromFile("testdeck.txt")
	testdeck.print()
	os.Remove("testdeck.txt")
}

// Array is a fixed length list of things.
// A slice is an array that can grow or shrink.
// Arrays and slices only can contain one data type.
