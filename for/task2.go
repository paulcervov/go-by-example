package main

import "fmt"

// Write a program that finds the sum of two-digit numbers divisible by 8 in a sequence of numbers.
// The program in the first line receives the number n - the number of numbers in the sequence,
// in the second line - n numbers included in this sequence.
func main() {
	var i, digitCount, digit, digitSum int
	fmt.Scan(&digitCount)

	for i = 1; i <= digitCount; i++ {
		fmt.Scan(&digit)

		if (digit > 9) && (digit < 100) && ((digit % 8) == 0) {
			digitSum += digit
		}
	}

	fmt.Println(digitSum)
}
