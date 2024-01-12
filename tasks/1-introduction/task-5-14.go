package main

import "fmt"

// Given a natural number, print its last digit.
// 123 -> 3
func main() {

	var number int

	fmt.Scan(&number) // doesn't exceed 10000

	number = number % 10

	fmt.Println(number)
}
