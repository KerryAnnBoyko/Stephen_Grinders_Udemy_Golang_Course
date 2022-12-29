package main

import "strconv"

func fizzbuzz(count int) []string {
	output := []string{}
	for i := 1; i <= count; i++ {
		if i%15 == 0 {
			output = append(output, "FizzBuzz!")
		} else if i%5 == 0 {
			output = append(output, "buzz")
		} else if i%3 == 0 {
			output = append(output, "fizz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}
	return output
}
