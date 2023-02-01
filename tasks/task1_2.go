package main

import "fmt"

// Given a natural three-digit number,
// determine if all its digits are different
func main() {
	var number, x, y, z int

	fmt.Scan(&number)

	x = number / 100        // 1.23 -> 1
	y = (number % 100) / 10 // 2.3 -> 1
	z = number % 10         // 3

	if (x != y) && (y != z) && (x != z) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
