package main

import "fmt"

// Given a non-negative integer not exceeding 10000.
// Find and print the first digit of the number.
func main() {

	var number, modulo, result int

	fmt.Scan(&number)

	if number < 10 {

		result = number
	} else if number == 10000 {

		result = 1
	} else {

		if (number > 9) && (number < 100) {
			modulo = 10
		} else if (number > 99) && (number < 1000) {
			modulo = 100
		} else if (number > 990) && (number < 10000) {
			modulo = 1000
		}

		result = (number - (number % modulo)) / modulo
	}

	fmt.Println(result)
}
