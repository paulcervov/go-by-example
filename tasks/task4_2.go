package main

import "fmt"

// Given a non-negative integer not exceeding 10000.
// Find and print the first digit of the number.
func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println(getFirstNumberDigit(number))
}

// 9999 / 10 = 999.9 -> 999
// 999 / 10 = 99.9 -> 99
// 99 / 10 = 9.9 -> 9
func getFirstNumberDigit(number int) int {
	if number < 10 {
		return number
	} else {
		return getFirstNumberDigit(number / 10)
	}
}
