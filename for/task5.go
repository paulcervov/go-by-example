package main

import "fmt"

// Write a program that reads integers from the console, one number per line.
// For each number entered check:
// * if the number is less than 10, then skip this number;
// * if the number is greater than 100, then stop reading the numbers;
// * otherwise print this number back to the console on a separate line.
func main() {
	var number int

	for fmt.Scan(&number); number <= 100; fmt.Scan(&number) {
		if number >= 10 {
			fmt.Println(number)
		}
	}
}
