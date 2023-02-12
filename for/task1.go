package main

import "fmt"

// It is required to write a program, during which two natural numbers
// A and B are read from the keyboard (each is not more than 100, A < B).
// Print the sum of all numbers from A to B inclusive.
func main() {
	var a, b, i, sum int

	fmt.Scan(&a, &b)

	for i = a; i <= b; i++ {
		sum += i
	}

	fmt.Println(sum)
}
