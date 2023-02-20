package main

import "fmt"

// The deposit in the bank is x rubles. It increases annually by p percent.
// Every year the amount of the contribution becomes larger.
// Determine in how many years the contribution will be at least y rubles.
func main() {
	var x, p, y, i int

	for fmt.Scan(&x, &p, &y); x < y; i++ {
		x += (x * p) / 100
	}

	fmt.Println(i)
}
