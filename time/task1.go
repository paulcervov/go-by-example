package main

import (
	"fmt"
	"time"
)

// Write a program that counts the number of minutes in a time period.
// Sample Input: 1h30m
// Sample Output: 1h30m = 90 min
func main() {
	var s string

	fmt.Scan(&s)

	d, _ := time.ParseDuration(s)

	fmt.Println(s, "=", d.Minutes(), "min")
}
