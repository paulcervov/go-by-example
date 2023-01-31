package main

import "fmt"

func main() {

	var number int

	fmt.Scan(&number)

	if threeDigitIntSum(number/1000) == threeDigitIntSum(number%1000) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func threeDigitIntSum(digit int) int {
	return (digit / 100) + ((digit / 10) % 10) + (digit % 10)
}
