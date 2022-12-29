package main

import "fmt"

func main() {
	output := fizzbuzz(15)
	for _, e := range output {
		fmt.Println(e)
	}
}
