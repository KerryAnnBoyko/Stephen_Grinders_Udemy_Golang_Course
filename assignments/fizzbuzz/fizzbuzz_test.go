package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	fb := fizzbuzz(15)

	if len(fb) != 15 {
		t.Errorf("TestNewDeck: Expected deck to have a length of %v, recieved %v", 15, len(fb))
	}

	if fb[0] != "1" {
		t.Errorf("Expected first element to be '1', recieved %v", fb[0])
	}

	if fb[2] != "fizz" {
		t.Errorf("Expected third element to be 'fizz', recieved %v", fb[2])
	}

	if fb[4] != "buzz" {
		t.Errorf("Expected fifth element to be 'buzz', recieved %v", fb[4])
	}

	if fb[len(fb)-1] != "FizzBuzz!" {
		t.Errorf("Expected last element to be 'FizzBuzz!', recieved %v", fb[4])
	}
}
