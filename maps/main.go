package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "FF0000",
		"green": "00FF00",
		"blue":  "0000FF",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}

/* Maps:
    * all keys must be the same type
		* all values must be the same type.
		* can be iterated over
		* don't need to know all keys at compile time.
		* passes by reference!
	Structs:
	  * Values can be different type
		* Keys don't support indexing (and therefore no iterating)
		* You need to know all fields at compile time
		* passes by value! (to pass by reference, use the pointer)
*/
