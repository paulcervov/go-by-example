package main

import "fmt"

// Find the first number between 1 and n inclusive that is a multiple of c but NOT a multiple of d.
func main() {

	var i, max, a, b int

	fmt.Scan(&max, &a, &b)

	for i = 1; i <= max; i++ {

		if (i%a == 0) && (i%b != 0) {
			fmt.Println(i)
			break
		}
	}
}
