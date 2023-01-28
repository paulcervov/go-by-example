package main

import "fmt"

// The hour hand has turned d degrees since the beginning of the day.
// Determine how many whole hours h and whole minutes m are now.
func main() {

	var d, h, m int
	fmt.Scan(&d)

	// 1h is (360 / 12) = 30d
	h = d / 30 // int 2.33 is 2

	// 1d is (720 / 360) = 2m
	m = 2 * (d % 30)

	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
