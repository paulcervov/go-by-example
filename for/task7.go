package main

import "fmt"

// Two numbers are given. Determine the digits included in the record of both the first and second numbers.
// The program receives two numbers as input. It is guaranteed that digits in numbers do not repeat. Numbers ranging from 0 to 10000.
// The program should display the numbers that are in both numbers, separated by a space. The numbers are displayed in the order they appear in the first number.
// Sample Input: 564 8954
// Sample Output: 5 4
func main() {
	var a, remainderA, nextDigitA, b, remainderB, nextDigitB, aCounter, bCounter int

	fmt.Scan(&a, &b)

	remainderA = a
	for aCounter = 1000; aCounter > 0; aCounter /= 10 {

		nextDigitA = remainderA / aCounter
		if nextDigitA == 0 {
			continue
		}

		remainderB = b
		for bCounter = 1000; bCounter > 0; bCounter /= 10 {

			nextDigitB = remainderB / bCounter
			if nextDigitB == 0 {
				continue
			}

			if nextDigitB == nextDigitA {
				fmt.Printf("%d ", nextDigitB)
			}

			remainderB = remainderB % bCounter
		}

		remainderA = remainderA % aCounter
	}
}
