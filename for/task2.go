package main

import "fmt"

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
