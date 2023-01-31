package main

import "fmt"

// Given a natural number, print its last digit.
func main() {

	var number int
	fmt.Scan(&number)

	number = number % 10

	fmt.Println(number)
}
