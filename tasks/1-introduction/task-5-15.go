package main

import "fmt"

// Given a natural number, print its tens number (second number from the right)
// 123 -> 2
func main() {

	var number int

	fmt.Scan(&number) // doesn't exceed 10000

	number = (number % 100) / 10

	fmt.Println(number)
}
